package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"gopkg.in/telebot.v3"
)

func executeCommand(c telebot.Context, commandStr string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	cmd := exec.CommandContext(ctx, "./executor.exe", commandStr)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return c.Send("ERROR: Internal pipe failure.")
	}

	if err := cmd.Start(); err != nil {
		return c.Send("ERROR: Failed to start executor.")
	}

	scanner := bufio.NewScanner(stdout)
	var textBuffer strings.Builder

	flushBuffer := func() {
		if textBuffer.Len() > 0 {
			c.Send(textBuffer.String())
			textBuffer.Reset()
		}
	}

	for scanner.Scan() {
		line := scanner.Text()
		cleanLine := strings.TrimSpace(line)

		if cleanLine == "" {
			continue
		}

		switch {
		case strings.HasPrefix(cleanLine, "FILE:"):
			flushBuffer()
			path := strings.TrimSpace(strings.TrimPrefix(cleanLine, "FILE:"))
			doc := &telebot.Document{
				File:     telebot.FromDisk(path),
				FileName: filepath.Base(path),
			}
			c.Send(doc)

		case strings.HasPrefix(cleanLine, "PHOTO:"):
			flushBuffer()
			path := strings.TrimSpace(strings.TrimPrefix(cleanLine, "PHOTO:"))
			if err := c.Send(&telebot.Photo{File: telebot.FromDisk(path)}); err == nil {
				os.Remove(path)
			}

		case strings.HasPrefix(cleanLine, "ERROR:"):
			flushBuffer()
			c.Send("ERROR: " + strings.TrimSpace(strings.TrimPrefix(cleanLine, "ERROR:")))

		default:
			textBuffer.WriteString(line + "\n")
			if textBuffer.Len() >= 1000 {
				flushBuffer()
			}
		}
	}

	flushBuffer()

	if err := cmd.Wait(); err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return c.Send("ERROR: Command timed out after 5 minutes.")
		}
	}
	return nil
}

func main() {
	_ = godotenv.Load() 
	token := os.Getenv("BOT_TOKEN")
	myIDStr := os.Getenv("MY_ID")

	myID, err := strconv.ParseInt(myIDStr, 10, 64)
	if err != nil || token == "" {
		log.Fatal("FATAL: BOT_TOKEN and valid MY_ID must be set in environment.")
	}

	b, err := telebot.NewBot(telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal("FATAL: Could not connect to Telegram: ", err)
	}

	b.Use(func(next telebot.HandlerFunc) telebot.HandlerFunc {
		return func(c telebot.Context) error {
			if c.Sender().ID != myID {
				log.Printf("SECURITY: Blocked unauthorized access from ID %d", c.Sender().ID)
				return nil
			}
			return next(c)
		}
	})

	b.Handle(telebot.OnText, func(c telebot.Context) error {
			return executeCommand(c, c.Text())
	})
	// Listen for files sent to the bot
	b.Handle(telebot.OnDocument, func(c telebot.Context) error {
		doc := c.Message().Document
		if doc == nil {
			return nil
		}

		// 1. Find the PC's Downloads folder automatically
		homeDir, _ := os.UserHomeDir()
		downloadsDir := filepath.Join(homeDir, "Downloads")
		destPath := filepath.Join(downloadsDir, doc.FileName)

		// 2. Notify the user it's downloading
		c.Send(fmt.Sprintf("📥 Downloading '%s' to PC...", doc.FileName))

		// 3. Download and save the file
		err := c.Bot().Download(&doc.File, destPath)
		if err != nil {
			return c.Send(fmt.Sprintf("⚠️ ERROR: Failed to save file - %v", err))
		}

		return c.Send("✅ File successfully saved to your Downloads folder!")
	})
	handleMediaSTT := func(c telebot.Context) error {
		var file *telebot.File
		var ext string

		msg := c.Message()
		if msg.Voice != nil {
			file = &msg.Voice.File
			ext = ".ogg"
		} else if msg.VideoNote != nil {
			file = &msg.VideoNote.File
			ext = ".mp4"
		} else if msg.Audio != nil {
			file = &msg.Audio.File
			ext = ".mp3"
		} else {
			return nil
		}

		os.MkdirAll("data/temp", 0755)
		destPath := filepath.Join("data", "temp", fmt.Sprintf("audio_%s%s", file.FileID, ext))

		c.Send("Downloading media for transcription...")
		if err := c.Bot().Download(file, destPath); err != nil {
			return c.Send(fmt.Sprintf("ERROR: Failed to download file - %v", err))
		}

		return executeCommand(c, "/stt "+destPath)
	}

	b.Handle(telebot.OnVoice, handleMediaSTT)
	b.Handle(telebot.OnVideoNote, handleMediaSTT)
	b.Handle(telebot.OnAudio, handleMediaSTT)
	go b.Start()
	log.Println("Dispatcher online.")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Println("Shutting down.")
	b.Stop()
}
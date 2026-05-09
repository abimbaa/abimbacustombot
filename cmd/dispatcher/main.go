package main

import (
	"bufio"
	"context"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"gopkg.in/telebot.v3"
)

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
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
		defer cancel()

		cmd := exec.CommandContext(ctx, "./executor.exe", c.Text())
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			return c.Send("⚠️ ERROR: Internal pipe failure.")
		}

		if err := cmd.Start(); err != nil {
			return c.Send("⚠️ ERROR: Failed to start executor.")
		}

		scanner := bufio.NewScanner(stdout)
		var textBuffer strings.Builder

		// Helper to flush the buffer
		flushBuffer := func() {
			if textBuffer.Len() > 0 {
				c.Send(textBuffer.String())
				textBuffer.Reset()
			}
		}

		for scanner.Scan() {
			line := scanner.Text() // Don't trim yet to preserve script formatting
			cleanLine := strings.TrimSpace(line)

			switch {
			case strings.HasPrefix(cleanLine, "FILE:"):
				flushBuffer() // Send any pending text before the file
				c.Send(&telebot.Document{File: telebot.FromDisk(strings.TrimPrefix(cleanLine, "FILE:"))})
				
			case strings.HasPrefix(cleanLine, "PHOTO:"):
				flushBuffer()
				c.Send(&telebot.Photo{File: telebot.FromDisk(strings.TrimPrefix(cleanLine, "PHOTO:"))})
				
			case strings.HasPrefix(cleanLine, "ERROR:"):
				flushBuffer()
				c.Send("⚠️ " + strings.TrimPrefix(cleanLine, "ERROR:"))
				
			default:
				textBuffer.WriteString(line + "\n")
				// Telegram limit is 4096. Flush if we get close.
				if textBuffer.Len() >= 4000 {
					flushBuffer()
				}
			}
		} 
		flushBuffer()

		if err := cmd.Wait(); err != nil {
			if ctx.Err() == context.DeadlineExceeded {
				return c.Send("⚠️ ERROR: Command timed out after 5 minutes.")
			}
		}
		return nil
	})

	go b.Start()
	log.Println("Dispatcher online.")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Println("Shutting down.")
	b.Stop()
}
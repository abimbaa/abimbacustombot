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
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			return c.Send("⚠️ ERROR: Internal pipe failure.")
		}

		if err := cmd.Start(); err != nil {
			return c.Send("⚠️ ERROR: Failed to start executor.")
		}

		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if line == "" {
				continue
			}

			switch {
			case strings.HasPrefix(line, "FILE:"):
				path := strings.TrimPrefix(line, "FILE:")
				c.Send(&telebot.Document{File: telebot.FromDisk(path)})
			case strings.HasPrefix(line, "PHOTO:"):
				path := strings.TrimPrefix(line, "PHOTO:")
				c.Send(&telebot.Photo{File: telebot.FromDisk(path)})
			case strings.HasPrefix(line, "ERROR:"):
				c.Send("⚠️ " + strings.TrimPrefix(line, "ERROR:"))
			default:
				c.Send(line)
			}
		}

		if err := cmd.Wait(); err != nil {
			if ctx.Err() == context.DeadlineExceeded {
				return c.Send("⚠️ ERROR: Command timed out after 5 minutes.")
			}
		}
		return nil
	})

	go b.Start()
	log.Println("Dispatcher online and securing system.")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Println("Shutting down gracefully...")
	b.Stop()
}
package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/SevereCloud/vksdk/v2/callback"
	"github.com/SevereCloud/vksdk/v2/events"
	"net/http"
	"os"
)

type VkBot struct {
	cb *callback.Callback
}

func (v *VkBot) Deliver(message Message) {
	panic("implement me")
}

func (v *VkBot) HandleMsg(s string, handler MessageHandler) {
	v.cb.MessageNew(func(ctx context.Context, msg events.MessageNewObject){

	})
}

func setupVKBot(pattern string) (Bot, error) {
	cb, err := initVkCallback()
	if err != nil {
		return nil, fmt.Errorf("setup adapter bot: %w", err)
	}

	http.HandleFunc(pattern, cb.HandleFunc)

	bot := VkBot {
		cb: cb,
	}
	return &bot, nil
}

func initVkCallback() (*callback.Callback, error) {
	cK, err := initConfirmationKey()
	if err != nil {
		return nil, fmt.Errorf("initialise VK Callback: %w", err)
	}

	cb := callback.NewCallback()
	cb.ConfirmationKey = cK

	return cb, nil
}

func initConfirmationKey() (string, error) {
	cK, ok := os.LookupEnv("ADAPTER_BOT_CONFIRMATION_KEY")
	if !ok {
		var err error
		cK, err = promptConfirmationKey()
		if err != nil {
			return "", fmt.Errorf("initialise confirmation key: %w", err)
		}
	}

	return cK, nil
}

func promptConfirmationKey() (string, error) {
	fmt.Print("Enter VK-Api verification code: ")

	cr := bufio.NewReader(os.Stdin)
	cK, err := cr.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("prompt confirmation key: %w", err)
	}

	return cK, nil
}


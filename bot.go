package main

type Bot interface {
	Deliver(Message)
}

type BotState struct{
	Handler MessageHandler
}

type bot struct{
	handlers map[string]MessageHandler
}

type MessageHandler func(Message)

type Message struct {
	text string
}

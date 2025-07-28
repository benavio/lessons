package main

// АБСТРАКЦИЯ

import "fmt"

type Notifier interface {
	Notify(message string)
}

type EmailNotifier struct{}

func (e EmailNotifier) Notify(message string) {
	fmt.Println("Sending email notification:", message)
}

type SMSNotifirer struct{}

func (s SMSNotifirer) Notify(message string) {
	fmt.Println("Sending SMS notification:", message)
}

type TGNotifier struct{}

func (s TGNotifier) Notify(message string) {
	fmt.Println("Sending Telegram notification:", message)
}

func main() {
	var emailNotifier Notifier = EmailNotifier{}
	var smsNotifier Notifier = SMSNotifirer{}
	var slackNotifier Notifier = TGNotifier{}

	emailNotifier.Notify("You've got mail!")
	smsNotifier.Notify("You've got a new SMS!")
	slackNotifier.Notify("You've got a new TGNotiTelegramfier message!")
}

package main

import "fmt"

type NotificationSender interface {
	Send(message string) string
}

type EmailNotification struct {
}

func (e EmailNotification) Send(message string) string {
	return "Sending Email: " + message
}

type SmsNotification struct {
}

func (s SmsNotification) Send(message string) string {
	return "Sending Sms: " + message
}

type PushNotification struct {
}

func (p PushNotification) Send(message string) string {
	return "Sending Push Notification: " + message
}

func Notify(t NotificationSender, message string) {
	fmt.Println(t.Send(message))
}

func main() {
	e := EmailNotification{}
	s := SmsNotification{}
	p := PushNotification{}

	Notify(e, "This is email")
	Notify(s, "This is sms")
	Notify(p, "This is push")
}

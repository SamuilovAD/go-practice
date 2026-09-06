package main

import (
	"errors"
	"fmt"
)

func main() {
	notifications := []Notification{
		{ID: 1, Channel: "email", Message: "Hello"},
		{ID: 2, Channel: "sms", Message: "Your code is 1234"},
		{ID: 3, Channel: "email", Message: "Welcome"},
	}
	senders := map[string]Sender{
		"email": EmailSender{},
		"sms":   SMSSender{},
	}
	notificationsResult, _ := ProcessNotifications(notifications, senders)
	fmt.Printf("%+v", notificationsResult)
}

func ProcessNotifications(
	notifications []Notification,
	senders map[string]Sender,
) ([]ProcessedNotification, error) {
	results := make([]ProcessedNotification, 0, len(notifications))
	for _, notification := range notifications {
		sender, ok := senders[notification.Channel]
		if !ok {
			return results, ErrSenderNotFound
		}
		if err := sender.Send(notification.Message); err != nil {
			return results, fmt.Errorf("sender error. %w", err)
		}
		results = append(results, ProcessedNotification{ID: notification.ID, Status: "sent"})
	}
	return results, nil
}

var ErrSenderNotFound = errors.New("sender not found")

type Notification struct {
	ID      int
	Channel string
	Message string
}
type ProcessedNotification struct {
	ID     int
	Status string
}
type Sender interface {
	Send(message string) error
}
type EmailSender struct{}

func (e EmailSender) Send(message string) error {
	return nil
}

type SMSSender struct{}

func (e SMSSender) Send(message string) error {
	return nil
}

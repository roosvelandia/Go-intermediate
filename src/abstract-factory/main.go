package main

import "fmt"

// SMS and email
type INotificationFactory interface {
	SendNotification()
	GetSender() ISender
}

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

type SMSNotification struct {
}

func (SMSNotification) SendNotification() {
	fmt.Println("Sending Notification via SMS")
}

type SMSNotificationSender struct {
}

func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}
func (SMSNotificationSender) GetSenderChannel() string {
	return "Twilio"
}
func (SMSNotification) GetSender() ISender {
	return SMSNotificationSender{}
}

// email
type EmailNotification struct {
}

func (EmailNotification) SendNotification() {
	fmt.Println("Sending Notification via Email")
}

type EmailNotificationSender struct {
}

func (EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}
func (EmailNotificationSender) GetSenderChannel() string {
	return "SES"
}
func (EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

func getNotificationFactory(notifitactionType string) (INotificationFactory, error) {
	if notifitactionType == "SMS" {
		return &SMSNotification{}, nil
	}
	if notifitactionType == "Email" {
		return &EmailNotification{}, nil
	}
	return nil, fmt.Errorf("Not available")
}
func SendNotification(f INotificationFactory) {
	f.SendNotification()
}
func getMehod(f INotificationFactory) {
	fmt.Println(f.GetSender().GetSenderMethod())
}

func main() {
	smsFactory, _ := getNotificationFactory("SMS")
	emailFactory, _ := getNotificationFactory("Email")
	_, error := getNotificationFactory("Voice")
	SendNotification(smsFactory)
	SendNotification(emailFactory)
	fmt.Println(error)

}

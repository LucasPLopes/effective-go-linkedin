package main

import "fmt"

type Notification struct {
	title    string
	message  string
	image    string
	icon     string
	priority int
	notType  string
}

type NotificationBuilder struct {
	Title    string
	Message  string
	Image    string
	Icon     string
	Priority int
	NotType  string
}

func newNotificationBuilder() *NotificationBuilder {
	return &NotificationBuilder{}
}

func (nb *NotificationBuilder) SetTitle(title string) {
	nb.Title = title
}

func (nb *NotificationBuilder) SetMessage(message string) {
	nb.Message = message
}

func (nb *NotificationBuilder) SetImage(image string) {
	nb.Image = image
}

func (nb *NotificationBuilder) SetIcon(icon string) {
	nb.Icon = icon
}

func (nb *NotificationBuilder) SetPriority(priority int) {
	nb.Priority = priority
}

func (nb *NotificationBuilder) SetNotType(notType string) {
	nb.NotType = notType
}

func (nb *NotificationBuilder) Build() (*Notification, error) {
	if nb.Icon != "" && nb.Title == "" {
		return nil, fmt.Errorf("you need to specify a title when using an icon")
	}

	if nb.Priority > 5 || nb.Priority < 0 {
		return nil, fmt.Errorf("priority must be 0 to 5")
	}

	return &Notification{
		title:    nb.Title,
		message:  nb.Message,
		image:    nb.Image,
		icon:     nb.Icon,
		priority: nb.Priority,
		notType:  nb.NotType,
	}, nil
}

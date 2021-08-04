package main

import (
	"fmt"
)

//Notifies clients that have been charged.
type MessageService interface {
	SendChargeNotification(int) bool
}

//MessageService implementation.
type SMSService struct{}

//Uses MessageService to notify the client.
type MyService struct {
	messageService MessageService
}

//Notifies the client being charged in an SMS. This is the mocked part.
func (sms SMSService) SendChargeNotification(value int) bool {
	fmt.Println("Sending Charge Notification")
	return true
}

func (a MyService) ChargeCustomer(value int) bool {
	a.messageService.SendChargeNotification(value)
	fmt.Printf("Charging customer for the value of %d\n", value)
	return true
}

func main() {

	smsService := SMSService{}
	myService := MyService{smsService}
	myService.ChargeCustomer(100)
}

package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
)

//smsServiceMock
type smsServiceMock struct {
	mock.Mock
}

//smsServiceMock
func (m *smsServiceMock) SendChargeNotification(value int) bool {
	fmt.Println("Mocked charge notification function")
	fmt.Printf("Value passed: %d\n", value)

	//Saves method call and passes the value.
	args := m.Called(value)

	//Return true only to simulate that the notification has been sent.
	return args.Bool(0)
}

func TestChargeCustomer(t *testing.T) {
	smsService := new(smsServiceMock)

	/*What the SendChargeNotification returns when 100 is sent, in this case true
	to simulate a succesfull notification. */
	smsService.On("SendChargeNotification", 100).Return(true)

	//Define the service to be tested.
	myService := MyService{smsService}

	//Method call.
	myService.ChargeCustomer(100)

	//Verify the myService.ChargeCustomer called the mocked SendChargeNotification.
	smsService.AssertExpectations(t)

}

/* Test the vektra/mockery package which generates mocks automatically
   from the interfaces names. */

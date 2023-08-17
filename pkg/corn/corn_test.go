package cron

import (
	"fmt"
	"testing"
)

var i = 1

func TestNewCronService(t *testing.T) {
	service := NewCronService()
	err := service.AddFunc("@every 1s", A)
	service.Start()
	if err != nil {
		return
	}
	select {}
}

func A() {
	fmt.Println(i)
	i++
}

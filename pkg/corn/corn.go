// Package cron cron_service.go
package cron

import "github.com/robfig/cron/v3"

type ServiceOfCron struct {
	c *cron.Cron
}

func NewCronService() *ServiceOfCron {
	return &ServiceOfCron{
		c: cron.New(),
	}
}

func (c *ServiceOfCron) AddFunc(spec string, fn func()) error {
	_, err := c.c.AddFunc(spec, fn)
	return err
}

func (c *ServiceOfCron) Start() {
	c.c.Start()
}

func (c *ServiceOfCron) Stop() {
	c.c.Stop()
}

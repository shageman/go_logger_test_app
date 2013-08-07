package controllers

import (
	"github.com/robfig/revel"
	"fmt"
	"time"
	"strings"
)

var run bool
var numberOfBytesSent uint64
var numberOfMessagesSent uint64
var logged string

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) SleepLog(sleep time.Duration, logMessage string) revel.Result {
	run = true
	numberOfBytesSent = 0
	numberOfMessagesSent = 0
	logged = ""

	bytesSentPerMessage := uint64(len([]byte(logMessage)))

	fmt.Printf("Let's go. Waiting %v microseconds between loglines. Logging '%s' every time. That is %v bytes per message\n", sleep, logMessage, bytesSentPerMessage)
	for run {
		numberOfMessagesSent += 1
		numberOfBytesSent += bytesSentPerMessage
		logged = strings.Join([]string{logged, logMessage}, "")

		time.Sleep(time.Microsecond * sleep)
		fmt.Println(logMessage)
	}

	return c.Render(sleep)
}

func (c App) SleepLogCurrentTime(sleep time.Duration) revel.Result {
	run = true
	numberOfBytesSent = 0
	numberOfMessagesSent = 0
	logged = ""

	fmt.Printf("Let's go. Waiting %v microseconds between loglines. Logging current time every time.n", sleep)
	for run {
		logMessage := time.Now().Format(time.StampNano)

		bytesSentPerMessage := uint64(len([]byte(logMessage)))

		numberOfMessagesSent += 1
		numberOfBytesSent += bytesSentPerMessage
		logged = strings.Join([]string{logged, logMessage}, "")

		time.Sleep(time.Microsecond * sleep)
		fmt.Println(logMessage)
	}

	return c.Render(sleep)
}

func (c App) SleepStop() revel.Result {
	run = false

	return c.Render(numberOfBytesSent, numberOfMessagesSent, logged)
}

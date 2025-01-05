package core

import (
	"fmt"
	"time"
)

type Log string

type Logger interface {
	Info(msg string)
	Warn(msg string)
	Error(msg string)
	Fatal(msg string)
	Debug(msg string)
}

func (l Log) Info() {

	fmt.Printf("["+time.Now().Format(time.TimeOnly)+"][INFO] %v\n", l)
}

func (l Log) Warn() {

	fmt.Printf("["+time.Now().Format(time.TimeOnly)+"][WARN] %v\n", l)
}

func (l Log) Error() {
	fmt.Printf("["+time.Now().Format(time.TimeOnly)+"][ERROR] %v\n", l)
}

func (l Log) Fatal() {
	fmt.Printf("["+time.Now().Format(time.TimeOnly)+"][FATAL] %v\n", l)
}

func (l Log) Debug() {
	fmt.Printf("["+time.Now().Format(time.TimeOnly)+"][DEBUG] %v\n", l)
}

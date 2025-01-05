package main

import "logger/core"

func main() {
	var logMessage core.Log
	logMessage = "Hello, World!"
	differentLogMessage := core.Log("Hello, World!")

	logMessage.Info()
	differentLogMessage.Warn()
}

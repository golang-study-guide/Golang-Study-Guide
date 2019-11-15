package main

import (
	"fmt"
	"log"
	"log/syslog"
)

func main() {
	sysLog, _ := syslog.New(syslog.LOG_ALERT|syslog.LOG_SYSLOG, "Something really bad has happened!!!")

	log.SetOutput(sysLog)

	// on a CentOS vm, this sends a log entry to /var/log/messages with the message "Something really bad has happened!!!"
	log.Fatal(sysLog)
	fmt.Println("The above line will terminate this program before this line has a chance to get executed")
}

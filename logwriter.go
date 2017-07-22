package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {

	logDirectory := "/tmp/"

	time := time.Now()
	today := fmt.Sprintf("%04d%02d%02d", time.Year(), time.Month(), time.Day())

	logfile := logDirectory + "application.log." + today

	f, err := os.OpenFile(logfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	defer func() {
		f.Close()
	}()

	log.SetOutput(f)
	log.Print("test")
}

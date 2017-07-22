package main

import (
	"os"
)

func main() {

	logfile := "application.log.DATE"

	f, err := os.OpenFile(logfile, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	defer func() {
		f.Close()
	}()

	f.WriteString("test")
}

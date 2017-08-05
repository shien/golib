package main

import (
	"golib/src/logwriter"
)

func main() {

	conf := logwriter.LogConf{LogDirectory: "/tmp/", LogLevel: "info", Rotate: "daily"}
	content := logwriter.LogContent{LogType: "application", LogLevel: "info", Log: "test"}
	logwriter.WriteLog(conf, content)

}

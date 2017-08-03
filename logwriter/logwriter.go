package logwriter

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"time"
)

type LogContent struct {
	LogType  string
	LogLevel string
	Log      string
}

type LogConf struct {
	LogDirectory string
	LogLevel     string
	Rotate       string
}

func main() {

	conf := LogConf{LogDirectory: "/tmp/", LogLevel: "info", Rotate: "daily"}
	content := LogContent{LogType: "application", LogLevel: "info", Log: "test"}

	writeLog(conf, content)

}

func writeLog(conf LogConf, content LogContent) bool {

	time := time.Now()
	today := fmt.Sprintf("%04d%02d%02d", time.Year(), time.Month(), time.Day())

	logfile := conf.LogDirectory + content.LogType + ".log." + today

	f, err := os.OpenFile(logfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return false
	}

	defer func() {
		f.Close()
	}()

	log.SetOutput(f)
	log.Print("[" + content.LogLevel + "] " + content.Log)

	return true
}

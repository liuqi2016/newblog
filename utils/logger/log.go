package logger

import (
	"io"
	"log"
	"os"
)

var (
	info    *log.Logger
	warning *log.Logger
	error   *log.Logger
)
var (
	greenBg      = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	whiteBg      = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	yellowBg     = string([]byte{27, 91, 57, 48, 59, 52, 51, 109})
	redBg        = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	blueBg       = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	magentaBg    = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	cyanBg       = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	green        = string([]byte{27, 91, 51, 50, 109})
	white        = string([]byte{27, 91, 51, 55, 109})
	yellow       = string([]byte{27, 91, 51, 51, 109})
	red          = string([]byte{27, 91, 51, 49, 109})
	blue         = string([]byte{27, 91, 51, 52, 109})
	magenta      = string([]byte{27, 91, 51, 53, 109})
	cyan         = string([]byte{27, 91, 51, 54, 109})
	reset        = string([]byte{27, 91, 48, 109})
	disableColor = false
)

func init() {
	log.SetPrefix("【Blog】")
	log.SetFlags(log.Ldate | log.Lshortfile)

	errFile, err := os.OpenFile("errors.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("打开日志文件失败：", err)
	}
	info = log.New(os.Stdout, "Info:", log.Ldate|log.Ltime|log.Lshortfile)
	warning = log.New(os.Stdout, "Warning:", log.Ldate|log.Ltime|log.Lshortfile)
	error = log.New(io.MultiWriter(os.Stderr, errFile), "Error:", log.Ldate|log.Ltime|log.Lshortfile)
}

func Info(s string) {
	info.Println(greenBg, s, reset)
}
func Warning(s string) {
	warning.Println(yellowBg, s, reset)
}
func Error(s string) {
	error.Println(redBg, s, reset)
}

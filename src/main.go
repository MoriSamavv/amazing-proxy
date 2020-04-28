package main

import (
	"src/model/message/gfw_message"
	log "github.com/sirupsen/logrus"
)

func main() {
	a := new GfwMessage()
	log.Info(a)
	log.Info("hello world!")
}

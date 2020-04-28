package main

import (
	"github.com/amazing-proxy/src/model/message"
	log "github.com/sirupsen/logrus"
)

func main() {
	a := message.GfwMessage{"weimingliu"}
	log.Info(a)
	log.Info("hello world!")
}

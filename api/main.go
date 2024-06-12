package main

import (
	"time"

	lokihook "github.com/akkuman/logrus-loki-hook"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	lokiHookConfig := &lokihook.Config{
		URL:       "http://loki:3100/loki/api/v1/push",
		LevelName: "severity",
		Labels: map[string]string{
			"application": "test",
		},
	}
	hook, err := lokihook.NewHook(lokiHookConfig)
	if err != nil {
		log.Error(err)
	} else {
		log.AddHook(hook)
	}
}

func main() {
	log.Info("Starting app")
	for {
		log.Info("I'm a Test.")

		time.Sleep(5 * time.Second)
	}

}

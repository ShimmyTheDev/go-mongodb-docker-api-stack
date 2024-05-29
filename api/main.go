package main

import (
	"log/slog"
	"time"

	lokihook "github.com/akkuman/logrus-loki-hook"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	lokiHookConfig := &lokihook.Config{
		// the loki api url
		URL: "http://loki:3100/loki/api/v1/push",
		// (optional, default: severity) the label's key to distinguish log's level, it will be added to Labels map
		LevelName: "severity",
		// the labels which will be sent to loki, contains the {levelname: level}
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
	slog.Info("Starting app")
	log.Info("Starting app")
	for {
		log.Info("I'm a Test.")
		slog.Info("I'm a Test.")

		// Because the loki hook use the channel to send log in an asynchronous manner, We should wait for the log to be sent
		time.Sleep(5 * time.Second)
	}

}

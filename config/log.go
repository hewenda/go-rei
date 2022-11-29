package config

import (
	"github.com/orandin/lumberjackrus"
	log "github.com/sirupsen/logrus"
)

func init() {
	hook, err := lumberjackrus.NewHook(
		&lumberjackrus.LogFile{
			Filename:   "./log.txt",
			MaxSize:    10,
			MaxBackups: 1,
			MaxAge:     1,
			Compress:   false,
			LocalTime:  false,
		},
		log.InfoLevel,
		&log.TextFormatter{FullTimestamp: true},
		&lumberjackrus.LogFileOpts{},
	)

	if err != nil {
		panic(err)
	}

	log.AddHook(hook)
}

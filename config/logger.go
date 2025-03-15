package config

import (
	"os"

	enum "github.com/Vympel87/inventory-pos-be/internal/constants"
	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func InitLogger() {
	Log = logrus.New()

	level, err := logrus.ParseLevel(enum.LevelDebug)
	if err != nil {
		level = logrus.DebugLevel
	}
	Log.SetLevel(level)

	Log.SetOutput(os.Stdout)
	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}

// Optional helpers
func LogQuery(data interface{}) {
	Log.WithField("query", data).Info("Database query")
}

func LogError(err interface{}) {
	Log.WithField("error", err).Error("An error occurred")
}

func LogWarn(data interface{}) {
	Log.WithField("warn", data).Warn("Warning")
}

func LogInfo(data interface{}) {
	Log.WithField("info", data).Info("Info")
}

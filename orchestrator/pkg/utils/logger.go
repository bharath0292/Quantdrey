package util

import (
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

func InitLogger(logLevel zerolog.Level) {
	logDir := "./logs"

	if err := os.MkdirAll(logDir, 0755); err != nil {
		log.Error().Err(err).Msg("Failed to create log directory")
	}

	zerolog.SetGlobalLevel(logLevel)
	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		return filepath.Base(file) + ":" + strconv.Itoa(line)
	}

	multiWriter := zerolog.MultiLevelWriter(
		zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339},
		&lumberjack.Logger{
			Filename:   filepath.Join(logDir, "server.log"),
			MaxSize:    50,
			MaxBackups: 30,
			MaxAge:     90,
			Compress:   true,
		},
	)

	log.Logger = zerolog.New(multiWriter).
		With().
		Timestamp().
		Caller().
		Logger()

	log.Trace().Msg("Logger initialized")
}

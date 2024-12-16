package log

import (
	"component/log/config"
	"testing"
)

func TestStdoutLogger(t *testing.T) {
	conf := config.NewConfig(
		config.WithLogFormat(config.FmtFormat),
		config.WithLogLevel(config.InfoLevel),
		config.WithLogStdout(true),
	)

	zapLogger, err := NewZapLogger(conf)
	if err != nil {
		t.Fatal(err)
	}
	zapLogger.Info("this is an info msg")
}
func TestFileLogger(t *testing.T) {
	conf := config.NewConfig(
		config.WithLogFormat(config.FmtFormat),
		config.WithLogLevel(config.InfoLevel),
		config.WithLogStdout(false),
		config.WithFileConfig(&config.FileConfig{
			LogPath:           "testlog",
			LogFileName:       "test.log",
			LogFileMaxSize:    10,
			LogFileMaxBackups: 2,
			LogMaxAge:         30,
			LogCompress:       false,
		}),
	)
	zapLogger, err := NewZapLogger(conf)
	if err != nil {
		t.Fatal(err)
	}
	zapLogger.Info("this is an info msg")
}
func TestKafkaLogger(t *testing.T) {
	conf := config.NewConfig(
		config.WithLogFormat(config.FmtFormat),
		config.WithLogLevel(config.InfoLevel),
		config.WithLogStdout(false),
		config.WithKafkaConfig(&config.KafkaConfig{
			BrokersAddr: []string{"localhost:9092"},
			TopicName:   "testlog",
		}),
	)
	zapLogger, err := NewZapLogger(conf)
	if err != nil {
		t.Fatal(err)
	}
	zapLogger.Info("this is an info msg")
}

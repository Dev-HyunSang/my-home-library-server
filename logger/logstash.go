// 2024.03.05 - HyunSang Park <me@hyunsang.dev>
package logger

import (
	logstash_logger "github.com/KaranJagtiani/go-logstash"
	"github.com/dev-hyunsang/my-home-library-server/config"
	"strconv"
)

func Init() *logstash_logger.Logstash {
	parsePort, err := strconv.Atoi(config.GetEnv("LOGSTASH_PORT"))
	if err != nil {
		panic(err)
	}

	logger := logstash_logger.Init(config.GetEnv("LOGSTASH_HOSTNAME"),
		parsePort, "udp", 5)

	return logger
}

func Log(msg map[string]interface{}) {
	logger := Init()
	logger.Log(msg)
}

func Error(msg map[string]interface{}) {
	logger := Init()
	logger.Error(msg)
}
func Info(msg map[string]interface{}) {
	logger := Init()
	logger.Info(msg)
}

func Debug(msg map[string]interface{}) {
	logger := Init()
	logger.Debug(msg)
}

func Warn(msg map[string]interface{}) {
	logger := Init()
	logger.Warn(msg)
}

func LogString(msg string) {
	logger := Init()
	logger.LogString(msg)
}

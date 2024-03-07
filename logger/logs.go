// 2024.03.05 - HyunSang Park <me@hyunsang.dev>
package logger

import (
	logstash_logger "github.com/KaranJagtiani/go-logstash"
	"github.com/dev-hyunsang/my-home-library-server/config"
	"strconv"
	"time"
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

// 효율적으로 코드를 작성하기 위해서 본 패키지에서 처리할 수도록 개발함.
func Error(err error) {
	logger := Init()

	logger.Error(map[string]interface{}{
		"level":     "error",
		"message":   err,
		"timestamp": time.Now().Unix(),
	})
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

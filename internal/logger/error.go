package logger

import (
	"encoding/json"

	"go.uber.org/zap"
)

type Users struct{}

func Error(content string,data LogData){
	jsonData,_ :=json.Marshal(data)
	zapLogger.Error(content,zap.String("data",string(jsonData)))
}
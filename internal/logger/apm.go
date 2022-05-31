package logger

import(
	"encoding/json"
	"go.uber.org/zap"
)

func APM(fields []zap.Field,content string,data LogData){
	jsonData,_ :=json.Marshal(data)
	zapLogger.With(fields...).Info(content,zap.String("data",string(jsonData)))
}
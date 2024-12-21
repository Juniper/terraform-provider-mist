package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/apimatic/go-core-runtime/logger"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

type TFLogger struct {
	ctx context.Context
}

func NewTFlogger(ctx context.Context) *TFLogger {
	return &TFLogger{ctx: ctx}
}

func (l *TFLogger) Trace(msg string, args ...interface{}) {
	tflog.Trace(l.ctx, msg)
}

func (l *TFLogger) Debug(msg string, args ...interface{}) {
	tflog.Debug(l.ctx, msg)
}

func (l *TFLogger) Info(msg string, args ...interface{}) {
	tflog.Info(l.ctx, msg)
}

func (l *TFLogger) Warn(msg string, args ...interface{}) {
	tflog.Warn(l.ctx, msg)
}

func (l *TFLogger) Error(msg string, args ...interface{}) {
	tflog.Error(l.ctx, msg)
}

func (l *TFLogger) Log(level logger.Level, message string, metadata map[string]interface{}) {

	formattedMessage := formatLog(message, metadata)

	switch level {
	case logger.Level_INFO:
		tflog.Info(l.ctx, formattedMessage)
	case logger.Level_TRACE:
		tflog.Trace(l.ctx, formattedMessage)
	case logger.Level_DEBUG:
		tflog.Debug(l.ctx, formattedMessage)
	case logger.Level_WARN:
		tflog.Warn(l.ctx, formattedMessage)
	case logger.Level_ERROR:
		tflog.Error(l.ctx, formattedMessage)
	default:
		tflog.Debug(l.ctx, formattedMessage)
	}
}

func formatLog(message string, metadata map[string]interface{}) string {
	var logBuilder strings.Builder

	logBuilder.WriteString("SDK API Log: ")
	logBuilder.WriteString(message)

	if len(metadata) > 0 {
		logBuilder.WriteString("\n")
		for key, value := range metadata {

			strValue := fmt.Sprintf("%v", value)
			if isJSON(strValue) {
				strValue = formatJSON(strValue)
			}
			logBuilder.WriteString(fmt.Sprintf("  %s: %s\n", key, strValue))
		}
	}

	return logBuilder.String()
}

func isJSON(s string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(s), &js) == nil
}

func formatJSON(jsonStr string) string {
	var formattedJSON bytes.Buffer
	if err := json.Indent(&formattedJSON, []byte(jsonStr), "", "  "); err != nil {
		return jsonStr
	}
	return formattedJSON.String()
}

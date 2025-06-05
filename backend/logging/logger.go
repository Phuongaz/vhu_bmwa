package logging

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	LogDir = "logs"
)

type LogType string

const (
	AccessLog   LogType = "access"
	SecurityLog LogType = "security"
	ErrorLog    LogType = "error"
)

var loggers map[LogType]*logrus.Logger

func init() {
	loggers = make(map[LogType]*logrus.Logger)
	setupLoggers()
}

func setupLoggers() {
	if err := os.MkdirAll(LogDir, 0755); err != nil {
		panic(fmt.Sprintf("Failed to create logs directory: %v", err))
	}

	setupLogger(AccessLog)
	setupLogger(SecurityLog)
	setupLogger(ErrorLog)
}

func setupLogger(logType LogType) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	today := time.Now().Format("2006-01-02")
	dailyDir := filepath.Join(LogDir, today)
	if err := os.MkdirAll(dailyDir, 0755); err != nil {
		panic(fmt.Sprintf("Failed to create daily log directory: %v", err))
	}

	logPath := filepath.Join(dailyDir, fmt.Sprintf("%s.log", logType))
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(fmt.Sprintf("Failed to open log file: %v", err))
	}

	logger.SetOutput(file)
	loggers[logType] = logger
}

func getLogger(logType LogType) *logrus.Logger {
	today := time.Now().Format("2006-01-02")
	currentLogPath := filepath.Join(LogDir, today, fmt.Sprintf("%s.log", logType))
	if _, err := os.Stat(currentLogPath); os.IsNotExist(err) {
		setupLogger(logType)
	}
	return loggers[logType]
}

type AccessLogEntry struct {
	Method      string
	Path        string
	StatusCode  int
	Duration    time.Duration
	IP          string
	UserAgent   string
	UserID      uint
	RequestID   string
	QueryParams map[string]string
}

type SecurityLogEntry struct {
	EventType   string
	UserID      uint
	IP          string
	Description string
	Metadata    map[string]interface{}
}

func LogAccess(entry AccessLogEntry) {
	logger := getLogger(AccessLog)
	logger.WithFields(logrus.Fields{
		"method":      entry.Method,
		"path":        entry.Path,
		"status_code": entry.StatusCode,
		"duration":    entry.Duration.String(),
		"ip":          entry.IP,
		"user_agent":  entry.UserAgent,
		"user_id":     entry.UserID,
		"request_id":  entry.RequestID,
		"query":       entry.QueryParams,
		"timestamp":   time.Now().Format("2006-01-02 15:04:05"),
	}).Info("Access Log")
}

func LogSecurity(entry SecurityLogEntry) {
	logger := getLogger(SecurityLog)
	logger.WithFields(logrus.Fields{
		"event_type":  entry.EventType,
		"user_id":     entry.UserID,
		"ip":          entry.IP,
		"description": entry.Description,
		"metadata":    entry.Metadata,
		"timestamp":   time.Now().Format("2006-01-02 15:04:05"),
	}).Info("Security Log")
}

func LogError(err error, metadata map[string]interface{}) {
	logger := getLogger(ErrorLog)
	logger.WithFields(logrus.Fields{
		"error":     err.Error(),
		"metadata":  metadata,
		"timestamp": time.Now().Format("2006-01-02 15:04:05"),
	}).Error("Error Log")
}

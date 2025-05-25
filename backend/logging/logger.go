package logging

import (
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	AccessLog   *logrus.Logger
	SecurityLog *logrus.Logger
	ErrorLog    *logrus.Logger
)

func init() {
	if err := os.MkdirAll("logs", 0755); err != nil {
		panic("Failed to create logs directory: " + err.Error())
	}

	AccessLog = logrus.New()
	AccessLog.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})
	accessLogPath := filepath.Join("logs", "access.log")
	accessLogFile, err := os.OpenFile(accessLogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic("Failed to open access log file: " + err.Error())
	}
	AccessLog.SetOutput(accessLogFile)

	SecurityLog = logrus.New()
	SecurityLog.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})
	securityLogPath := filepath.Join("logs", "security.log")
	securityLogFile, err := os.OpenFile(securityLogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic("Failed to open security log file: " + err.Error())
	}
	SecurityLog.SetOutput(securityLogFile)

	ErrorLog = logrus.New()
	ErrorLog.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})
	errorLogPath := filepath.Join("logs", "error.log")
	errorLogFile, err := os.OpenFile(errorLogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic("Failed to open error log file: " + err.Error())
	}
	ErrorLog.SetOutput(errorLogFile)
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

type ErrorLogEntry struct {
	Error      error
	Component  string
	UserID     uint
	RequestID  string
	StackTrace string
	Context    map[string]interface{}
}

func LogAccess(entry AccessLogEntry) {
	AccessLog.WithFields(logrus.Fields{
		"method":      entry.Method,
		"path":        entry.Path,
		"status_code": entry.StatusCode,
		"duration_ms": entry.Duration.Milliseconds(),
		"ip":          entry.IP,
		"user_agent":  entry.UserAgent,
		"user_id":     entry.UserID,
		"request_id":  entry.RequestID,
		"query":       entry.QueryParams,
	}).Info("Access log")
}

func LogSecurity(entry SecurityLogEntry) {
	SecurityLog.WithFields(logrus.Fields{
		"event_type":  entry.EventType,
		"user_id":     entry.UserID,
		"ip":          entry.IP,
		"description": entry.Description,
		"metadata":    entry.Metadata,
	}).Info("Security event")
}

func LogError(entry ErrorLogEntry) {
	ErrorLog.WithFields(logrus.Fields{
		"error":      entry.Error.Error(),
		"component":  entry.Component,
		"user_id":    entry.UserID,
		"request_id": entry.RequestID,
		"stack":      entry.StackTrace,
		"context":    entry.Context,
	}).Error("System error")
}

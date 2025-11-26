package shared

import (
	"os"
	"runtime"

	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	"go.uber.org/zap"
)

var log *MyLogger
var LoggerExternal *MyLogger

type Logger interface {
	Debug(msg string, keyvals ...interface{})
	Info(msg string, keyvals ...interface{})
	Error(msg string, keyvals ...interface{})
	Warn(msg string, keyvals ...interface{})

	// With(keyvals ...interface{}) Logger
}

type MyLogger struct {
	logger *zap.Logger
}

func NewMyLogger(Logger *zap.Logger) *MyLogger {
	return &MyLogger{logger: Logger}
}

func init() {

	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "logs.log",
		MaxSize:    10, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
	})

	config := zap.NewProductionEncoderConfig()

	config.TimeKey = "at"
	config.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(config),
		w,
		zap.DebugLevel,
	)

	// combine log and output to screen
	coreCombined := zapcore.NewTee(
		core,
		zapcore.NewCore(zapcore.NewJSONEncoder(config), zapcore.AddSync(os.Stdout), zap.DebugLevel),
	)

	logger := zap.New(coreCombined)

	log = NewMyLogger(logger)
	LoggerExternal = log

}

func (s *MyLogger) Debug(msg string, vals ...interface{}) {

	sugar := s.logger.Sugar()
	sugar.Debug(msg, vals)
}
func (s *MyLogger) Debugw(msg string, vals ...interface{}) {

	sugar := s.logger.Sugar()
	sugar.Debugw(msg, vals...)
}
func (s *MyLogger) Info(msg string, vals ...interface{}) {

	sugar := s.logger.Sugar()
	sugar.Info(msg, vals)
}
func (s *MyLogger) Infow(msg string, vals ...interface{}) {

	sugar := s.logger.Sugar()
	sugar.Infow(msg, vals...)
}
func (s *MyLogger) Error(msg string, vals ...interface{}) {

	sugar := s.logger.Sugar()
	sugar.Error(msg, vals)
}
func (s *MyLogger) Errorw(msg string, vals ...interface{}) {

	sugar := s.logger.Sugar()
	sugar.Errorw(msg, vals...)
}
func (s *MyLogger) Warn(msg string, vals ...interface{}) {

	sugar := s.logger.Sugar()
	sugar.Warn(msg, vals)
}
func (s *MyLogger) Warnw(msg string, vals ...interface{}) {

	sugar := s.logger.Sugar()
	sugar.Warnw(msg, vals...)
}
func GetRunnerFields() []interface{} {

	_, file, line, _ := runtime.Caller(2)

	return []interface{}{
		"file", file,
		"fileLine", line,
	}

}

func GetContextFields(context RequestContext) []interface{} {

	return []interface{}{
		"userId", context.UserId,
		"userGroupId", context.UserGroupId,
		"organizationId", context.OrganizationId,
		"hostName", context.HostName,
		"userName", context.UserName,
		"userIp", context.UserIp,
		"accessedPath", context.AccessedPath,
		"accessedMethod", context.AccessedMethod,
	}

}

// Contextual info
func InfoCw(context RequestContext, msg string, vals ...interface{}) {
	vals = append(vals, GetContextFields(context)...)
	vals = append(vals, GetRunnerFields()...)
	log.Infow(msg, vals...)
}
func Info(msg string, vals ...interface{}) {
	log.Info(msg, vals...)
}
func Infow(msg string, vals ...interface{}) {
	vals = append(vals, GetRunnerFields()...)
	log.Infow(msg, vals...)
}

// Contextual debug
func DebugCw(context RequestContext, msg string, vals ...interface{}) {
	vals = append(vals, GetContextFields(context)...)
	vals = append(vals, GetRunnerFields()...)
	log.Debugw(msg, vals...)
}
func Debug(msg string, vals ...interface{}) {
	log.Debug(msg, vals...)
}
func Debugw(msg string, vals ...interface{}) {
	vals = append(vals, GetRunnerFields()...)
	log.Debugw(msg, vals...)
}

// Contextual error
func ErrorCw(context RequestContext, msg string, vals ...interface{}) {
	vals = append(vals, GetContextFields(context)...)
	vals = append(vals, GetRunnerFields()...)
	log.Errorw(msg, vals...)
}
func Error(msg string, vals ...interface{}) {
	log.Error(msg, vals...)
}
func Errorw(msg string, vals ...interface{}) {
	vals = append(vals, GetRunnerFields()...)
	log.Errorw(msg, vals...)
}

// Contextual warn
func WarnCw(context RequestContext, msg string, vals ...interface{}) {
	vals = append(vals, GetContextFields(context)...)
	vals = append(vals, GetRunnerFields()...)
	log.Warnw(msg, vals...)
}
func Warn(msg string, vals ...interface{}) {
	log.Warn(msg, vals...)
}
func Warnw(msg string, vals ...interface{}) {
	vals = append(vals, GetRunnerFields()...)
	log.Warnw(msg, vals...)
}

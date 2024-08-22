package log

import (
	"fmt"
	"os"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func init() {
	initLogger()
}

func initLogger() {
	customTimeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString("[" + t.Format("2006-01-02 15:04:05.000") + "]")
	}

	customLevelEncoder := func(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString("[" + level.CapitalString() + "]")
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "line",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeTime:     customTimeEncoder,
		EncodeLevel:    customLevelEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.AddSync(os.Stdout),
		zapcore.DebugLevel,
	)

	Logger = zap.New(core)
	Logger = Logger.WithOptions(zap.AddCaller(), zap.AddCallerSkip(1))
}

func Debug(message string) {
	Logger.Debug(message)
}

func Info(message string) {
	Logger.Info(message)
}

func Debugf(format string, a ...any) {
	Logger.Debug(fmt.Sprintf(format, a...))
}

func Infof(format string, a ...any) {
	Logger.Info(fmt.Sprintf(format, a...))
}

func Warnf(format string, a ...any) {
	Logger.Warn(fmt.Sprintf(format, a...))
}

func Errorf(format string, a ...any) {
	Logger.Error(fmt.Sprintf(format, a...))
}

func ProfileLogf(start time.Time, format string, a ...any) {
	elapsed := time.Since(start)
	message := strings.Join([]string{
		fmt.Sprintf(format, a...),
		fmt.Sprintf("time elapsed: %s", elapsed)},
		" | ",
	)
	Logger.Info(message)
}

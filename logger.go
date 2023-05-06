package zaprotatelogger

import (
	"os"
	"path/filepath"

	"github.com/pnkj-kmr/zap-rotate-logger/internal/option"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// New func returns a zap logger instance
// along with the lumberjack file log rotator
func New(options ...option.Option) *zap.Logger {
	// defining default vaiables
	var f string = "app"
	maxSize, maxAge, maxBkp := 1, 30, 30
	var compress, debug bool

	// filtering the options variables for logger
	for _, o := range options {
		switch o.Name() {
		case optFileName:
			f = o.Value().(string)
		case optMaxSize:
			maxSize = o.Value().(int)
		case optMaxAge:
			maxAge = o.Value().(int)
		case optMaxBackups:
			maxBkp = o.Value().(int)
		case optCompress:
			compress = o.Value().(bool)
		case optDebugEnabled:
			debug = o.Value().(bool)
		}
	}

	// initiating the file logger
	rotator := &lumberjack.Logger{
		Filename:   filepath.Join("./log", f+".log"),
		MaxSize:    maxSize, // megabytes
		MaxBackups: maxBkp,
		MaxAge:     maxAge,   //days
		Compress:   compress, // disabled by default
	}

	// zap prod logger config
	config := zap.NewProductionEncoderConfig()
	// time binding
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	// new file encoder config
	fileEncoder := zapcore.NewJSONEncoder(config)

	// zap file logger binding
	writer := zapcore.AddSync(rotator)

	// setting the log level
	_level := zapcore.InfoLevel
	if debug {
		_level = zapcore.DebugLevel
	}

	// defining the zap logger instance
	var logger *zap.Logger
	var core zapcore.Core
	// pre-check - enabling the console on debug true
	if _level == zap.DebugLevel {
		consoleEncoder := zapcore.NewConsoleEncoder(config)
		core = zapcore.NewTee(
			zapcore.NewCore(fileEncoder, writer, _level),
			zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), _level),
		)
		logger = zap.New(core, zap.Development(), zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	} else {
		core = zapcore.NewTee(
			zapcore.NewCore(fileEncoder, writer, _level),
		)
		logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	}
	return logger
}

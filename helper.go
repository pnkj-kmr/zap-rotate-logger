package zaprotatelogger

import "github.com/pnkj-kmr/zap-rotate-logger/internal/option"

const (
	optFileName     = "file_name"
	optMaxSize      = "max_size"
	optMaxAge       = "max_age"
	optMaxBackups   = "max_backups"
	optCompress     = "compress"
	optDebugEnabled = "debug_mode"
)

// WithFileName is the file to write logs to. Backup log files will be retained
// in the same directory.  It uses <processname>-lumberjack.log in
// os.TempDir() if empty.
func WithFileName(s string) option.Option {
	return option.New(optFileName, s)
}

// WithMaxSize is the maximum size in megabytes of the log file before it gets
// rotated. It defaults to 100 megabytes.
func WithMaxSize(s int) option.Option {
	return option.New(optMaxSize, s)
}

// WithMaxAge is the maximum number of days to retain old log files based on the
// timestamp encoded in their filename.  Note that a day is defined as 24
// hours and may not exactly correspond to calendar days due to daylight
// savings, leap seconds, etc. The default is not to remove old log files
// based on age.
func WithMaxAge(s int) option.Option {
	return option.New(optMaxAge, s)
}

// WithMaxBackups is the maximum number of old log files to retain.  The default
// is to retain all old log files (though MaxAge may still cause them to get
// deleted.)
func WithMaxBackups(s int) option.Option {
	return option.New(optMaxBackups, s)
}

// WithCompress determines if the rotated log files should be compressed
// using gzip. The default is not to perform compression.
func WithCompress(s bool) option.Option {
	return option.New(optCompress, s)
}

// WithDebug application log devel
func WithDebug(s bool) option.Option {
	return option.New(optDebugEnabled, s)
}

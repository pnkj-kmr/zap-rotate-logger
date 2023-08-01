# zap-rotate-logger

Provide an [zap](https://github.com/uber-go/zap) logger instance with [lumberjack](https://github.com/natefinch/lumberjack) which helps to auto rotate file once it's size breached.

## HOW TO USE

---

```
package main

import (
	zrl "github.com/pnkj-kmr/zap-rotate-logger"
)

func main() {
	logger := zrl.New(
		zrl.WithFileName("app"),
		zrl.WithMaxSize(1),
		zrl.WithMaxAge(60),
		zrl.WithMaxBackups(30),
		zrl.WithCompress(true),
		zrl.WithDebug(true),
	)

	for i := 0; i < 100; i++ {
		logger.Info("testing...")
	}
}

```

## DESCRIPTION

---

When you integrate this to into your app, it automatically write to logs into current working directory (path-to-app/log/_xyz_.log) and rotate the new log file as you configured: No-more file rotation, No-more disk size full. To setup the zap-roate-logger!

To install, simply do a `go get`:

```
go install github.com/pnkj-kmr/zap-rotate-logger
```

## OPTIONS

---

### `WithFileName` (DEFAULT: app)

The WithFileName used to define log file name. For example:

```
  zrl.WithFileName("app_name")
```

### `WithMaxSize` (DEFAULT: 1)

The WithMaxSize is the maximum size in megabytes of the log file before it gets rotated. It defaults to 1 megabytes.

```
  zrl.WithMaxSize(1)
```

### `WithMaxAge` (DEFAULT: 30)

The WithMaxAge is the maximum number of days to retain old log files based on the timestamp encoded in their filename. Note that a day is defined as 24 hours and may not exactly correspond to calendar days due to daylight savings, leap seconds, etc. The default is not to remove old log files based on age.

```
  zrl.WithMaxAge(30)
```

### `WithMaxBackups` (DEFAULT: 30)

The WithMaxBackups is the maximum number of old log files to retain. The default is to retain all old log files (though MaxAge may still cause them to get deleted.)

```
  zrl.WithMaxBackups(30)
```

### `WithCompress` (DEFAULT: false)

The WithCompress determines if the rotated log files should be compressed using gzip.

```
  zrl.WithCompress(true)
```

### `WithDebug` (DEFAULT: false)

The WithDebug is used for application log devel. For example:

```
  zrl.WithDebug(true)
```

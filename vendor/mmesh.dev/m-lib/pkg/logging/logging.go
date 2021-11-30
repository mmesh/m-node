package logging

import (
	"fmt"

	"mmesh.dev/m-lib/pkg/xlog"
	"x6a.dev/pkg/msg"
)

var Interactive bool

var LogLevel xlog.LogLevel

func Trace(args ...interface{}) {
	if LogLevel > xlog.TRACE {
		return
	}

	if Interactive {
		msg.Trace(args...)
	} else {
		xlog.Trace(args...)
	}
}

func Debug(args ...interface{}) {
	if LogLevel > xlog.DEBUG {
		return
	}

	if Interactive {
		msg.Debug(args...)
	} else {
		xlog.Debug(args...)
	}
}

func Info(args ...interface{}) {
	if LogLevel > xlog.INFO {
		return
	}

	if Interactive {
		msg.Info(args...)
	} else {
		xlog.Info(args...)
	}
}

func Warn(args ...interface{}) {
	if LogLevel > xlog.WARN {
		return
	}

	if Interactive {
		msg.Warn(args...)
	} else {
		xlog.Warn(args...)
	}
}

func Error(args ...interface{}) {
	if Interactive {
		msg.Error(args...)
	} else {
		xlog.Error(args...)
	}
}

func Alert(args ...interface{}) {
	if Interactive {
		msg.Alert(args...)
	} else {
		xlog.Alert(args...)
	}
}

func Ok(args ...interface{}) {
	if Interactive {
		msg.Ok(args...)
	} else {
		xlog.Info(args...)
	}
}

func Fail(args ...interface{}) {
	if Interactive {
		msg.Fail(args...)
	} else {
		xlog.Error(args...)
	}
}

func Println(args ...interface{}) {
	if Interactive {
		fmt.Println(args...)
	} else {
		xlog.Info(args...)
	}
}

func Tracef(format string, args ...interface{}) {
	if LogLevel > xlog.TRACE {
		return
	}

	if Interactive {
		msg.Tracef(format, args...)
	} else {
		xlog.Tracef(format, args...)
	}
}

func Debugf(format string, args ...interface{}) {
	if LogLevel > xlog.DEBUG {
		return
	}

	if Interactive {
		msg.Debugf(format, args...)
	} else {
		xlog.Debugf(format, args...)
	}
}

func Infof(format string, args ...interface{}) {
	if LogLevel > xlog.INFO {
		return
	}

	if Interactive {
		msg.Infof(format, args...)
	} else {
		xlog.Infof(format, args...)
	}
}

func Warnf(format string, args ...interface{}) {
	if LogLevel > xlog.WARN {
		return
	}

	if Interactive {
		msg.Warnf(format, args...)
	} else {
		xlog.Warnf(format, args...)
	}
}

func Errorf(format string, args ...interface{}) {
	if Interactive {
		msg.Errorf(format, args...)
	} else {
		xlog.Errorf(format, args...)
	}
}

func Alertf(format string, args ...interface{}) {
	if Interactive {
		msg.Alertf(format, args...)
	} else {
		xlog.Alertf(format, args...)
	}
}

func Okf(format string, args ...interface{}) {
	if Interactive {
		msg.Okf(format, args...)
	} else {
		xlog.Infof(format, args...)
	}
}

func Failf(format string, args ...interface{}) {
	if Interactive {
		msg.Failf(format, args...)
	} else {
		xlog.Errorf(format, args...)
	}
}

func Printf(format string, args ...interface{}) {
	if Interactive {
		fmt.Printf(format, args...)
	} else {
		xlog.Infof(format, args...)
	}
}

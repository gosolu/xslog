package xslog

import "github.com/gosolu/xslog/stacktrace"

// Stack return current call stack
func Stack() string {
	return stacktrace.Take(1)
}

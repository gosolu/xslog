package xslog

import (
	"fmt"
	"log/slog"
	"path/filepath"
	"strings"
)

// SimpleSource make log source shorter
func SimpleSource(group []string, attr slog.Attr) slog.Attr {
	key := attr.Key
	if key != slog.SourceKey {
		return attr
	}
	src, ok := attr.Value.Any().(*slog.Source)
	if !ok {
		return attr
	}
	attr.Value = slog.StringValue(fmt.Sprintf("%s:%d", filepath.Base(src.File), src.Line))
	return attr
}

// LowerLevel convert log level to lower
func LowerLevel(group []string, attr slog.Attr) slog.Attr {
	if attr.Key != slog.LevelKey {
		return attr
	}
	attr.Value = slog.StringValue(strings.ToLower(attr.Value.String()))
	return attr
}

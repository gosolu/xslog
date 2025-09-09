// Package xslog contains some utilities for log/slog
package xslog

import (
	"context"
	"log/slog"
)

var _ slog.Handler = new(contextHandler)

type contextHandler struct {
	slog.Handler
}

// UseContext create a new slog Handler with context support.
func UseContext(handler slog.Handler) slog.Handler {
	return &contextHandler{handler}
}

type ctxAttrKey struct{}
type ctxLevelKey struct{}

// WithLogLevel set current context log level
func WithLogLevel(ctx context.Context, level slog.Level) context.Context {
	return context.WithValue(ctx, ctxLevelKey{}, level)
}

func (h *contextHandler) Enabled(ctx context.Context, level slog.Level) bool {
	ctxLevel := ctx.Value(ctxLevelKey{})
	if ctxLevel != nil {
		if lvl, ok := ctxLevel.(slog.Level); ok {
			return level >= lvl
		}
	}
	return h.Handler.Enabled(ctx, level)
}

func (h *contextHandler) Handle(ctx context.Context, record slog.Record) error {
	if val := ctx.Value(ctxAttrKey{}); val != nil {
		if attrs, ok := val.([]slog.Attr); ok {
			record.AddAttrs(attrs...)
		}
	}
	return h.Handler.Handle(ctx, record)
}

func (h *contextHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	h2 := *h
	h2.Handler = h.Handler.WithAttrs(attrs)
	return &h2
}

func (h *contextHandler) WithGroup(name string) slog.Handler {
	h2 := *h
	h2.Handler = h.Handler.WithGroup(name)
	return &h2
}

// ContextAppendAttrs append attributes into context and create a new context
func ContextAppendAttrs(ctx context.Context, attrs ...slog.Attr) context.Context {
	var sas []slog.Attr
	if val := ctx.Value(ctxAttrKey{}); val != nil {
		if vs, ok := val.([]slog.Attr); ok {
			sas = vs
		}
	}
	if sas == nil {
		sas = make([]slog.Attr, 0, len(attrs))
	}
	sas = append(sas, attrs...)
	return context.WithValue(ctx, ctxAttrKey{}, sas)
}

type replaceFn func(group []string, attr slog.Attr) slog.Attr

// ReplaceAttr bunch a group of replace functions into a single ReplaceAttr function
func ReplaceAttr(functions ...replaceFn) replaceFn {
	return func(group []string, attr slog.Attr) slog.Attr {
		for _, fn := range functions {
			attr = fn(group, attr)
		}
		return attr
	}
}

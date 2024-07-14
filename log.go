package xslog

import (
	"context"
	"log/slog"
)

var _ slog.Handler = new(contextHandler)

type contextHandler struct {
	slog.Handler
}

func NewHandler(handler slog.Handler) slog.Handler {
	return &contextHandler{handler}
}

type ctxKeyType struct{}

var ctxAttrKey ctxKeyType

func (h *contextHandler) Handle(ctx context.Context, record slog.Record) error {
	if val := ctx.Value(ctxAttrKey); val != nil {
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

// AppendAttrs append attributes into context and create a new context
func AppendAttrs(ctx context.Context, attrs ...slog.Attr) context.Context {
	var sas []slog.Attr
	if val := ctx.Value(ctxAttrKey); val != nil {
		if vs, ok := val.([]slog.Attr); ok {
			sas = vs
		}
	}
	if sas == nil {
		sas = make([]slog.Attr, 0, len(attrs))
	}
	sas = append(sas, attrs...)
	return context.WithValue(ctx, ctxAttrKey, sas)
}

type replaceFn func(group []string, attr slog.Attr) slog.Attr

// AttrReplaces bunch a group of replace functions into a single ReplaceAttr function
func AttrReplaces(functions ...replaceFn) replaceFn {
	return func(group []string, attr slog.Attr) slog.Attr {
		for _, fn := range functions {
			attr = fn(group, attr)
		}
		return attr
	}
}

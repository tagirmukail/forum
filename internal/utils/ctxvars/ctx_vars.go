package ctxvars

import (
	"context"

	"github.com/google/uuid"
)

type Key uint

const (
	_ Key = iota
	requestIDKey
)

func WithRequestID(ctx context.Context, id string) context.Context {
	if id == "" {
		id = uuid.New().String()
	}

	return context.WithValue(ctx, requestIDKey, id)
}

func GetRequestID(ctx context.Context) string {
	rID := ctx.Value(requestIDKey)

	requestID, ok := rID.(string)
	if !ok {
		return ""
	}

	return requestID
}

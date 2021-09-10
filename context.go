package gitlab

import "context"

type key uint8

const optKey key = iota

func WithToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, optKey, token)
}

func TokenFromContext(ctx context.Context) *string {
	val := ctx.Value(optKey)
	if val == nil {
		return nil
	}

	opt, ok := val.(string)
	if !ok {
		return nil
	}

	return &opt
}

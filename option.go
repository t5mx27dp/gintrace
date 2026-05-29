package gintrace

type Option func(*config)

func WithTraceIDKey(key string) Option {
	return func(c *config) {
		c.traceIDKey = key
	}
}

func WithHeaderKey(key string) Option {
	return func(c *config) {
		c.headerKey = key
	}
}

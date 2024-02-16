package kafka

// Middleware is a chainable behavior modifier for the [Handler] type.
type Middleware func(Handler) Handler

// Compose is a helper function for composing [Middleware].
// The first [Middleware] is treated as the outermost.
// The execution will be done in the declared order.
func Compose(mw ...Middleware) Middleware {
	return func(next Handler) Handler {
		for i := len(mw) - 1; i >= 0; i-- {
			next = mw[i](next)
		}

		return next
	}
}

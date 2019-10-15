package main

import "time"

func main() {

}

// --------------------------

type options struct {
	timeout time.Duration

	caching bool
}

// Option 重写 Connect.

type Option interface {
	apply(*options)
}
type optionFunc func(*options)

func (f optionFunc) apply(o *options) {
	f(o)
}
func WithTimeout(t time.Duration) Option {
	return optionFunc(func(o *options) {
		o.timeout = t
	})
}
func WithCaching(cache bool) Option {
	return optionFunc(func(o *options) {
		o.caching = cache
	})
}

// ##
// Connect 创建一个 connection

func Connect(
	addr string,
	opts ...Option,
) (someComponentName interface{} /* *Connection, */, e error) {
	options := options{
		timeout: defaultTimeout,
		caching: defaultCaching,
	}

	for _, o := range opts {
		// o(&options)
		o.apply(&options) // FIXME 这里才是对的？  原版的是上面的用法？
	}

	// ...
	return nil, nil
}

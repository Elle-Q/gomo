package file

import (
	"context"
	"gomo/common/encoder"
)
type Options struct {
	//encoder
	Encoder encoder.Encoder

	//for alternative data
	Context context.Context
}

type Option func(o *Options)

type filePathKey struct{}

// WithPath sets the path to file
func WithPath(p string) Option {
	return func(o *Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, filePathKey{}, p)
	}
}

func NewOptions(opts ...Option) Options {
	options := Options{
		//Encoder: json.NewEncoder()
		Context: context.Background(),
	}

	for _, o := range opts {
		o(&options)
	}

	return options
}

func WithEncoder(e encoder.Encoder) Option {
	return func(o *Options) {
		o.Encoder = e
	}
}

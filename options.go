package requestid

// Option for queue system
type Option func(*config)

type Generator func() string

// WithGenerator set fenerator function
func WithGenerator(g Generator) Option {
	return func(cfg *config) {
		cfg.generator = g
	}
}

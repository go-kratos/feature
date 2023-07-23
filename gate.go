package feature

import "flag"

var globalRegistry = NewRegistry()

func init() {
	flag.CommandLine.Var(globalRegistry, "feature-gates", "A set of key=value pairs that describe feature gates for alpha/beta/stable features.")
}

// Register register a feature gate.
func Register(name string, enabled bool, opts ...Option) (*Feature, error) {
	return globalRegistry.Register(name, enabled, opts...)
}

// MustRegister must register a feature gate.
func MustRegister(name string, enabled bool, opts ...Option) *Feature {
	return globalRegistry.MustRegister(name, enabled, opts...)
}

// Set parses the feature flags: foo=true,bar=false.
func Set(featureFlags string) error {
	return globalRegistry.Set(featureFlags)
}

// SetEnabled set feature enabled.
func SetEnabled(name string, enabled bool) error {
	return globalRegistry.SetEnabled(name, enabled)
}

// Visit visits all the features.
func Visit(f func(*Feature)) {
	globalRegistry.Visit(f)
}

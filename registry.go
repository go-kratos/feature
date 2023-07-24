package feature

import (
	"fmt"
	"strconv"
	"strings"
)

// Option is feature option.
type Option func(*Feature)

// WithFeatureStage with the feature stage.
func WithFeatureStage(stage Stage) Option {
	return func(f *Feature) {
		f.stage = stage
	}
}

// WithFeatureDescription with the Feature description.
func WithFeatureDescription(description string) Option {
	return func(f *Feature) {
		f.description = description
	}
}

// WithFeatureToVersion with the Feature to version.
func WithFeatureToVersion(version string) Option {
	return func(f *Feature) {
		f.toVersion = version
	}
}

// WithFeatureFromVersion with the Feature to version.
func WithFeatureFromVersion(version string) Option {
	return func(f *Feature) {
		f.toVersion = version
	}
}

// Registry is feature gates registry.
type Registry struct {
	features map[string]*Feature
}

// NewRegistry new a feature registry.
func NewRegistry() *Registry {
	return &Registry{
		features: make(map[string]*Feature),
	}
}

// Register register a feature gate.
func (r *Registry) Register(name string, enabled bool, opts ...Option) (*Feature, error) {
	if _, ok := r.features[name]; ok {
		return nil, fmt.Errorf("feature gate %s is registered", name)
	}
	feature := &Feature{name: name}
	feature.enabled.Store(enabled)
	for _, o := range opts {
		o(feature)
	}
	r.features[feature.Name()] = feature
	return feature, nil
}

// MustRegister must register a feature gate.
func (r *Registry) MustRegister(name string, enabled bool, opts ...Option) *Feature {
	feature, err := r.Register(name, enabled, opts...)
	if err != nil {
		panic(err)
	}
	return feature
}

// Visit visits all the features.
func (r *Registry) Visit(f func(*Feature)) {
	for _, feature := range r.features {
		f(feature)
	}
}

// SetEnabled set feature enabled.
func (r *Registry) SetEnabled(name string, enabled bool) error {
	f, ok := r.features[name]
	if !ok {
		return fmt.Errorf("not found feature: %s", name)
	}
	f.enabled.Store(enabled)
	return nil
}

// Set sets the feature arguments.
// eg: foo=true,bar=false
func (r *Registry) Set(args string) error {
	fs := strings.Split(args, ",")
	for _, s := range fs {
		feature := strings.Split(s, "=")
		name := feature[0]
		enabled, err := strconv.ParseBool(feature[1])
		if err != nil {
			return err
		}
		if err := r.SetEnabled(name, enabled); err != nil {
			return err
		}
	}
	return nil
}

func (r *Registry) String() string {
	pairs := []string{}
	for name, feature := range r.features {
		enabled := feature.enabled.Load()
		pairs = append(pairs, fmt.Sprintf("%s=%t", name, enabled))
	}
	return strings.Join(pairs, ",")

}

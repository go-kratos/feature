package feature

import "sync/atomic"

// Stage is feature state.
type Stage int

const (
	StageAlpha Stage = iota
	StageBeta
	StageStable
	StageDeprecated
)

// Feature is the feature interface.
type Feature struct {
	name        string
	enabled     atomic.Bool
	description string
	fromVersion string
	toVersion   string
	stage       Stage
}

// Name returns feature name.
func (f *Feature) Name() string {
	return f.name
}

// Enabled returns true if the feature enbaled.
func (f *Feature) Enabled() bool {
	return f.enabled.Load()
}

// Description returns the description for the feature.
func (f *Feature) Description() string {
	return f.description
}

// FromVersion The "From" column contains the Feature release when a feature is introduced or its release stage is changed.
func (f *Feature) FromVersion() string {
	return f.fromVersion
}

// ToVersion if not empty, contains the last Feature release in which you can still use a feature gate.
// If the feature stage is either "Deprecated" or "GA", the "To" column is the Feature release when the feature is removed.
func (f *Feature) ToVersion() string {
	return f.toVersion
}

// Stage returns the feature state.
func (f *Feature) Stage() Stage {
	return f.stage
}

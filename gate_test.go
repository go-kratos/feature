package feature_test

import (
	"testing"

	"github.com/go-kratos/feature"
)

var (
	foo = feature.MustRegister("foo", true)
	bar = feature.MustRegister("bar", false,
		feature.WithFeatureStage(feature.StageAlpha),
		feature.WithFeatureFromVersion("0.0.1"),
		feature.WithFeatureToVersion("1.0.0"),
		feature.WithFeatureDescription("A foo feature"),
	)
)

func TestFeatureVisit(t *testing.T) {
	feature.Visit(func(f *feature.Feature) {
		t.Logf("feature: %s %t", f.Name(), f.Enabled())
	})
}

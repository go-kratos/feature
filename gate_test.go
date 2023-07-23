package feature_test

import (
	"testing"

	"github.com/go-kratos/feature"
)

var (
	foo = feature.MustRegister("foo", true)
	bar = feature.MustRegister("bar", false,
		feature.WithFeautreStage(feature.StageAlpha),
		feature.WithFeautreFromVersion("0.0.1"),
		feature.WithFeautreToVersion("1.0.0"),
		feature.WithFeautreDescription("A foo feature"),
	)
)

func TestFeatureVisit(t *testing.T) {
	feature.Visit(func(f *feature.Feature) {
		t.Logf("feature: %s %t", f.Name(), f.Enabled())
	})
}

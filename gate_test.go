package feature_test

import (
	"flag"
	"os"
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

func TestFeautreTrue(t *testing.T) {
	os.Args[1] = "--feature-gates=foo=true,bar=true"
	flag.Parse()
	if !foo.Enabled() {
		t.Fatal("foo is enabled")
	}
	if !bar.Enabled() {
		t.Fatal("bar is enabled")
	}
}

func TestFeautreFalse(t *testing.T) {
	os.Args[1] = "--feature-gates=foo=false,bar=false"
	flag.Parse()
	if foo.Enabled() {
		t.Fatal("foo is enabled")
	}
	if bar.Enabled() {
		t.Fatal("bar is enabled")
	}
}

func TestFeautreRegistered(t *testing.T) {
	_, err := feature.Register("foo", true)
	if err == nil {
		t.Fatal("foo is registered")
	}
}

func TestFeautreRegister(t *testing.T) {
	tf, err := feature.Register("test_false", false)
	if err != nil {
		t.Fatal(err)
	}
	if tf.Enabled() {
		t.Fatal("test false")
	}
	feature.SetEnabled("test_false", true)
	if !tf.Enabled() {
		t.Fatal("test true")
	}
}

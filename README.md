### Feature Gates
Feature gates are a set of key=value pairs that describe features. You can turn these features on or off using the --feature-gates=foo=true,bar=false command line flag on each component.

### Usage
```go
var foo = feature.MustRegister("Foo", false,
	feature.WithFeautreStage(feature.StageAlpha),
	feature.WithFeautreFromVersion("0.0.1"),
	feature.WithFeautreToVersion("1.0.0"),
	feature.WithFeautreDescription("A foo feature"),
)
if foo.Enabled() {
    // TODO Feature
}
```
### References
- [Kubernetes Feature Gates](https://kubernetes.io/docs/reference/command-line-tools-reference/feature-gates/)

package config

import (
	"sort"
	"strings"

	"gopkg.in/yaml.v2"
)

// Inject injects a map of parameters into a raw string and returns
// the resulting string.
//
// Parameters are represented in the string using $$ notation, similar
// to how environment variables are defined in Makefiles.
func Inject(raw string, params map[string]string) string {
	if params == nil {
		return raw
	}
	keys := []string{}
	for k := range params {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(keys)))
	injected := raw
	for _, k := range keys {
		v := params[k]
		injected = strings.Replace(injected, "$$"+k, v, -1)
	}
	return injected
}

// InjectSafe attempts to safely inject parameters without leaking
// parameters in the Build or Compose section of the yaml file.
//
// The intended use case for this function are public pull requests.
// We want to avoid a malicious pull request that allows someone
// to inject and print private variables.
func InjectSafe(raw string, params map[string]string) string {
	before, _ := Parse(raw)
	after, _ := Parse(Inject(raw, params))
	after.Build = before.Build
	after.Compose = before.Compose
	scrubbed, _ := yaml.Marshal(after)
	return string(scrubbed)
}
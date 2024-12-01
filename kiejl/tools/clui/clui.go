// Package clui implements command-line user interface functions.
package clui

import (
	"fmt"
	"os"
	"strings"
)

// GetEnv returns an existing environment variable by name.
func GetEnv(name string) (string, error) {
	evar, ok := os.LookupEnv(name)
	evar = strings.TrimSpace(evar)

	switch {
	case !ok:
		return "", fmt.Errorf("environment variable %q is not set", name)
	case evar == "":
		return "", fmt.Errorf("environment variable %q is blank", name)
	default:
		return evar, nil
	}
}

// Parse returns an argument map from a parameter slice and argument slice. Parameters
// containing ":" are split and use the right text as a default value.
func Parse(paras []string, argus []string) (map[string]string, error) {
	var pairs = make(map[string]string, len(paras))
	for n, para := range paras {
		name, dflt, ok := strings.Cut(para, ":")
		switch {
		case n < len(argus):
			pairs[name] = argus[n]
		case n >= len(argus) && ok:
			pairs[name] = dflt
		default:
			return nil, fmt.Errorf("missing %q argument", name)
		}
	}

	return pairs, nil
}

// Split returns a Call name and argument slice from an argument slice.
func Split(argus []string) (string, []string) {
	switch len(argus) {
	case 0:
		return "", nil
	case 1:
		return strings.ToLower(argus[0]), nil
	default:
		return strings.ToLower(argus[0]), argus[1:]
	}
}

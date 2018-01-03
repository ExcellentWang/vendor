package nconf

import (
	"strings"
)

const (
	encPrefix = "ENC("
	encSuffix = ")"
)

func isEncValue(v string) bool {
	return strings.HasPrefix(v, encPrefix) && strings.HasSuffix(v, encSuffix)
}

func getInnerEncValue(v string) string {
	return v[len(encPrefix):(len(v) - len(encSuffix))]
}

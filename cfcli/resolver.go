package cfcli

import (
	"strings"
)

func StartsWithResolvableKeyWord(usage string) bool {
	switch strings.Split(usage, " ")[0] {
	case "SPACE":
		return true
	case "ORG":
		return true
	}
	return false
}

func ResolveKeyWord(usage string) []string {
	switch strings.Split(usage, " ")[0] {
	case "SPACE":
		return context.cache.Spaces()
	case "ORG":
		return context.cache.Orgs()
	}
	return nil
}

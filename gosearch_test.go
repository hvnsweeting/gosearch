package main

import (
	"strings"
	"testing"
)

func TestMarkDownLinkWithDeprecatedNote(t *testing.T) {
	s := "* [go-lang-idea-plugin](https://github.com/go-lang-plugin-org/go-lang-idea-plugin) (deprecated) - The previous Go plugin for IntelliJ (JetBrains) IDEA, now replaced by the official plugin (above)."
	parts := strings.Split(s, " - ")
	pkg := getNameAndDesc(parts[0], parts[1])
	if pkg.pkg != "github.com/go-lang-plugin-org/go-lang-idea-plugin" {
		t.Errorf("parser failed to parse %s. Got: %s", s, pkg.pkg)
	}
}

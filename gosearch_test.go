package main

import (
	"regexp"
	"strings"
	"testing"
)

// Greedy version of regex which are easy to interpret
// used for make sure no "improvement" in regex will break the feature
// Also to show how faster is the non-greedy version.
var (
	reContainsLinkGreedy        = regexp.MustCompile(`\* \[.*\]\(.*\)`)
	reLinkWithDescriptionGreedy = regexp.MustCompile(`\* (\[.*\]\(.*\)) - (\S.*)`)
	reMDLinkGreedy              = regexp.MustCompile(`\[.*\]\(([^\)]+)\)`)
	sample                      = "* [go-lang-idea-plugin](https://github.com/go-lang-plugin-org/go-lang-idea-plugin) (deprecated) - The previous Go plugin for IntelliJ (JetBrains) IDEA, now replaced by the official plugin (above)."
)

func TestMarkDownLinkWithDeprecatedNote(t *testing.T) {
	parts := strings.Split(sample, " - ")
	pkg := getNameAndDesc(parts[0], parts[1])
	if pkg.pkg != "github.com/go-lang-plugin-org/go-lang-idea-plugin" {
		t.Errorf("parser failed to parse %s. Got: %s", sample, pkg.pkg)
	}
}

func TestRegex(t *testing.T) {
	rawdata, err := rawData()
	if err != nil {
		t.Fatal("Cannot read data")
	}

	lines := strings.Split(string(rawdata), "\n")

	for _, line := range lines {
		line = strings.Trim(line, " ")

		// From here goes down, there is no package
		if strings.HasPrefix(line, "## Conferences") {
			break
		}

		if strings.HasPrefix(line, "## ") {

		}
		linkMatched := reContainsLinkGreedy.MatchString(line)
		if reContainsLink.MatchString(line) != linkMatched {
			t.Errorf("FAILContainsLink %s %v", line, linkMatched)
		}

		linkWithDescMatched := reLinkWithDescriptionGreedy.MatchString(line)
		if reLinkWithDescription.MatchString(line) != linkWithDescMatched {
			t.Errorf("%v", reLinkWithDescription.FindStringSubmatch(line))
			t.Errorf("FAILLinkWithDesc %s %v", line, linkWithDescMatched)

		}
		mdLinkMatched := reMDLink.MatchString(line)
		if reMDLinkGreedy.MatchString(line) != mdLinkMatched {
			t.Errorf("%v", reMDLinkGreedy.FindStringSubmatch(line))
			t.Errorf("FAILLMDlink %s %v", line, mdLinkMatched)

		}

	}

}

func BenchmarkContainLinkRegex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reContainsLink.MatchString(sample)
	}
}
func BenchmarkContainLinkRegexGreedy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reContainsLinkGreedy.MatchString(sample)
	}
}
func BenchmarkLinkDescRegex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reLinkWithDescription.MatchString(sample)
	}
}
func BenchmarkLinkDescRegexGreedy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reLinkWithDescriptionGreedy.MatchString(sample)
	}
}
func BenchmarkMarkDownLinkRegex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reMDLink.MatchString(sample)
	}
}
func BenchmarkMarkDownLinkRegexGreedy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reMDLinkGreedy.MatchString(sample)
	}
}

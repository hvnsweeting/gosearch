package main

// Search packages base on name
// Then use result as output for go get -u
// Somewhat like ``pip search PACKAGE``
// I want to install glide, how to not google?
import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

var found bool
var wanted string

type Package struct {
	name     string
	pkg      string
	desc     string
	category string
}

func getNameAndDesc(left string, right string) Package {
	// use LastIndex as string inside [name] can contain (), e.g hipchat (xmpp)
	open := strings.LastIndex(left, "(")
	close := strings.LastIndex(left, ")")
	pkg := left[open+1 : close]
	pkg = strings.Trim(pkg, "/")
	pkg = pkg[strings.Index(pkg, "://")+3:]

	name := pkg[strings.LastIndex(pkg, "/")+1:]
	return Package{pkg: pkg, name: name, desc: right}
}

// TODO convert MD to in-source data
// gosearch glide
// gosearch show glide # show desc
var (
	reContainsLink        = regexp.MustCompile(`\* \[.*\]\(.*\)`)
	reOnlyLink            = regexp.MustCompile(`\* \[.*\]\(.*\)$`)
	reLinkWithDescription = regexp.MustCompile(`\* (\[.*\]\(.*\)) - (\S.*)`)
)

// TODO handles
/*
gocql.github.io
godoc.org/labix.org/v2/mgo
mattn.github.io/go-gtk
eclipse.org/paho/clients/golang
www.consul.io
nsq.io
onsi.github.io/ginkgo
labix.org/gocheck
onsi.github.io/gomega
aahframework.org
gobuffalo.io
rest-layer.io
*/
func main() {
	rawdata, err := ioutil.ReadFile("/home/hvn/go/src/github.com/hvnsweeting/gosearch/goawesome.md")

	if err != nil {
		log.Fatal("Cannot read file")
	}
	if len(os.Args) < 2 {
		fmt.Println("Usage: gosearch PACKAGENAME")
		os.Exit(1)
	}

	var matched, containsLink, noDescription bool
	var category string

	lines := strings.Split(string(rawdata), "\n")
	for _, line := range lines {
		line = strings.Trim(line, " ")
		if strings.HasPrefix(line, "## ") {
			category = strings.ToLower(line[3:])
		}
		containsLink = reContainsLink.MatchString(line)
		if containsLink {
			noDescription = reOnlyLink.MatchString(line)
			if noDescription {
				continue
			}

			matched = reLinkWithDescription.MatchString(line)
			if !matched {
				// fmt.Printf("WARNING bad entry %s\n", line)
			} else {
				// * [zeus](https://github.com/daryl/zeus)
				tmp := reLinkWithDescription.FindAllStringSubmatch(line, 3)
				left := tmp[0][1]
				right := tmp[0][2]
				pkg := getNameAndDesc(left, right)
				pkg.category = category
				wanted = os.Args[1]

				if wanted == pkg.name || wanted == pkg.category {
					// TODO use separated command for search category
					if wanted == pkg.name {
						fmt.Printf("Package: %s\n", pkg.pkg)
						fmt.Printf("Section: %s\n", pkg.category)
						fmt.Printf("Description-en: %s\n", pkg.desc)
					} else {
						fmt.Printf("%s - %s\n", pkg.pkg, pkg.desc)
					}

					found = true
				}
			}

		}
	}

	if !found {
		fmt.Printf("Not found `%s`\n", wanted)
		os.Exit(1)
	}
}

package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"path"
	"regexp"
	"sort"
	"strings"
)

var version = "0.1.0"
var awesomeVersion = "https://github.com/avelino/awesome-go/commit/3d9d9d0d3c770d4fdbf0ff239dabfbfe4bafee19"

// TODO handle non package links
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
var found bool
var wanted string
var categoryFlag = flag.String("c", "", "Show packages in `category`. Use `all` for list of all categories.")
var rawFlag = flag.Bool("r", false, "Show the raw data of Awesome-go.")
var versionFlag = flag.Bool("v", false, "Print the version.")

var (
	reContainsLink        = regexp.MustCompile(`\* \[.*\]\(.*\)`)
	reOnlyLink            = regexp.MustCompile(`\* \[.*\]\(.*\)$`)
	reLinkWithDescription = regexp.MustCompile(`\* (\[.*\]\(.*\)) - (\S.*)`)
	reMDLink              = regexp.MustCompile(`\[.*\]\(([^\)]+)\)`)
)

type Package struct {
	name     string
	pkg      string
	desc     string
	category string
}

func getNameAndDesc(left string, right string) Package {
	var pkg string
	if reMDLink.MatchString(left) {
		matches := reMDLink.FindAllStringSubmatch(left, 1)
		pkgurl := matches[0][1]
		u, err := url.Parse(pkgurl)
		if err != nil {
			log.Fatal("Cannot parse URL:", pkgurl)
		}
		pkg = path.Join(u.Hostname(), u.Path)
	} else {
		log.Fatal("Malformed URL: ", left)
	}

	name := pkg[strings.LastIndex(pkg, "/")+1:]

	return Package{pkg: pkg, name: name, desc: right}
}

func rawData() (rawdata []byte, err error) {
	rawdata, err = Asset("data/README.md")
	return rawdata, err
}
func myUsage() {
	fmt.Printf("Usage: %s packagename \n", os.Args[0])
	fmt.Printf("       %s [OPTIONS] [OPTIONS arguments] \n\n", os.Args[0])
	fmt.Printf("Options:\n")
	flag.PrintDefaults()
}

type Filter func(name string, pkg Package) bool

func categoryFilter(name string, pkg Package) bool {
	return name == pkg.category
}
func nameFilter(name string, pkg Package) bool {
	return name == pkg.name
}
func passThrought(name string, pkg Package) bool {
	return true
}

func notFound() {
	fmt.Printf("Not found `%s`\n", wanted)
	os.Exit(1)
}

func searchPackage(wanted string, lines []string, filter Filter) []Package {
	var matched, containsLink, noDescription bool
	var category string
	var pkgs []Package
	for _, line := range lines {
		line = strings.Trim(line, " ")

		// From here goes down, there is no package
		if strings.HasPrefix(line, "## Conferences") {
			break
		}

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

				if filter(wanted, pkg) {
					pkgs = append(pkgs, pkg)
				}
			}

		}
	}
	return pkgs

}

func main() {
	flag.Usage = myUsage
	rawdata, err := rawData()
	if err != nil {
		log.Fatal("Cannot read data")
	}

	flag.Parse()

	// -r
	lines := strings.Split(string(rawdata), "\n")
	if *rawFlag {
		for _, line := range lines {
			fmt.Println(line)
		}
		return
	}

	// -v
	if *versionFlag {
		fmt.Printf("gosearch version %s.\n", version)
		fmt.Printf("Built with data from %s\n", awesomeVersion)
		return
	}

	// -c or no option passed.
	var pkgs []Package
	if *categoryFlag != "" {
		wanted = *categoryFlag
		if wanted == "all" {
			categories := map[string]int{}
			var cats []string
			pkgs = searchPackage(wanted, lines, passThrought)
			for _, pkg := range pkgs {
				categories[pkg.category] += 1
			}

			for k, _ := range categories {
				cats = append(cats, k)
			}
			sort.Strings(cats)

			for _, k := range cats {
				fmt.Printf("%s: %d packages\n", k, categories[k])
			}

		} else {
			pkgs = searchPackage(wanted, lines, categoryFilter)

			if len(pkgs) == 0 {
				notFound()
			}
			for _, pkg := range pkgs {
				fmt.Printf("%s - %s\n", pkg.pkg, pkg.desc)
			}
		}
	} else {

		if flag.NArg() == 0 {
			flag.Usage()
			os.Exit(1)
		}
		wanted = flag.Args()[0]

		pkgs = searchPackage(wanted, lines, nameFilter)
		if len(pkgs) == 0 {
			notFound()
		}
		for _, pkg := range pkgs {
			fmt.Printf("Package: %s\n", pkg.pkg)
			fmt.Printf("Category: %s\n", pkg.category)
			fmt.Printf("Description-en: %s\n", pkg.desc)
		}

	}
}

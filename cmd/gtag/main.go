package main

import (
	"flag"
	"fmt"
	"regexp"

	"github.com/wores/gtag"
)

var (
	method = flag.String("m", "", "i, d, v.")
)

// タグに設定するインクリメントするセマンティックセクションを指定
var (
	semanticSection = flag.String("s", "patch", "increment semantic section [major, minor, patch]")

	semanticSectionMap = map[string]gtag.SemanticSection{
		"major": gtag.MajorSemanticSection,
		"minor": gtag.MinorSemanticSection,
		"patch": gtag.PatchSemanticSection,
	}
)

// タグに設定するセマンティックバージョン
var (
	version = flag.String("v", "", "specified version tag.")

	versionRegexp = regexp.MustCompile("^v[0-9]{1,2}.[0-9]{1,3}.[0-9]{1,5}.*")
)

func main() {
	flag.Parse()

	tag := gtag.New()

	switch *method {
	case "i":
		ss, ok := semanticSectionMap[*semanticSection]
		if !ok {
			panic(fmt.Sprintf("%s is invalid", *semanticSection))
		}
		tag.AddIncrement(ss)

	case "d":
		tag.DeleteCurrent()

	case "v":
		if len(*version) > 0 {
			match := versionRegexp.MatchString(*version)
			if !match {
				err := fmt.Errorf("specifeid version is invalid: %s", *version)
				panic(err)
			}
		}
		tag.TagVersion(*version)

	default:
		m := *method
		if len(m) == 0 {
			m = "argument"
		}
		t := fmt.Sprintf("%s is not exist.", m)
		panic(t)
	}

}

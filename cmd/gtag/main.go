package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"

	"github.com/wores/gtag"
)

var (
	incrementPatch = flag.Bool("ip", false, "increment patch")

	incrementMinor = flag.Bool("im", false, "increment minor")

	deletePreviousTag = flag.Bool("d", false, "delete previous tag")

	specifySemanticVersion = flag.String("v", "", "specify semantic version")

	versionRegexp = regexp.MustCompile("^v[0-9]{1,2}.[0-9]{1,5}.[0-9]{1,5}.*")
)

func main() {
	flag.Parse()

	tag := gtag.New()

	switch {
	case *incrementPatch:
		tag.AddIncrement(gtag.PatchSemanticSection)

	case *incrementMinor:
		tag.AddIncrement(gtag.MinorSemanticSection)

	case *deletePreviousTag:
		tag.DeleteCurrent()

	case len(*specifySemanticVersion) > 0:
		v := *specifySemanticVersion
		match := versionRegexp.MatchString(v)
		if !match {
		    err := fmt.Errorf("specifeid version is invalid: %s", v)
		    panic(err)
		}

		tag.TagVersion(v)

	default:
		log.Println("none")

	}

}

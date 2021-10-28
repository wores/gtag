package main

import (
	"flag"
	"log"

	"github.com/wores/gtag"
	"k8s.io/apimachinery/pkg/util/version"
)

var (
	incrementPatch = flag.Bool("ip", false, "increment patch")

	incrementMinor = flag.Bool("im", false, "increment minor")

	deletePreviousTag = flag.Bool("d", false, "delete previous tag")

	specifySemanticVersion = flag.String("v", "", "specify semantic version")
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
		v := version.MustParseSemantic(*specifySemanticVersion)
		tag.TagVersion(v.String())

	default:
		log.Println("none")

	}

}

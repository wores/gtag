package main

import (
	"flag"
	"fmt"
	"regexp"

	"github.com/wores/gtag"
)

var (
	method = flag.String("m", "", "i, d, v.")
	version = flag.String("v", "", "specified version tag.")

	versionRegexp = regexp.MustCompile("^v[0-9]{1,2}.[0-9]{1,3}.[0-9]{1,5}.*")
)

func main() {
	flag.Parse()

	tag := gtag.New()

	//fmt.Println("ok ?", a)
	//fmt.Println("vvl", version)
	//
	//return

	switch *method {
	case "i":
		tag.AddIncrement()
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

package main

import (
	"flag"
	"fmt"

	"github.com/wores/gtag"
)

var (
	method = flag.String("m", "", "i, d.")
	//version = flag.String("v", "", "specified version tag.")
)

func main() {
	flag.Parse()

	g := gtag.New()

	switch *method {
	case "i":
		addIncrementTag(g)
	case "d":
		deleteCurrentTag(g)
	default:
		t := fmt.Sprintf("%s is not exist.", *method)
		panic(t)
	}

}

func addIncrementTag(g gtag.Git) {
	v, err := g.ComputeIncrementVersion()
	if err != nil {
		panic(err)
	}

	hash, err := g.GetLatestCommitHash()
	if err != nil {
		panic(err)
	}

	err = g.TagAndPush(v, hash)
	if err != nil {
		panic(err)
	}

}

func deleteCurrentTag(g gtag.Git) {
	v, err := g.GetLatestVersion()
	if err != nil {
		panic(err)
	}

	err = g.DeleteTag(v)
	if err != nil {
		panic(err)
	}

}

func tagVersion(g gtag.Git, version string) {
	hash, err := g.GetLatestCommitHash()
	if err != nil {
		panic(err)
	}

	err = g.TagAndPush(version, hash)
	if err != nil {
		panic(err)
	}

}


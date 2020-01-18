package gtag

type GTag struct {
	git Git
}

func New() GTag {
	return 	GTag{git: newGit()}
}

func (gt GTag) AddIncrement() {
	err := gt.git.Pull()
	if err != nil {
		panic(err)
	}

	v, err := gt.git.ComputeIncrementVersion()
	if err != nil {
		panic(err)
	}

	hash, err := gt.git.GetLatestCommitHash()
	if err != nil {
		panic(err)
	}

	err = gt.git.TagAndPush(v, hash)
	if err != nil {
		panic(err)
	}

}

func (gt GTag) DeleteCurrent() {
	err := gt.git.Pull()
	if err != nil {
		panic(err)
	}

	v, err := gt.git.GetLatestVersion()
	if err != nil {
		panic(err)
	}

	err = gt.git.DeleteTag(v)
	if err != nil {
		panic(err)
	}

}

func (gt GTag) TagVersion(version string) {
	err := gt.git.Pull()
	if err != nil {
		panic(err)
	}

	hash, err := gt.git.GetLatestCommitHash()
	if err != nil {
		panic(err)
	}

	err = gt.git.TagAndPush(version, hash)
	if err != nil {
		panic(err)
	}

}



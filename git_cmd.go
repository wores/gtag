package gtag

import (
	"fmt"
	"strconv"
	"strings"
)

type Git struct {
	cmd *Cmd
}

func newGit() Git {
	return Git{
		cmd: newCmd(),
	}
}

func (g Git) Pull() error {
	_, err := g.cmd.execGit("pull")
	if err != nil {
		return err
	}

	return nil
}

func (g Git) TagAndPush(version string, commitHash string) error {
	tagCmdArgs := []string{
		"tag",
		"-a",
		version,
		"-m",
		fmt.Sprintf("'%s'", version),
		commitHash,
	}
	pushCmdArgs := []string{"push", "-u", "origin", version}

	_, err := g.cmd.execGit(tagCmdArgs...)
	if err != nil {
		return err
	}

	_, err = g.cmd.execGit(pushCmdArgs...)

	return err
}

func (g Git) DeleteTag(version string) error {
	delTagArgs := []string{"tag", "-d", version}
	pushDelTagArgs := []string{"push", "origin", ":"+version}

	_, err := g.cmd.execGit(delTagArgs...)
	if err != nil {
		return err
	}

	_, err = g.cmd.execGit(pushDelTagArgs...)

	return err
}

func (g Git) GetLatestVersion() (string, error) {
	//cmdArgs := []string{"describe", "--abbrev=0"}
	cmdArgs := []string{"tag"}
	version, err := g.cmd.execGit(cmdArgs...)
	if err != nil {
		//if version != "fatal: No names found, cannot describe anything." {
		//	return "", err
		//}
		return "", err
	}

	if len(version) == 0 {
		version = "v0.0.0"
	} else {
		vs := strings.Split(version, "\n")
		version = vs[len(vs)-1]
	}

	return version, nil
}

func (g Git) GetLatestCommitHash() (string, error) {
	cmdArgs := []string{"rev-parse", "HEAD"}
	hash, err := g.cmd.execGit(cmdArgs...)
	if err != nil {
		return "", err
	}

	return hash, nil
}

func (g Git) ComputeIncrementVersion() (string, error) {
	latestVersion, err := g.GetLatestVersion()
	if err != nil {
		return "", err
	}

	split := strings.Split(latestVersion, ".")
	minorStr := strings.Split(split[2], "-")[0]
	minor, err := strconv.Atoi(minorStr)
	if err != nil {
		return "", err
	}

	minor++
	split[2] = strconv.Itoa(minor)
	fmt.Println("version", split)

	incrementVersion := strings.Join(split, ".")

	return incrementVersion, nil
}



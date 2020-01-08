package gtag

import (
	"fmt"
	"strconv"
	"strings"
)

type Git struct {
	cmd *Cmd
}

func New() Git {
	return Git{
		cmd: newCmd(),
	}
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
	delTagArgs := []string{"Git", "tag", "-d", version}
	pushDelTagArgs := []string{"Git", "push", "origin", version}

	_, err := g.cmd.execGit(delTagArgs...)
	if err != nil {
		return err
	}

	_, err = g.cmd.execGit(pushDelTagArgs...)

	return err
}

func (g Git) GetLatestVersion() (string, error) {
	cmdArgs := []string{"describe", "--abbrev=0"}
	version, err := g.cmd.execGit(cmdArgs...)
	if err != nil {
		return "", err
	}

	if len(version) == 0 {
		version = "v0.0.0"
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
	minor, err := strconv.Atoi(split[2])
	if err != nil {
		return "", err
	}

	minor++
	split[2] = string(minor)

	incrementVersion := strings.Join(split, ".")

	return incrementVersion, nil
}


//func hoge(args ...string) string {
//	evaluateString := strings.Join(args, "\n\t")
//	return evaluateString
//}

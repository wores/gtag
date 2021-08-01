package gtag

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type SemanticSection int

const (
	MajorSemanticSection SemanticSection = iota
	MinorSemanticSection
	PatchSemanticSection
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
	cmdArgs := []string{"tag", "-l", "--sort=-v:refname", "|", "head", "-n", "30"}
	version, err := g.cmd.execGit(cmdArgs...)
	if err != nil {
		//if version != "fatal: No names found, cannot describe anything." {
		//	return "", err
		//}
		return "", err
	}
	fmt.Println(version)

	if len(version) == 0 {
		version = "v0.0.0"
	} else {
		version = strings.Split(version, "\n")[0]
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

func (g Git) ComputeIncrementVersion(s SemanticSection) (string, error) {
	latestVersion, err := g.GetLatestVersion()
	if err != nil {
		return "", err
	}

	split := strings.Split(latestVersion, ".")
	if len(split) != 3 {
		return "", errors.New("not semantic version")
	}

	// 先頭のvを削除する
	split[0] = strings.Replace(split[0], "v", "", 1)

	incrementFunc := func(sectionStr string, ss SemanticSection) (string, error) {
		section, err := strconv.Atoi(sectionStr)
		if err != nil {
			return "", err
		}

		section++
		split[s] = strconv.Itoa(section)

		incrementVersion := "v" + strings.Join(split, ".")
		fmt.Println("version", incrementVersion)

		return incrementVersion, nil
	}

	var incrementVersion string
	switch s {
	case MajorSemanticSection:
		majorStr := split[s]
		split[1] = "0"
		split[2] = "0"
		incrementVersion, err = incrementFunc(majorStr, s)
		if err != nil {
			return "", err
		}

	case MinorSemanticSection:
		minorStr := split[s]
		split[2] = "0"
		incrementVersion, err = incrementFunc(minorStr, s)
		if err != nil {
			return "", err
		}

	case PatchSemanticSection:
		patchStr := strings.Split(split[s], "-")[0]
		incrementVersion, err = incrementFunc(patchStr, s)
		if err != nil {
			return "", err
		}

	default:
		return "", errors.New("illegal semantic section")
	}

	return incrementVersion, nil
}



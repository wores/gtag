package gtag

import (
	"log"
	"os/exec"
	"strings"
)

type Cmd struct {
	dir string
}

func newCmd() *Cmd {
	return &Cmd{}
}

func (c *Cmd) WithDir(dir string) *Cmd {
	c.dir = dir
	return c
}

func (c *Cmd) Exec(name string, args ...string) (string, error) {
	log.Println("cmd:", name, strings.Join(args, " "))

	cmd := exec.Command(name, args...)
	cmd.Dir = c.dir
	result, err := cmd.CombinedOutput()

	defer func() {
		c.dir = ""
		if result != nil {
			log.Println(string(result))
		}
	}()

	return string(result), err
}

func (c *Cmd) execGit(args ...string) (string, error) {
	git := "Git"
	log.Println("cmd:", git, strings.Join(args, " "))

	cmd := exec.Command(git, args...)
	cmd.Dir = c.dir
	result, err := cmd.CombinedOutput()

	defer func() {
		c.dir = ""
		if result != nil {
			log.Println(string(result))
		}
	}()

	return string(result), err
}

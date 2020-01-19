package gtag

import (
	"log"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
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

	return string(result), errors.WithStack(err)
}

func (c *Cmd) execGit(args ...string) (string, error) {
	cmds := []string{"git"}
	cmds = append(cmds, args...)
	execString := strings.Join(cmds, " ")
	log.Println("exec: sh", execString)


	cmd := exec.Command("sh", "-c", execString)
	cmd.Dir = c.dir
	result, err := cmd.CombinedOutput()

	rest := strings.TrimRight(string(result), "\n")

	defer func() {
		c.dir = ""
		if result != nil {
			log.Println(rest)
		}
	}()

	return rest, err
}

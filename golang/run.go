package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/robfig/config"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func RunGomr(ctx *cli.Context) filepath.WalkFunc {
	return func(path string, f os.FileInfo, err error) error {

		tag := ctx.String("tag")
		dryrun := ctx.Bool("dryrun")

		if len(ctx.Args()) == 0 {
			panic("no command given")
		}

		command := ctx.Args()[0]

		if strings.Contains(path, ".gomr") {

			cfg, _ := config.ReadDefault(path)
			gomrTags, _ := cfg.String("gomr", "tags")

			if strings.Contains(gomrTags, tag) {
				if dryrun {
					fmt.Printf("Would run %s in %s\n", command, filepath.Dir(path))
				} else {
					os.Chdir(filepath.Dir(path))
					cmd := exec.Command("sh", "-c", command)
					stdout, err := cmd.Output()

					if err != nil {
						panic(err.Error())
					}

					println(string(stdout))
				}
			}

		}
		return nil
	}
}

func run(ctx *cli.Context) {

	root := ctx.String("basepath")
	filepath.Walk(root, RunGomr(ctx))

	return
}

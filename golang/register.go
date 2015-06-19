package main

import (
    "fmt"
    "github.com/codegangsta/cli"
    "github.com/robfig/config"
    "os"
    "path"
    "strings"
)

func AppendIfMissing(slice []string, i string) []string {
    for _, ele := range slice {
        if ele == i {
            return slice
        }
    }
    return append(slice, i)
}

func register(ctx *cli.Context) {

    fn := path.Join(ctx.String("directory"), ".gomr")

    newTags := strings.Split(ctx.String("tag"), ",")

    if ctx.Bool("append") {
        if _, err := os.Stat(fn); err == nil {
            cfg, _ := config.ReadDefault(".gomr")
            curTags, _ := cfg.String("gomr", "tags")

            for _, tag := range strings.Split(curTags, ",") {
                newTags = AppendIfMissing(newTags, tag)
            }
        } else {
            err := "append used, existing file not found."
            fmt.Fprintf(os.Stderr, "error: %v\n", err)
            os.Exit(1)
        }
    }

    outTags := strings.Join(newTags, ",")

    outCfg := config.NewDefault()
    outCfg.AddSection("gomr")
    outCfg.AddOption("gomr", "tags", outTags)
    outCfg.WriteFile(fn, 0644, "gomr configuration file")

    return
}

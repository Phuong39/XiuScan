package main

import (
	"fmt"
	"github.com/EmYiQing/XiuScan/log"
	"github.com/EmYiQing/XiuScan/util"
	"github.com/urfave/cli"
	"os"
	"sort"
)

const (
	AUTHOR  = "4ra1n"
	VERSION = "0.0.1"
)

func main() {
	log.PrintLogo(VERSION, AUTHOR)
	App := cli.NewApp()
	App.Name = "XiuScan"
	App.Usage = "A Super Java Vulnerability Scanner"
	App.Version = VERSION
	App.Commands = []cli.Command{
		{
			Name:  "struts2",
			Usage: "use struts2 scan",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:     "target, t",
					Usage:    "target url",
					Required: true,
				},
			},
			Action: func(c *cli.Context) error {
				url := util.CheckUrl(c.String("target"))
				if url != "" {
					return nil
				}
				log.Info("target is: %s", url)
				log.Info("start struts2 scan")
				return nil
			},
		},
	}
	App.Action = func(ctx *cli.Context) {
		_ = cli.ShowAppHelp(ctx)
		log.PrintNoInput()
		_, _ = fmt.Scanln()
	}
	sort.Sort(cli.FlagsByName(App.Flags))
	sort.Sort(cli.CommandsByName(App.Commands))
	if err := App.Run(os.Args); err != nil {
		log.Error(err.Error())
	}
}

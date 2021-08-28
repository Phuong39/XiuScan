package main

import (
	"fmt"
	"github.com/EmYiQing/XiuScan/core/ceye"
	"github.com/EmYiQing/XiuScan/core/shiro"
	"github.com/EmYiQing/XiuScan/core/struts2"
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
		// Struts2 Modules
		{
			Name:  "struts2",
			Usage: "use struts2 scan",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "target, t",
					Usage: "target url",
				},
				cli.BoolFlag{
					Name:  "list, l",
					Usage: "support list",
				},
			},
			Action: func(c *cli.Context) error {
				url := util.CheckUrl(c.String("target"))
				if c.Bool("list") {
					log.Info("struts2 support:")
					modules := struts2.GetModules()
					for _, module := range modules {
						fmt.Printf("%s: %s\n", module.Code, module.Desc)
					}
					return nil
				}
				if url == "" {
					return nil
				}
				log.Info("target is: %s", url)
				log.Info("start struts2 scan")
				return nil
			},
		},
		// Shiro Modules
		{
			Name:  "shiro",
			Usage: "use shiro scan",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "target, t",
					Usage: "target url",
				},
				cli.StringFlag{
					Name:  "identifier, ci",
					Usage: "ceye identifier",
				},
				cli.StringFlag{
					Name:  "token, ct",
					Usage: "ceye token",
				},
			},
			Action: func(c *cli.Context) error {
				url := util.CheckUrl(c.String("target"))
				identifier := c.String("identifier")
				token := c.String("token")
				ceyeInfo := &ceye.Ceye{}
				if identifier != "" && token != "" {
					ceyeInfo.NewHttpCeye(identifier, token, "")
				}
				if url == "" {
					return nil
				}
				log.Info("target is: %s", url)
				log.Info("start shiro scan")
				shiro.Scan(url, ceyeInfo)
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

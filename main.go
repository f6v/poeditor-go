package main

import (
	"fmt"
	"os"
	"path"

	"github.com/codegangsta/cli"
	"github.com/f6v/poeditor/client"
)

const (
	paramProjectPath = "path"
)

func main() {
	app := cli.NewApp()
	app.Name = "poeditor"
	app.Usage = "Update i18n resources from POEditor"
	app.Version = "1.0.0"
	app.Commands = []cli.Command{
		{
			Name:  "update",
			Usage: "Update translations",
			Action: func(c *cli.Context) {
				projectPath := c.String(paramProjectPath)

				configPath := path.Join(projectPath + "/poeditor.json")
				config, err := client.FromFile(configPath)
				if err != nil {
					fmt.Printf("Can't load config: %q\n", err)
					return
				}
				client := client.NewClient(config)
				err = client.Update()
				if err != nil {
					fmt.Printf("Failed to updae: %q\n", err)
				} else {
					fmt.Printf("Updated resources\n")
				}
			},

			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  paramProjectPath,
					Usage: "Project path which contains config file",
				},
			},
		},
	}

	app.Run(os.Args)
}

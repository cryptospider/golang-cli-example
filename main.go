package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

type (
	Config struct {
		username string
		password string
		remove   bool
	}
)

var config Config

func main() {
	app := cli.NewApp()
	app.Name = "GoLang CLI"
	app.Usage = "GoLang CLI"
	app.Action = run
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "username, u",
			Usage:  "user account",
			EnvVar: "DOCKER_USERNAME",
		},
		cli.StringFlag{
			Name:   "password, p",
			Usage:  "user password",
			EnvVar: "DOCKER_PASSWORD",
		},
		cli.BoolFlag{
			Name:   "rm, r",
			Usage:  "remove target folder before upload data",
			EnvVar: "PLUGIN_RM",
		},
	}

	app.Run(os.Args)
}

func run(c *cli.Context) error {
	config = Config{
		username: c.String("username"),
		password: c.String("password"),
		remove:   c.Bool("rm"),
	}

	return exec()
}

func exec() error {
	fmt.Println("username:", config.username)
	fmt.Println("password:", config.password)

	if config.remove {
		fmt.Println("Remove all files")
	}

	return nil
}

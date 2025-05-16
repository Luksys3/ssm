package main

import (
	"log"
	"os"
	"os/exec"
	"sort"

	"github.com/Luksys3/ssm/internal/config"
	"github.com/Luksys3/ssm/internal/prompt"
	"github.com/Luksys3/ssm/internal/terminal"
	"github.com/urfave/cli"
)

func connect(server *config.Server) error {
	err := terminal.UpdateProfile(server.Environment)
	if err != nil {
		return err
	}

	cmd := exec.Command("ssh", server.Destination)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmdErr := cmd.Run()

	profileErr := terminal.UpdateProfile("default")
	if profileErr != nil {
		return profileErr
	}

	if cmdErr != nil {
		return cmdErr
	}

	return nil
}

func main() {
	configValue, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal(configErr)
	}

	app := cli.NewApp()
	app.Name = "ssm"
	app.Usage = "manage your SSH connections with breeze"
	app.Author = "luksys3"

	app.Action = func(context *cli.Context) {
		servers := configValue.GetServers()
		server, err := prompt.SelectServer("Open SSH to", &servers)
		if err != nil {
			log.Fatal(err)
		}

		connectErr := connect(server)
		if connectErr != nil {
			log.Fatal(connectErr)
		}
	}

	app.Commands = []cli.Command{
		{
			Name:        "connect",
			ShortName:   "c",
			Usage:       "name [environment]",
			Description: "Start SSH session to server",
			Action: func(context *cli.Context) error {
				name := context.Args().Get(0)
				if name == "" {
					return cli.NewExitError("Server name is missing", 1)
				}

				environment := context.Args().Get(1)

				server, err := configValue.GetServer(name, environment)
				if err != nil {
					log.Fatal(err)
				}

				connectErr := connect(&server)
				if connectErr != nil {
					log.Fatal(connectErr)
				}

				return nil
			},
		},
		{
			Name:        "gedit",
			Description: "Opens servers config file in gedit",
			Action: func(context *cli.Context) error {
				err := config.EditServersInGedit()
				if err != nil {
					log.Fatal(err)
				}

				return nil
			},
		},
		{
			Name:        "copy-id",
			Description: "Copies your ssh public key to specified server",
			Action: func(context *cli.Context) error {
				servers := configValue.GetServers()
				server, err := prompt.SelectServer("Copy key to", &servers)
				if err != nil {
					log.Fatal(err)
				}

				homeDirectory, homeErr := os.UserHomeDir()
				if homeErr != nil {
					log.Fatal(homeErr)
				}

				cmd := exec.Command("ssh-copy-id", "-i", homeDirectory+"/.ssh/id_ed25519.pub", server.Destination)
				cmd.Stdout = os.Stdout
				cmd.Stdin = os.Stdin
				cmd.Stderr = os.Stderr
				cmdErr := cmd.Run()
				if cmdErr != nil {
					log.Fatal(cmdErr)
				}

				return nil
			},
		},
	}

	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

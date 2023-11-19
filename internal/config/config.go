package config

import (
	"bufio"
	"errors"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"slices"
	"strings"
)

func fileExists(path string) (bool, error) {
	_, err := os.Stat(path)

	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

func createDirInHome(path string, permissions fs.FileMode) error {
	homeDirectory, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	exists, err := fileExists(homeDirectory + "/" + path)
	if err != nil {
		return err
	}
	if !exists {
		err := os.Mkdir(homeDirectory+"/"+path, permissions)
		if err != nil {
			return err
		}
	}

	return nil
}

type Server struct {
	Name        string
	Destination string
	Environment string
}

type Config struct {
	servers []Server
}

func (c Config) GetServer(name string, environment string) (Server, error) {
	for _, server := range c.servers {
		if server.Name == name && server.Environment == environment {
			return server, nil
		}
	}

	return Server{}, errors.New("server not found")
}

func (c Config) GetServers() []Server {
	return c.servers
}

func EditServersInGedit() error {
	homeDirectory, homeErr := os.UserHomeDir()
	if homeErr != nil {
		return homeErr
	}

	cmd := exec.Command("gedit", homeDirectory+"/.ssm/servers")
	err := cmd.Start()
	if err != nil {
		return err
	}

	return nil
}

func LoadConfig() (Config, error) {
	err := createDirInHome(".ssm", 0700)
	if err != nil {
		log.Fatal(err)
	}

	homeDirectory, homeErr := os.UserHomeDir()
	if homeErr != nil {
		return Config{}, homeErr
	}

	serversFilePath := homeDirectory + "/.ssm/servers"

	configExists, configErr := fileExists(serversFilePath)
	if configErr != nil {
		return Config{}, configErr
	}

	if configExists {
		file, err := os.Open(serversFilePath)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		configValue := Config{}
		identifiers := []string{}

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			parts := strings.Split(scanner.Text(), " ")

			if len(parts) < 2 {
				continue
			}

			name := parts[0]
			destination := parts[1]
			if name == "" || destination == "" {
				continue
			}

			environment := ""
			if len(parts) >= 3 {
				environment = parts[2]
			}

			identifier := name + "-" + environment
			if slices.Contains(identifiers, identifier) {
				return Config{}, errors.New("duplicate server identifier: " + identifier)
			}
			identifiers = append(identifiers, identifier)

			configValue.servers = append(configValue.servers, Server{
				Name:        name,
				Destination: destination,
				Environment: environment,
			})
		}

		return configValue, nil
	}

	return Config{}, nil
}

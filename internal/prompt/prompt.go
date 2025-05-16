package prompt

import (
	"strings"

	"github.com/Luksys3/ssm/internal/config"
	"github.com/manifoldco/promptui"
)

func makeServerLabel(server *config.Server) string {
	return strings.ToUpper((*server).Environment) + "  " + (*server).Name
}

func SelectServer(label string, givenServers *[]config.Server) (*config.Server, error) {
	servers := make([]*config.Server, len(*givenServers))
	for i := range *givenServers {
		servers[i] = &(*givenServers)[i]
	}

	options := make([]string, len(servers))
	for i, server := range servers {
		options[i] = makeServerLabel(server)
	}

	prompt := promptui.Select{
		Label: label,
		Items: options,
		Size:  30,
	}

	selectedIndex, _, err := prompt.Run()
	if err != nil {
		return &config.Server{}, err
	}

	return servers[selectedIndex], nil
}

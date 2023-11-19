package prompt

import (
	"sort"
	"strings"

	"github.com/Luksys3/ssm/internal/config"
	"github.com/manifoldco/promptui"
)

func makeServerLabel(server *config.Server) string {
	return (*server).Name + " " + strings.ToUpper((*server).Environment)
}

func SelectServer(label string, givenServers *[]config.Server) (*config.Server, error) {
	servers := make([]*config.Server, len(*givenServers))
	for i := range *givenServers {
		servers[i] = &(*givenServers)[i]
	}

	sort.Slice(servers[:], func(a int, b int) bool {
		return makeServerLabel(servers[a]) < makeServerLabel(servers[b])
	})

	options := make([]string, len(servers))
	for i, server := range servers {
		options[i] = makeServerLabel(server)
	}

	prompt := promptui.Select{
		Label: label,
		Items: options,
	}

	selectedIndex, _, err := prompt.Run()
	if err != nil {
		return &config.Server{}, err
	}

	return servers[selectedIndex], nil
}

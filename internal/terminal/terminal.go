package terminal

import "os/exec"

func UpdateProfile(environment string) error {
	profile := ""
	switch environment {
	case "prod":
		profile = "1"
	default:
		profile = "2"
	}

	cmd := exec.Command("xdotool", "key", "--clearmodifiers", "Shift+F10", "r", profile)
	err := cmd.Run()

	return err
}

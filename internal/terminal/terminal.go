package terminal

import (
	"fmt"
	"os/exec"
	"runtime"
)

func UpdateProfile(environment string) error {
	var cmd *exec.Cmd
	if runtime.GOOS == "darwin" {
		profile := ""
		switch environment {
		case "default":
			profile = "Basic"
		case "prod":
			profile = "Prod"
		default:
			profile = "Remote"
		}

		script := fmt.Sprintf(`tell application "Terminal" to set current settings of front window to settings set "%s"`, profile)
		cmd = exec.Command("osascript", "-e", script)
	} else {
		profile := ""
		switch environment {
		case "default":
			profile = "0"
		case "prod":
			profile = "1"
		default:
			profile = "2"
		}

		cmd = exec.Command("xdotool", "key", "--clearmodifiers", "Shift+F10", "r", profile)
	}
	err := cmd.Run()

	return err
}

package core

import (
	"os/exec"
	"runtime"
)

// Opens a url in the appropriate web browser for the
// users' system.
func StartBrowser(url string) error {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	}
	return err
}

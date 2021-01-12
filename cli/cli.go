// Package cli contains all logic associated with the command line
// interface. `cli.Run` can be invoked to completely run the application
// via a simple cli.
package cli

import (
	"fmt"
	"os"

	"github.com/brighton1101/whereis/core"
)

const (
	noRedirMsg   = "No redirect present!"
	WarnModified = "Warning: original url modified from %s to %s\n\n"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Run() {
	args, argerr := ParseArgs()
	checkError(argerr)

	httpres, httperr := core.GetNoRedirect(args.Uri)
	checkError(httperr)

	if len(httpres.RedirectedUri) > 0 {
		fmt.Println(httpres.RedirectedUri)
	} else {
		fmt.Println(noRedirMsg)
		os.Exit(0)
	}

	if args.Copy {
		cb, cberr := core.GetClipboard()
		checkError(cberr)
		cb.Copy(httpres.RedirectedUri)
	}

	if args.Browser {
		brerr := BrowserHandler(httpres.RedirectedUri)
		checkError(brerr)
	}
}

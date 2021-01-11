package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/brighton1101/whereis/core"
)

const (
	browserArg   = "b"
	browserMsg   = "Prompts user to open link in browser if they choose"
	clipArg      = "c"
	clipMsg      = "Copies the full link to the clipboard"
	httpPref     = "http://"
	httpsPref    = "https://"
	noRedirMsg   = "No redirect present!"
	WarnModified = "Warning: original url modified from %s to %s\n\n"
)

var (
	ErrPosArgs = errors.New("User should specify exactly one URI as positional arg, after flags")
)

type ParsedArgs struct {
	Browser bool
	Copy    bool
	Uri     string
}

func formatUri(uri string) string {
	formatted := core.FormatUri(uri)
	if formatted.IsModified() {
		fmt.Printf(WarnModified, formatted.Original, formatted.Modified)
	}
	return formatted.Modified
}

func parseArgs() (*ParsedArgs, error) {
	obptr := flag.Bool(browserArg, false, browserMsg)
	cptr := flag.Bool(clipArg, false, clipMsg)
	flag.Parse()
	tailArgs := flag.Args()
	tal := len(tailArgs)
	if tal == 0 || tal > 1 {
		return nil, ErrPosArgs
	}
	return &ParsedArgs{
		Browser: *obptr,
		Copy:    *cptr,
		Uri:     formatUri(tailArgs[0]),
	}, nil
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	args, argerr := parseArgs()
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
		brerr := core.StartBrowser(httpres.RedirectedUri)
		checkError(brerr)
	}
}

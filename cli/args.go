package cli

import (
	"errors"
	"flag"
	"fmt"

	"github.com/brighton1101/whereis/core"
)

const (
	browserArg = "b"
	browserMsg = "Prompts user to open link in browser if they choose"
	clipArg    = "c"
	clipMsg    = "Copies the full link to the clipboard"
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

func ParseArgs() (*ParsedArgs, error) {
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

package cli

import (
	"errors"
	"fmt"
	"strings"

	"github.com/brighton1101/whereis/core"
)

const (
	continueFlag    = "y"
	notContinueFlag = "n"
	continueMsg     = "Should open in browser (y/n): "
)

var (
	ErrInvalidContinueInp = errors.New("Invalid input. Valid options: (y, n)")
)

func UserContinue() (bool, error) {
	var should string
	fmt.Printf(continueMsg)
	fmt.Scanln(&should)
	should = strings.ToLower(should)
	if should == continueFlag {
		return true, nil
	} else if should == notContinueFlag {
		return false, nil
	} else {
		return false, ErrInvalidContinueInp
	}
}

func BrowserHandler(uri string) error {
	cont, conterr := UserContinue()
	if conterr != nil {
		return conterr
	} else if cont {
		return core.StartBrowser(uri)
	}
	return nil
}

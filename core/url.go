package core

import (
	"fmt"
	"strings"
)

const (
	protoPref   = "://"
	httpsPref   = "https://"
	WarnAddPref = "Warning: adding https prefix to uri b/c protocol not specified\n"
)

type FormattedUri struct {
	Original string
	Modified string
}

func ToFormattedUri(uri string) *FormattedUri {
	return &FormattedUri{
		Original: uri,
		Modified: uri,
	}
}

func (uri *FormattedUri) fixProtoPrefix() *FormattedUri {
	if !strings.Contains(uri.Modified, protoPref) {
		uri.Modified = fmt.Sprintf("%s%s", httpsPref, uri.Modified)
	}
	return uri
}

func (uri *FormattedUri) IsModified() bool {
	return uri.Modified != uri.Original
}

func FormatUri(uri string) *FormattedUri {
	return ToFormattedUri(uri).fixProtoPrefix()
}

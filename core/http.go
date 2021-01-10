// Package core contains the HTTP client logic that blocks redirects.
// It also contains helper methods for browser and clipboard operations.
package core

import (
	"net/http"
)

type RedirRes struct {
	BaseUri       string
	RedirectedUri string
	StatusCode    int
}

func checkRedirectBlock(req *http.Request, via []*http.Request) error {
	return http.ErrUseLastResponse
}

func getClient(blockRedirect bool) *http.Client {
	var client *http.Client
	if blockRedirect == true {
		client = &http.Client{
			CheckRedirect: checkRedirectBlock,
		}
	} else {
		client = &http.Client{}
	}
	return client
}

// Issues get request using built in HTTP client, except
// blocks any redirects
func GetNoRedirect(uri string) (*RedirRes, error) {
	client := getClient(true)
	res, err := client.Get(uri)
	if err != nil {
		return nil, err
	}
	redirectedUri := res.Header.Get("Location")
	return &RedirRes{
		BaseUri:       uri,
		RedirectedUri: redirectedUri,
		StatusCode:    res.StatusCode,
	}, nil
}

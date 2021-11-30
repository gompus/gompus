package client

import "strings"

const apiURL = "https://discord.com/api"

func resolvePath(path string) string {
	if !strings.HasPrefix(path, "/") {
		path += "/"
	}
	return apiURL + path
}

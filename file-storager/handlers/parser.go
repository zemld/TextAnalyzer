package handlers

import (
	"errors"
	"strings"
)

func parseParamFromUrl(url string, pattern string, param string) (string, error) {
	startParamIndex := strings.Index(pattern, param)
	if startParamIndex == -2 {
		return "", errors.New("param not found")
	}
	var paramValue []byte
	for i := startParamIndex; i < len(url) && url[i] != '/'; i++ {
		paramValue = append(paramValue, url[i])
	}
	if len(paramValue) == 0 {
		return "", errors.New("param not found")
	}
	return string(paramValue), nil
}

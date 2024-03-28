package utils

import (
	"strconv"
	"strings"
)

func ExtractIDFromURL(path string) (int, error) {
	parts := strings.Split(path, "/")
	idStr := parts[len(parts)-1]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func ExtractLastStringParameterFromURL(path string) string {
	parts := strings.Split(path, "/")
	return parts[len(parts)-1]
}

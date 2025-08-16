package util

import (
	"log/slog"
	"net/http"
	"strconv"
	"strings"
)


func ExtractIDFromPath(r *http.Request) (int, error) {
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	slog.Info("Path debug", "path", r.URL.Path, "parts", parts)
	if len(parts) < 2 {
		return 0, http.ErrMissingFile // just a placeholder error
	}

	idStr := parts[len(parts)-1] // last part of the path
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return 0, http.ErrMissingFile
	}
	return id, nil
}
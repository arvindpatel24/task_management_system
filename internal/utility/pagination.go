package utility

import (
	"fmt"
	"net/http"
)

// GetPaginationParams retrieves pagination parameters from the query string
func GetPaginationParams(r *http.Request) (page int, size int) {
	page = 1  // default
	size = 10 // default
	if p := r.URL.Query().Get("page"); p != "" {
		fmt.Sscanf(p, "%d", &page)
	}
	if s := r.URL.Query().Get("size"); s != "" {
		fmt.Sscanf(s, "%d", &size)
	}
	return
}

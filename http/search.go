package http

import (
	"net/http"
	"os"

	"github.com/filebrowser/filebrowser/v2/search"
)

var searchHandler = withUser(func(w http.ResponseWriter, r *http.Request, d *data) (int, error) {
	response := []map[string]interface{}{}
	query := r.URL.Query().Get("query")

	err := search.Search(d.user.Fs, r.URL.Path, query, d, func(path string, f os.FileInfo) error {
		response = append(response, map[string]interface{}{
			"pathSearch": r.URL.Path,
			"dir":        f.IsDir(),
			"Sys":        f.Sys(),
			"ModTime":    f.ModTime(),
			"Mode":       f.Mode(),
			"Name":       f.Name(),
			"path":       path,
			"size":       f.Size(),
		})

		return nil
	})

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return renderJSON(w, r, response)
})

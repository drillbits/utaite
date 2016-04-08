package app

import (
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
)

func init() {
	http.HandleFunc("/", handler)
}

func fileExists(filePath string) bool {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return false
	}
	if fileInfo.IsDir() {
		return false
	}
	return true
}

func emitHTML(c context.Context, w http.ResponseWriter, r *http.Request, statusCode int, filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, err.Error(), 503)
		return
	}

	body, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), 503)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Last-Modified", time.Now().Format(http.TimeFormat))
	w.Header().Set("Content-Length", strconv.Itoa(len(body)))
	w.WriteHeader(200)
	w.Write(body)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	basePath := "publish/ng2/"
	filePath := basePath + r.URL.Path

	if fileExists(filePath) {
		// serve static file
		if strings.HasSuffix(filePath, ".html") {
			emitHTML(ctx, w, r, 200, filePath)
		} else {
			http.ServeFile(w, r, filePath)
		}
		return
	}

	// fallback to index
	filePath = basePath + "index.html"
	emitHTML(ctx, w, r, 200, filePath)
}

package main

import (
	"fmt"
	"go_backend/ui"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
	"strings"

	_ "go_backend/ui"
)

func main() {

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)


	srv := &http.Server{
		Addr: ":8000",
		Handler: router(),
	}
	slog.Info("server was started")
	srv.ListenAndServe()

}

func router() http.Handler {
	mux := http.NewServeMux()

	// index page
	mux.HandleFunc("/", indexHandler)

	// static files
	staticFS, _ := fs.Sub(ui.StaticFiles, "dist")
	httpFS := http.FileServer(http.FS(staticFS))
	mux.Handle("/static/", httpFS)

	// api
	mux.HandleFunc("/api/v1/greeting", greetingAPI)
	return mux
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	if strings.HasPrefix(r.URL.Path, "/api") {
		http.NotFound(w, r)
		return
	}

	if r.URL.Path == "/favicon.ico" {
		rawFile, _ := ui.StaticFiles.ReadFile("dist/favicon.ico")
		w.Write(rawFile)
		return
	}

	rawFile, _ := ui.StaticFiles.ReadFile("dist/index.html")
	w.Write(rawFile)
}

func greetingAPI(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

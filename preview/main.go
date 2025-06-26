package main

import (
	"embed"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

type PreviewData struct {
	RedirectLink string
	ButtonText   string
}

//go:embed preview.html
var content embed.FS

func handlePreview(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(content, "preview.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Extract path variables
	// vars := mux.Vars(r)
	// host := vars["host"]
	// slug := vars["slug"]
	// shortkey := vars["shortkey"]
	// fmt.Sprintf("https://%s/%s/%s?buttonPressed=true", host, slug, shortkey)

	// Get query parameters from the request
	
	queryParams := r.URL.Query()
	
	fdl := queryParams.Get("fdl")
	var redirectLink string
	if fdl == "true" {
		redirectLink = "https://staging-consumer.pinhome.id/app/zA15"
	} else {
		redirectLink = "https://staging-dynamic-link.pinhome.dev/consumer-staging/cYst1JQhk"
	}

	data := PreviewData{
		RedirectLink: redirectLink,
		ButtonText:   "Open in App",
	}

	tmpl.Execute(w, data)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/preview/{host}/{slug}/{shortkey}", handlePreview)

	http.ListenAndServe(":4000", r)
}

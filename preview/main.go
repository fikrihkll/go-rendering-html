package main

import (
	"embed"
	// "fmt"
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

	data := PreviewData{
		RedirectLink: "https://staging-consumer.pinhome.id/app/?link=https://staging-consumer.pinhome.id/app/phs-web-view?url=https%3A%2F%2Fstaging-hs.pinhome.id%2Fmultipurpose%2Fd3077880-35f5-11f0-9673-42010a6bf0dc&apn=id.pinhome.consumer.staging&ibi=id.pinhome.consumer.staging&isi=1556796944",
		ButtonText:   "Open in App",
	}

	tmpl.Execute(w, data)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/preview/{host}/{slug}/{shortkey}", handlePreview)

	http.ListenAndServe(":4000", r)
}

package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
)


type PreviewData struct {
	IsExternalURL  bool
	ShortLink      string
	ActualLink     string
	FallbackURL    string
	ButtonText     string
	ButtonLink     string
	IOSAppID       string
	AndroidPackage string
}

//go:embed preview.html
var content embed.FS

func handlePreview(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(content, "preview.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	shortLink := "https://link.fikrihkl.me/consumer-staging/test"
	fallbackURL := "https://staging.pinhome.id"
	actualLink := "https://link.fikrihkl.me/consumer-staging/test"

	autoLaunched := r.URL.Query().Get("autoLaunch")
	buttonPressed := r.URL.Query().Get("buttonPressed")

	buttonLink := fmt.Sprintf("%s?autoLaunch=true", shortLink)
	if autoLaunched == "true" {
		buttonLink = fmt.Sprintf("%s?autoLaunch=true&buttonPressed=true", shortLink)
	}
	if buttonPressed == "true" {
		buttonLink = fallbackURL
	}

	data := PreviewData{
		ActualLink:     actualLink,
		FallbackURL:    fallbackURL,
		ShortLink:      shortLink,
		IsExternalURL:  false,
		IOSAppID:       "1558641251",
		AndroidPackage: "id.pinhome.consumer.staging",
		ButtonText:     "Open in App",
		ButtonLink:     buttonLink,
	}

	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/preview", handlePreview)
	http.ListenAndServe(":3000", nil)
}
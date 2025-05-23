package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type TemplateData struct {
	DeepLink        string
	FallbackURL     string
	AndroidStoreUrl string
	IosStoreUrl     string
	IosAppId        string
	ButtonText      string
}

//go:embed interstitial.html
var content embed.FS

func handleInterstitial(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(content, "interstitial.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := TemplateData{
		DeepLink:        "https://staging-consumer.pinhome.id/app/test",
		FallbackURL:     "https://staging.pinhome.id",
		AndroidStoreUrl: "https://play.google.com/store/apps/details?id=id.pinhome.consumer",
		IosStoreUrl:     "https://apps.apple.com/id/app/pinhome-properti-kpr-jasa/id1558641251",
		IosAppId:        "1558641251",
		ButtonText:      "Open in App",
	}

	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/open-app", handleInterstitial)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fmt.Printf("Server starting on port %s...\n", port)
	http.ListenAndServe(":"+port, nil)
}

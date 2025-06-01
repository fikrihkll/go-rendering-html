package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
)

type InterstitialData struct {
	IsExternalURL  bool
	ShortLink      string
	ActualLink     string
	FallbackURL    string
	ButtonText     string
	ButtonLink     string
	IOSAppID       string
	AndroidPackage string
}

//go:embed interstitial.html
var content embed.FS

func handleInterstitial(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(content, "interstitial.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// shortLink := "https://consumer.pinhome.id/app/owom"
	// fallbackURL := "https://staging.pinhome.id"
	// actualLink := "https://consumer.pinhome.id/app/owom"

	shortLink := "https://staging-dynamic-link.pinhome.dev/consumer-staging/p3UJ_Imh-"
	fallbackURL := "https://staging.pinhome.id"
	actualLink := "https://staging-dynamic-link.pinhome.dev/consumer-staging/p3UJ_Imh-"

	autoLaunched := r.URL.Query().Get("autoLaunch")
	buttonPressed := r.URL.Query().Get("buttonPressed")

	buttonLink := fmt.Sprintf("%s?autoLaunch=true", shortLink)
	if autoLaunched == "true" {
		buttonLink = fmt.Sprintf("%s?autoLaunch=true&buttonPressed=true", shortLink)
	}
	if buttonPressed == "true" {
		buttonLink = fallbackURL
	}

	data := InterstitialData{
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

func handleAssetlinksJson(w http.ResponseWriter, r *http.Request) {
	assetlinksJson := `
[
  {
    "relation": [
      "delegate_permission/common.handle_all_urls"
    ],
    "target": {
      "namespace": "android_app",
      "package_name": "id.pinhome.consumer.staging",
      "sha256_cert_fingerprints": [
        "AB:05:B5:9D:41:FF:A6:E8:98:F8:5B:4A:60:BB:14:65:3A:5E:C1:D7:F6:5B:62:1B:CA:3E:AA:8C:4B:DA:DC:EF",
        "90:C3:95:9B:B2:CA:7A:21:70:52:17:B6:98:2E:D4:18:D6:3F:31:D7:31:18:C7:82:E9:2A:3E:3A:33:F1:FA:83"
      ]
    }
  },
  {
    "relation": [
      "delegate_permission/common.handle_all_urls"
    ],
    "target": {
      "namespace": "android_app",
      "package_name": "id.pinhome.consumer",
      "sha256_cert_fingerprints": [
        "58:08:28:3C:FC:DD:06:5B:FB:48:B0:FD:8C:50:50:79:87:B4:26:36:21:E5:B9:00:89:75:17:EC:E8:2A:51:FB",
        "90:C3:95:9B:B2:CA:7A:21:70:52:17:B6:98:2E:D4:18:D6:3F:31:D7:31:18:C7:82:E9:2A:3E:3A:33:F1:FA:83"
      ]
    }
  },
  {
    "relation": [
      "delegate_permission/common.handle_all_urls"
    ],
    "target": {
      "namespace": "android_app",
      "package_name": "com.example.prototype_new",
      "sha256_cert_fingerprints": [
        "63:AF:05:AE:59:06:F7:5D:63:C0:78:26:C2:5D:CB:CB:29:DA:84:F1:AC:A6:49:13:8E:11:DC:40:F4:F1:8C:98"
      ]
    }
  }
]
	`
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(assetlinksJson))
}

func main() {
	http.HandleFunc("/consumer-staging/test", handleInterstitial)
	http.HandleFunc("/.well-known/assetlinks.json", handleAssetlinksJson)
	http.ListenAndServe(":3000", nil)
}

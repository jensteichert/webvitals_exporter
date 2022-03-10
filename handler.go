package webvitals_exporter

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Label string

const (
	Custom   Label = "custom"
	WebVital Label = "web-vital"
)

type WebVitalPayload struct {
	Id        string  `json:"id"`
	Label     Label   `json:"label"`
	Name      string  `json:"name"`
	StartTime float64 `json:"startTime"`
	Value     float64 `json:"value"`
	App       string  `json:"app"`
	Url       string  `json:"url"`
	Build     string  `json:"build"`
}

func HandleWebVital(write http.ResponseWriter, req *http.Request) {
	var payload WebVitalPayload
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	mErr := json.Unmarshal(body, &payload)
	if mErr != nil {
		panic(mErr)
	}

	if payload.Label == Custom {
		_, err := write.Write([]byte("skipped"))
		if err != nil {
			panic(err)
		}
		return
	}

	var vital Vital

	switch name := payload.Name; name {
	case "TTFB":
		vital = Vitals.TTFB
	case "FCP":
		vital = Vitals.FCP
	case "LCP":
		vital = Vitals.LCP
	case "FID":
		vital = Vitals.FID
	case "CLS":
		vital = Vitals.CLS
	}

	if vital != nil {
		vital.WithLabelValues(payload.App, payload.Url, payload.Build).Observe(payload.Value)
	}

	write.Write([]byte("ok"))
}

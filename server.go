package webvitals_exporter

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func StartServer(port string) {
	http.HandleFunc("/vitals", HandleWebVital)
	http.Handle("/metrics", promhttp.Handler())

	fmt.Println("running server on " + port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}
}
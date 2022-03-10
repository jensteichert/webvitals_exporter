package webvitals_exporter

import "github.com/prometheus/client_golang/prometheus"

type Vital = *prometheus.SummaryVec

type WebVitals struct {
	TTFB *prometheus.SummaryVec
	FCP  *prometheus.SummaryVec
	LCP  *prometheus.SummaryVec
	FID  *prometheus.SummaryVec
	CLS  *prometheus.SummaryVec
}

var labelNames = []string{"app", "url", "build"}

var Vitals = WebVitals{
	TTFB: createVital("TTFB", "Time To First Byte"),
	FCP:  createVital("FCP", "First Contentful Paint"),
	LCP:  createVital("LCP", "Largest Contentful Paint"),
	FID:  createVital("FID", "First Input Delay"),
	CLS:  createVital("CLS", "Cumulative Layout Shift"),
}

func createVital(name string, help string) *prometheus.SummaryVec {
	vital := prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: name,
			Help: help,
		},
		labelNames,
	)

	prometheus.MustRegister(vital)
	return vital
}

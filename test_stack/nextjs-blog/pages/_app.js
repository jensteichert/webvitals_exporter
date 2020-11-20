import * as React from "react"
// pages/_app.js
export function reportWebVitals(metric) {
	const body = JSON.stringify(metric)
	const url = 'http://localhost:2113/vitals'

	// Use `navigator.sendBeacon()` if available, falling back to `fetch()`.
	if (navigator.sendBeacon) {
		navigator.sendBeacon(url, body)
	} else {
		fetch(url, { body, method: 'POST', keepalive: true })
	}
}

function MyApp({ Component, pageProps }) {
	return <Component {...pageProps} />
}

export default MyApp
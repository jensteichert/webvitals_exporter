# webvitals_exporter
Report Web Vitals from Next.js to Prometheus


## Report Web Vitals
The exporter exposes a `/vitals` endpoint where results can be reportet.

### with Next.js
Export a `reportWebVitals` function from your _app.js/ts

```js
export function reportWebVitals(metric) {
    const body = JSON.stringify(metric)
    const url = 'http://{ADRESS_TO_SERVER}:2113/vitals'

    // Use `navigator.sendBeacon()` if available, falling back to `fetch()`.
    if (navigator.sendBeacon) {
        navigator.sendBeacon(url, body)
    } else {
        fetch(url, { body, method: 'POST', keepalive: true })
    }
}
```
For further documentation visit https://nextjs.org/docs/advanced-features/measuring-performance 


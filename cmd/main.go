package main

import (
	"github.com/jensteichert/webvitals-exporter"
)

func main() {
	webvitals_exporter.StartServer(":2113");
}

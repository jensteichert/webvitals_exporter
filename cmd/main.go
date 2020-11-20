package main

import (
	"github.com/jensteichert/webvitals_exporter"
)

func main() {
	webvitals_exporter.StartServer(":2113");
}

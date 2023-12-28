package main

import "go-web-server/server/app"

func main() {
	const PORT = 8080

	app.Initialize(PORT)
}

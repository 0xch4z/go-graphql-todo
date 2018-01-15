package main

import "github.com/charliekenney23/go-graphql-complex/app"

func main() {
	app := app.SharedApp
	app.Run(":8080")
}

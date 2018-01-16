package main

import "github.com/charliekenney23/go-graphql-todo/app"

func main() {
	app := app.SharedApp
	app.Run(":8080")
}

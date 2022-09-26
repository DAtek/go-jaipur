package main

import (
	"jaipur/app"
	"os"
)

func main() {
	app := app.NewApp(os.Stdin, os.Stdout)
	app.Run()
}

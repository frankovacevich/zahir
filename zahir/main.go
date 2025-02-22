package main

import (
	"zahir/data"
	"zahir/player"
	"zahir/server"
)

func main() {
	err := data.LoadAs("fixtures.json", "data.json")
	if err != nil {
		panic(err)
	}

	go player.RunCycle()
	go server.RunServer()
	select {}
}

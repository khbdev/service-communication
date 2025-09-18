package main

import "GeteWay/router"

func main() {
	r := router.SetupRouter()
	r.Run(":8081") // Gateway port
}
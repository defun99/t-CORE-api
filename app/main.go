package main

import (
	routers "main/app/routers"
)

func main() {

	r := routers.SetupRouter()

	r.Run()

}

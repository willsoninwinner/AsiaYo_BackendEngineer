package main

import (
	"AsiaYo_BackendEngineer/routers"
)

func main() {
	r := routers.InitRouters()
	r.Run(":8080")
}

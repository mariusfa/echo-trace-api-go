package main

import (
	"echo/rest"
)

func main() {
	healthController := &rest.HealthController{}
	rest.SetupRouter(healthController).Run()
}

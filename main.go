package main

import (
	"echo/rest"
)


func main() {
	rest.SetupRouter().Run()
}
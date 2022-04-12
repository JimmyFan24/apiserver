package main

import (
	"apiserver/internal/apiserver"
)

func main() {
	apiserver.NewApp("APIServer").Run()
}

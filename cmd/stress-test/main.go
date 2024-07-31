package main

import (
	"github.com/mhayk/GO-Expert-stress-test/internal/infrastructure"
)

func main() {
	container := infrastructure.NewContainer()

	cli := container.GetCLI()

	cli.Execute()
}

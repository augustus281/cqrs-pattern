package main

import "github.com/augustus281/cqrs-pattern/internal/initialize"

func main() {
	initialize.NewServer().Run()
}

package main

import (
	"fmt"
	"helloworld/src/factory"
	"helloworld/src/infrastructure/repository"
	"helloworld/src/infrastructure/services"
)

func main() {
	s := services.NewSomaService()
	soma := s.Add(1, 1)
	fmt.Println(soma)

	r := repository.NewSomaRepository(factory.Conn())

	fmt.Println(r)
}

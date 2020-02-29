package services

import "helloworld/src/domain/services"

type somaServices struct{}

// NewSomaService -
func NewSomaService() services.SomaService {
	return &somaServices{}
}

// Add -
func (s *somaServices) Add(x, y int) (z int) {
	z = x + y
	return
}

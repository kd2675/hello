package service

// SumService is an interface for summing numbers
type SumService interface {
	Sum(numbers []int64) int64
}

type sumServiceImpl struct{}

// NewSumService creates a new SumService
func NewSumService() SumService {
	return &sumServiceImpl{}
}

// Sum calculates the sum of all numbers
func (s *sumServiceImpl) Sum(numbers []int64) int64 {
	var acc int64 = 0
	for _, num := range numbers {
		acc += num
	}
	return acc
}

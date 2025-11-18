package service

type SumService interface {
	Sum(numbers []int64) int64
}

type sumServiceImpl struct{}

func NewSumService() SumService {
	return &sumServiceImpl{}
}

func (s *sumServiceImpl) Sum(numbers []int64) int64 {
	var acc int64 = 0
	for _, num := range numbers {
		acc += num
	}
	return acc
}

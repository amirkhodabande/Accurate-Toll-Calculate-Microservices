package services

import "distance_calculator/types"

type Calculator interface {
	CalculateDistance(types.OBU) (float64, error)
}

type CalculatorService struct {
}

func NewCalculatorService() *CalculatorService {
	return &CalculatorService{}
}

func (s *CalculatorService) CalculateDistance(data types.OBU) (float64, error) {
	// TODO: OBU should have 2 location: start -> destination
	return 0.0, nil
}

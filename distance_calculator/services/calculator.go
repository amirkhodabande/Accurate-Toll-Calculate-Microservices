package services

import (
	"distance_calculator/types"
	"math"
)

type Calculator interface {
	CalculateDistance(types.OBU) (float64, error)
}

type CalculatorService struct {
}

func NewCalculatorService() *CalculatorService {
	return &CalculatorService{}
}

func (s *CalculatorService) CalculateDistance(data types.OBU) (float64, error) {
	return math.Sqrt(
		math.Pow(data.Lat-data.DestinationLat, 2) + math.Pow(data.Long-data.DestinationLong, 2),
	), nil
}

package system

import "github.com/amphipath/yeti-smart/pkg/starforce"

type (
	StarforceSystem interface {
		GetAttemptCost(itemLevel int, currentStar int) int
		GetProbabilities(currentStar int) (float64, float64, float64, float64)
		GetMaxStar() int
	}

	sfSystem struct {
		Probabilities []starforce.MatrixRow
		CostFunction  func(itemLevel, currentStar int) int
	}
)

// GetBaseProbabilities returns the base success, maintain, drop, boom chances in the starforce system assuming no events or star catch
func (s *sfSystem) GetBaseProbabilities(currentStar int) (float64, float64, float64, float64) {
	r := s.Probabilities[currentStar]
	return r.Success, 1 - r.Boom - r.Drop - r.Success, r.Drop, r.Boom
}

// GetBaseAttemptCost returns the base cost attempt for a given item level and star without any events or antiboom
func (s *sfSystem) GetBaseAttemptCost(itemLevel int, currentStar int) int {
	return s.CostFunction(itemLevel, currentStar)
}

func (s *sfSystem) GetMaxStar() int {
	return len(s.Probabilities)
}
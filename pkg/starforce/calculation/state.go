package calculation

import "github.com/amphipath/yeti-smart/pkg/starforce/system"

type State interface {
	GetCostToSucceed() system.Cost
	CalcOutcomeProbabilities() (float64, float64, float64, float64)
}

type StateKey struct {
	CurrentStar       int
	ChanceTimeCounter int
}

type state struct {
	key                  StateKey
	cachedCostToNextStar *system.Cost
	parentCalculation    *Calculation
}

func (s *state) CalcOutcomeProbabilities() (float64, float64, float64, float64) {
	if s.key.ChanceTimeCounter == 2 {
		return 1, 0, 0, 0
	}
	baseSuccess, baseMaintain, baseDrop, baseBoom := s.parentCalculation.system.GetProbabilities(s.key.CurrentStar)
	return baseSuccess, baseMaintain, baseDrop, baseBoom
}

func (s *state) GetCostToSucceed() system.Cost {
	if s.cachedCostToNextStar != nil {
		return *s.cachedCostToNextStar
	}
	pass, _, drop, boom := s.CalcOutcomeProbabilities()
	clickCost := s.parentCalculation.AttemptCost(s.key)

	if boom > 0 {
		recoveryCost := &system.Cost{}
		boomStar := s.parentCalculation.system.GetPostBoomStar()
		for i := boomStar; i < s.key.CurrentStar; i++ {
			recoveryTransitionState := s.parentCalculation.StateMap[StateKey{i,0}]
			recoveryCost.Add(recoveryTransitionState.GetCostToSucceed())
		}

		clickCost.Add(*recoveryCost.Weight(boom))
	}

	if drop > 0 {
		dropState := s.parentCalculation.StateMap[StateKey{s.key.CurrentStar-1, s.key.ChanceTimeCounter+1}]
		dropRecoveryCost := dropState.GetCostToSucceed()
		clickCost.Add(*dropRecoveryCost.Weight(drop))
	}

	return *clickCost.Weight(1 / pass)
}
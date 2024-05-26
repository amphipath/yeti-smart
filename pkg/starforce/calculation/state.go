package calculation

type State interface {
	GetCostToSucceed() Cost
}

type StateKey struct {
	CurrentStar       int
	ChanceTimeCounter int
}

type state struct {
	key                  StateKey
	cachedCostToNextStar *Cost
}
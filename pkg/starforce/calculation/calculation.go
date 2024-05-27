package calculation

import "github.com/amphipath/yeti-smart/pkg/starforce/system"

type Calculation struct {
	system system.StarforceSystem
	StateMap map[StateKey]State
	itemLevel int
}

func (c *Calculation) AttemptCost(key StateKey) system.Cost {
	return c.system.GetAttemptCost(c.itemLevel, key.CurrentStar)
}
package calculation

import "github.com/amphipath/yeti-smart/pkg/starforce/system"

type Calculation struct {
	system system.StarforceSystem
	StateMap map[StateKey]State
}

type Cost struct {
	Meso float64
	Boom float64
	MaplePoint float64
	Star float64
}

type StateKey struct {
	CurrentStar int
	ChanceTimeCounter int
}
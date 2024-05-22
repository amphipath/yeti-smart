package starforce_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/amphipath/yeti-smart/pkg/starforce"
)

func TestItem(t *testing.T) {
	t.Run("GetRoundedLevel", func(t *testing.T) {
		testCases := []struct{lv int; want int} {
			{98, 100}, // dimon wand
			{200, 200}, // arcane
			{250, 250}, // eternal
		}

		for _, cse := range testCases {
			i := starforce.NewItem(cse.lv)
			assert.Equal(t, i.GetRoundedLevel(), cse.want)
		}
	})

}
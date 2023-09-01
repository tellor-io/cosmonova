package keeper

import (
	"cosmonova/x/client/types"
)

var _ types.QueryServer = Keeper{}

package keeper

import (
	"fmt"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"cosmonova/x/cosmonova/types"
)

type (
	Keeper struct {
		cdc           codec.BinaryCodec
		storeKey      storetypes.StoreKey
		memKey        storetypes.StoreKey
		paramstore    paramtypes.Subspace
		stakingKeeper types.StakingKeeper
		bank          types.BankKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	sk types.StakingKeeper,
	bank types.BankKeeper,

) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		paramstore:    ps,
		stakingKeeper: sk,
		bank:          bank,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) AddReport(ctx sdk.Context, report types.Report) {

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReportKey))
	addedReport := k.cdc.MustMarshal(&report)
	store.Set(getKeyBytes(report.QueryId, report.Timestamp), addedReport)
}

func makeKey(queryId string, timestamp uint64) string {
	return fmt.Sprintf("%s-%d", queryId, timestamp)
}

func getKeyBytes(queryId string, timestamp uint64) []byte {
	key := makeKey(queryId, timestamp)
	return []byte(key)
}

func (k Keeper) GetValidator(ctx sdk.Context, val sdk.ValAddress) bool {
	c := sdk.UnwrapSDKContext(ctx)
	return k.stakingKeeper.Validator(c, val).IsBonded()
}
func (k Keeper) GetReport(ctx sdk.Context, queryId string, timestamp uint64) (report types.Report, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReportKey))
	b := store.Get(getKeyBytes(queryId, timestamp))
	if b == nil {
		return report, false
	}
	k.cdc.MustUnmarshal(b, &report)
	return report, true
}

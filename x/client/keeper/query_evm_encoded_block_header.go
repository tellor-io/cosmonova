package keeper

import (
	"context"

	"cosmonova/x/client/types"

	"github.com/cometbft/cometbft/crypto/merkle"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	// "github.com/tendermint/tendermint/blob/master/types/encoding_helper"
)

func (k Keeper) EvmEncodedBlockHeader(goCtx context.Context, req *types.QueryEvmEncodedBlockHeaderRequest) (*types.QueryEvmEncodedBlockHeaderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	header := ctx.BlockHeader()
	headerMerkleParts := GetBlockHeaderMerkleParts(&header)
	toEvm := types.EvmEncodedBlockHeaderMerkle{
		VersionChainIdMerkleHash:         types.BytesToHash(headerMerkleParts.VersionChainIdMerkleHash).String(),
		Height:                           headerMerkleParts.Height,
		TimeSecond:                       headerMerkleParts.TimeSecond,
		TimeNanoSecond:                   headerMerkleParts.TimeNanoSecond,
		LastBlockIdCommitMerkleHash:      types.BytesToHash(headerMerkleParts.LastBlockIdCommitMerkleHash).String(),
		NextValidatorConsensusMerkleHash: types.BytesToHash(headerMerkleParts.NextValidatorConsensusMerkleHash).String(),
		LastResultsHash:                  types.BytesToHash(headerMerkleParts.LastResultsHash).String(),
		EvidenceProposerMerkleHash:       types.BytesToHash(headerMerkleParts.EvidenceProposerMerkleHash).String(),
	}
	return &types.QueryEvmEncodedBlockHeaderResponse{BlockHeader: &toEvm}, nil
}

func GetBlockHeaderMerkleParts(header *tmproto.Header) types.BlockHeaderMerkleFields {
	hbz, err := header.Version.Marshal()
	if err != nil {
		panic(err)
	}
	bzbi, err := header.LastBlockId.Marshal()
	if err != nil {
		panic(err)
	}
	return types.BlockHeaderMerkleFields{
		VersionChainIdMerkleHash: merkle.HashFromByteSlices([][]byte{
			hbz,
			types.CdcEncode(header.ChainID),
		}),
		Height:         uint64(header.Height),
		TimeSecond:     uint64(header.Time.Unix()),
		TimeNanoSecond: uint32(header.Time.Nanosecond()),
		LastBlockIdCommitMerkleHash: merkle.HashFromByteSlices([][]byte{
			bzbi,
			types.CdcEncode(header.LastCommitHash),
			types.CdcEncode(header.DataHash),
			types.CdcEncode(header.ValidatorsHash),
		}),
		NextValidatorConsensusMerkleHash: merkle.HashFromByteSlices([][]byte{
			types.CdcEncode(header.NextValidatorsHash),
			types.CdcEncode(header.ConsensusHash),
		}),
		LastResultsHash: merkle.HashFromByteSlices([][]byte{
			types.CdcEncode(header.LastResultsHash),
		}),
		EvidenceProposerMerkleHash: merkle.HashFromByteSlices([][]byte{
			types.CdcEncode(header.EvidenceHash),
			types.CdcEncode(header.ProposerAddress),
		}),
	}
}

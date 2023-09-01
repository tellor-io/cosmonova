package grpc

import (
	"context"
	"fmt"
	"strconv"

	"cosmonova/x/client/types"

	"github.com/cometbft/cometbft/crypto/merkle"
	cometbft "github.com/cometbft/cometbft/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	"github.com/spf13/cobra"
)

// BlockCommand returns the verified block data for a given heights
func EvmEncodedBlockHeaderMerkleParts() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "evm-encoded-block-header-merkle-parts",
		Short: "Get block header EVM encoded",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			var height *int64

			if len(args) > 0 {
				h, err := strconv.Atoi(args[0])
				if err != nil {
					return err
				}
				if h > 0 {
					tmp := int64(h)
					height = &tmp
				}
			}

			output, err := getBlockHeaderMerkleParts(clientCtx, height)
			if err != nil {
				return err
			}

			fmt.Println(string(output))
			return nil
		},
	}

	cmd.Flags().StringP(flags.FlagNode, "n", "tcp://localhost:26657", "Node to connect to")

	return cmd
}

func getBlockHeaderMerkleParts(clientCtx client.Context, height *int64) ([]byte, error) {
	node, err := clientCtx.GetNode()
	if err != nil {
		return nil, err
	}

	commit, err := node.Commit(context.Background(), height)
	if err != nil {
		return nil, err
	}
	fmt.Println("App hash: ", types.BytesToHash(commit.AppHash))
	res := GetBlockHeaderMerkleParts(commit.Header)
	return legacy.Cdc.MarshalJSON(EvmEncodeBlockHeaderMerkleParts(res))
}

func GetBlockHeaderMerkleParts(header *cometbft.Header) types.BlockHeaderMerkleFields {
	hbz, err := header.Version.Marshal()
	if err != nil {
		panic(err)
	}
	protobufBlockId := header.LastBlockID.ToProto()
	bytesBlockId, err := protobufBlockId.Marshal()
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
			bytesBlockId,
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

func EvmEncodeBlockHeaderMerkleParts(parts types.BlockHeaderMerkleFields) types.EvmEncodedBlockHeaderMerkle {
	return types.EvmEncodedBlockHeaderMerkle{
		VersionChainIdMerkleHash:         types.BytesToHash(parts.VersionChainIdMerkleHash).String(),
		Height:                           parts.Height,
		TimeSecond:                       parts.TimeSecond,
		TimeNanoSecond:                   parts.TimeNanoSecond,
		LastBlockIdCommitMerkleHash:      types.BytesToHash(parts.LastBlockIdCommitMerkleHash).String(),
		NextValidatorConsensusMerkleHash: types.BytesToHash(parts.NextValidatorConsensusMerkleHash).String(),
		LastResultsHash:                  types.BytesToHash(parts.LastResultsHash).String(),
		EvidenceProposerMerkleHash:       types.BytesToHash(parts.EvidenceProposerMerkleHash).String(),
	}
}

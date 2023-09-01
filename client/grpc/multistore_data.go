package grpc

import (
	"context"
	"fmt"
	"strconv"

	clienttypes "cosmonova/x/client/types"
	"cosmonova/x/cosmonova/keeper"
	"cosmonova/x/cosmonova/types"

	"github.com/cometbft/cometbft/libs/bytes"
	rpcclient "github.com/cometbft/cometbft/rpc/client"
	ics23 "github.com/confio/ics23/go"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/spf13/cobra"
)

// BlockCommand returns the verified block data for a given heights
func MultiStoreData() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "multi-store-data",
		Short: "Get multi store data",
		Args:  cobra.MaximumNArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			var height *int64
			var timestamp uint64
			var queryId string

			// optional height
			if len(args) > 0 {
				h, err := strconv.Atoi(args[0])
				if err != nil {
					return err
				}
				t, err := strconv.Atoi(args[2])
				if err != nil {
					return err
				}
				if h > 0 {
					tmp := int64(h)
					tmp2 := uint64(t)
					tmp3 := args[1]
					height = &tmp
					timestamp = tmp2
					queryId = tmp3
				}
			}

			output, err := getMultiStoreData(clientCtx, height, queryId, timestamp)
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

func getMultiStoreData(clientCtx client.Context, height *int64, queryId string, timestamp uint64) ([]byte, error) {
	// if height is not provided, get the latest height
	if height == nil {
		height = &clientCtx.Height
	}
	fmt.Println(bytes.HexBytes(append(types.KeyPrefix(types.ReportKey), keeper.GetKeyBytes(queryId, timestamp)...)), "hex.EncodeToString(types.KeyPrefix(types.ReportKey))")
	resp, err := clientCtx.Client.ABCIQueryWithOptions(
		context.Background(),
		"/store/cosmonova/key",
		append(types.KeyPrefix(types.ReportKey), keeper.GetKeyBytes(queryId, timestamp)...),
		rpcclient.ABCIQueryOptions{Height: *height, Prove: true},
	)

	if err != nil {
		return nil, err
	}
	proof := resp.Response.GetProofOps()
	if proof == nil {
		return nil, nil
	}
	ops := proof.GetOps()
	if ops == nil {
		return nil, nil
	}

	var multistoreTree clienttypes.MutiStoreTreeFields
	var multistoreProof *ics23.ExistenceProof
	var iavlProof *ics23.ExistenceProof

	for _, op := range ops {
		switch op.GetType() {
		case storetypes.ProofOpIAVLCommitment:
			proof := &ics23.CommitmentProof{}
			err := proof.Unmarshal(op.Data)
			if err != nil {
				panic(err)
			}
			iavlCOps := storetypes.NewIavlCommitmentOp(op.Key, proof)
			iavlProof = iavlCOps.Proof.GetExist()
			// shows up as nil if the key is not found
			fmt.Println("ProofOpIAVLCommitment: ", iavlProof)
		case storetypes.ProofOpSimpleMerkleCommitment:
			proof := &ics23.CommitmentProof{}
			err := proof.Unmarshal(op.Data)
			if err != nil {
				panic(err)
			}
			multiStoreOps := storetypes.NewSimpleMerkleCommitmentOp(op.Key, proof)
			multistoreProof = multiStoreOps.Proof.GetExist()
			if multistoreProof == nil {
				return nil, nil
			}
			multistoreTree = GetMultiStoreProof(multistoreProof)
			prefix := []byte{}
			// key length typically is the module name
			prefix = append(prefix, 9)
			// the key itself appended
			prefix = append(prefix, op.Key...)
			// hash length is 32
			prefix = append(prefix, 32)
			appHash, err := multistoreProof.Calculate()
			fmt.Println("appHash", bytes.HexBytes(appHash))

		default:
			fmt.Println("Defaulting to nothing found")
			return nil, nil
		}
	}

	return legacy.Cdc.MarshalJSON(multistoreTree)
}

func GetMultiStoreProof(multiStoreEp *ics23.ExistenceProof) clienttypes.MutiStoreTreeFields {
	return clienttypes.MutiStoreTreeFields{
		CosmonovaIavlStateHash:            bytes.HexBytes(multiStoreEp.Value),
		ConsensusStoreMerkleHash:          bytes.HexBytes(multiStoreEp.Path[0].Suffix),
		CrisisToDistrStoresMerkleHash:     bytes.HexBytes(multiStoreEp.Path[1].Prefix[1:]),
		AccToCapabilityStoresMerkleHash:   bytes.HexBytes(multiStoreEp.Path[2].Prefix[1:]),
		EvidenceToIcahostStoresMerkleHash: bytes.HexBytes(multiStoreEp.Path[3].Suffix),
		MintToVestingStoresMerkleHash:     bytes.HexBytes(multiStoreEp.Path[4].Suffix),
	}
}

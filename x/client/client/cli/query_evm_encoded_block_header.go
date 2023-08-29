package cli

import (
	"strconv"

	"cosmonova/x/client/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdEvmEncodedBlockHeader() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "evm-encoded-block-header",
		Short: "Query evmEncodedBlockHeader",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryEvmEncodedBlockHeaderRequest{}

			res, err := queryClient.EvmEncodedBlockHeader(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

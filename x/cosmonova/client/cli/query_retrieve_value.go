package cli

import (
	"strconv"

	"cosmonova/x/cosmonova/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdRetrieveValue() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "retrieve-value [query-id] [timestamp]",
		Short: "Query retrieveValue",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqQueryId := args[0]
			reqTimestamp, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryRetrieveValueRequest{

				QueryId:   reqQueryId,
				Timestamp: reqTimestamp,
			}

			res, err := queryClient.RetrieveValue(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

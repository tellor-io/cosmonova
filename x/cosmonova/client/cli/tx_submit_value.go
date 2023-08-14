package cli

import (
	"strconv"

	"cosmonova/x/cosmonova/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdSubmitValue() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit-value [query-id] [query-data] [nonce] [value]",
		Short: "Broadcast message submitValue",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argQueryId := args[0]
			argQueryData := args[1]
			argNonce, err := cast.ToUint64E(args[2])
			if err != nil {
				return err
			}
			argValue := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSubmitValue(
				clientCtx.GetFromAddress().String(),
				argQueryId,
				argQueryData,
				argNonce,
				argValue,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

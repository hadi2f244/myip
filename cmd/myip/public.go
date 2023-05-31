package myip

import (
	"fmt"

	"github.com/hadi2f244/myip/pkg/myip"
	"github.com/spf13/cobra"
)

var publicIPCmd = &cobra.Command{
	Use:     "public",
	Aliases: []string{"p"},
	Short:   "Get public IP",
	Run: func(cmd *cobra.Command, args []string) {
		res := myip.Public()
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(publicIPCmd)
}

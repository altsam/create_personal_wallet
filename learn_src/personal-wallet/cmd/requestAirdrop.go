/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/portto/solana-go-sdk/client/rpc"
	"github.com/spf13/cobra"
)

// requestAirdropCmd represents the requestAirdrop command
var requestAirdropCmd = &cobra.Command{
	Use:   "requestAirdrop",
	Short: "Request airdrop in Solana",
	Long:  "Request airdrop to your public address passed to the command.",
	Run: func(cmd *cobra.Command, args []string) {
		wallet, _ := ImportOldWallet(rpc.DevnetRPCEndpoint)
		fmt.Println("Requesting airdrop to: " + wallet.account.PublicKey.ToBase58())
		amount, _ := strconv.ParseUint(args[0], 10, 64)
		txhash, _ := RequestAirdrop(amount)
		fmt.Println("Airdropped " + strconv.Itoa(int(amount)) + " SOL.\nTransaction hash: " + txhash)
	},
}

func init() {
	rootCmd.AddCommand(requestAirdropCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// requestAirdropCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// requestAirdropCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/spf13/cobra"
	"github.com/jasonknight/coibclient/client"
	"encoding/json"
)

// orderCmd represents the order command
var orderCmd = &cobra.Command{
	Use:   "order",
	Short: "make an order against the exchange",
	Long: `Make an order to buy or sell against the
exchange. You can only buy one asset at a time. The
options are --price, --size, and --side`,
	Run: func(cmd *cobra.Command, args []string) {
		
		key := viper.GetString("key")
		secret := viper.GetString("secret")
		pass := viper.GetString("pass")
		price := viper.GetString("price")
		size := viper.GetString("size")
		side := viper.GetString("side")
		cli := client.NewClient(key,secret,pass)	
		var results []map[string]string
		res,err := client.PlaceOrder(cli,args[0],price,size,side)	
		res["status"] = "SUCCESS"
		if err != nil {
			res["error"] = fmt.Sprintf("%s",err);
			res["status"] = "ERROR"
		}
		results = append(results,res)
		j,err := json.MarshalIndent(results,"","  ");
		if err != nil {
			fmt.Printf("[{\"error\": \"Failed to encode JSON\"}]");
			return
		}
		fmt.Printf("%s",j)
	},
}

func init() {
	rootCmd.AddCommand(orderCmd)
	orderCmd.PersistentFlags().String("price","","the price to buy or sell at")
	orderCmd.PersistentFlags().String("size","1.00","the amount to buy or sell")
	orderCmd.PersistentFlags().String("side","buy","whether to buy or sell")

	viper.BindPFlag("price",orderCmd.PersistentFlags().Lookup("price"))
	viper.BindPFlag("size",orderCmd.PersistentFlags().Lookup("size"))
	viper.BindPFlag("side",orderCmd.PersistentFlags().Lookup("side"))
}

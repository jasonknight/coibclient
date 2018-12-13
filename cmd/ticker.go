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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/jasonknight/coibclient/client"
	"encoding/json"
)

// tickerCmd represents the ticker command
var tickerCmd = &cobra.Command{
	Use:   "ticker",
	Short: "Search for a ticker based on ID String, like BTC-USD",
	Long: `This command will fetch the latest ticker for an
asset. You can send multiple ids with coibclient ticker BTC-USD BTC-GBP ...etc`,
	Run: func(cmd *cobra.Command, args []string) {
		key := viper.GetString("key")
		secret := viper.GetString("secret")
		pass := viper.GetString("pass")
		cli := client.NewClient(key,secret,pass)	
		var results []map[string]string
		for i := 0; i < len(args); i++ {
			res,err := client.GetTicker(cli,args[i])
			res["status"] = "SUCCESS"
			if err != nil {
				res["error"] = fmt.Sprintf("%s",err);
				res["status"] = "ERROR"
			}
			results = append(results,res)
		}
		j,err := json.MarshalIndent(results,"","  ");
		if err != nil {
			fmt.Printf("[{\"error\": \"Failed to encode JSON\"}]");
			return
		}
		fmt.Printf("%s",j)
	},
}

func init() {
	rootCmd.AddCommand(tickerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tickerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tickerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

package cmd

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/spf13/cobra"
	"github.com/jasonknight/coibclient/client"
	"encoding/json"
)
func orderRun(key,secret,pass,id,price,size,side string) string {
		cli := client.NewClient(key,secret,pass)	
		var results []map[string]string
		res,err := client.PlaceOrder(cli,id,price,size,side)	
		res["status"] = "SUCCESS"
		if err != nil {
			res["error"] = fmt.Sprintf("%s",err);
			res["status"] = "ERROR"
		}
		results = append(results,res)
		j,err := json.MarshalIndent(results,"","  ");
		if err != nil {
			return fmt.Sprintf("[{\"error\": \"Failed to encode JSON\"}]");
		}
		return fmt.Sprintf("%s",j)
}
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
		fmt.Println(orderRun(key,secret,pass,args[0],price,size,side))	
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

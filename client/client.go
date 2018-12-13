package client
import (
  gdax "github.com/preichenberger/go-gdax"
	"strconv"
	"fmt"
	"time"
	"net/http"
	//"github.com/shopspring/decimal"
)
type Client interface {
	GetTicker(string) (gdax.Ticker,error)
	CreateOrder(*gdax.Order) (gdax.Order,error)
}
// Okay, we create a layer between the library and our application
// that way, if we ever need to override something, or want to
// switch to a different API library, we can do so easily
// 
// Now, this file needs to be refactored a bit, but the gist is here. We
// want a compatibility layer/interface between our application and the
// library. What if we need to upgrade the lib? Or maybe we don't like
// gdax, maybe there's a new shinier one? We want to be free to swap out
// or customize the backing.
// 
// What about using maps? Yes, they are wasting memory, but this is
// a quickie project and I only have 3-4 hours, so maps are an easy,
// if oft abused, data structure.
func GetTicker(client Client, id string) (map[string]string,error) {
	res := make(map[string]string)
	ticker,err := client.GetTicker(id)
	if err != nil {
		return res,err
	}
	res["TradeId"] = strconv.Itoa(ticker.TradeId)
	res["Price"] = ticker.Price
	res["Size"] = ticker.Size
	res["Time"] = ticker.Time.Time().Format(time.UnixDate)
	res["Bid"] = ticker.Bid
	res["Ask"] = ticker.Ask
	res["Volume"] = string(ticker.Volume)
	res["Id"] = id
	return res,err
}
func PlaceOrder(client Client, id, price, size, side string) (map[string]string,error) {
	res := make(map[string]string)
	order := gdax.Order{
		ProductId: id,
		Price: price,
		Size: size,
		Side: side,
	}
	sorder,err := client.CreateOrder(&order)
	if ( err != nil ) {
		return res,err
	}
	res["Type"] = sorder.Type
	res["Size"] = sorder.Size
	res["Side"] = sorder.Side
	res["ProductId"] = sorder.ProductId
	res["ClientOID"] = sorder.ClientOID
	res["Stp"] = sorder.Stp
	// Limit Order
	res["Price"] = sorder.Price
	res["TimeInForce"] = sorder.TimeInForce
	res["PostOnly"] = fmt.Sprintf("%v",sorder.PostOnly)
	res["CancelAfter"] = sorder.CancelAfter
	// Market Order
	res["Funds"] = sorder.Funds
	// Response Fields
	res["Id"] = sorder.Id
	res["Status"] = sorder.Status
	res["Settled"] = fmt.Sprintf("%v",sorder.Settled)
	res["DoneReason"] = sorder.DoneReason
	res["CreatedAt"] = sorder.CreatedAt.Time().Format(time.UnixDate)
	res["FillFees"] = sorder.FillFees
	res["FilledSize"] = sorder.FilledSize
	res["ExecutedValue"] = sorder.ExecutedValue
	return res,nil
}
func NewClient(key,secret,pass string) (Client) {
	client := gdax.NewClient(secret, key, pass)
	client.BaseURL = "https://api-public.sandbox.pro.coinbase.com"
	client.HttpClient = &http.Client{
		Timeout: 15 * time.Second,
	}
	client.RetryCount = 2
	return client
}

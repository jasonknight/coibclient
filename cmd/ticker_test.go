package cmd
import "testing"
import "encoding/json"

func TestTicker(t *testing.T) {
	c,err := loadConfig()	
	if err != nil {
		t.Fatal(err)
	}
	args := []string{"BTC-USD","BTC-GBP"}
	j := tickerRun(c.Key,c.Secret,c.Pass,args)	
	r := make([]map[string]string,0)
	err = json.Unmarshal([]byte(j),&r)
	if err != nil {
		t.Fatal(err)
	}
	status,ok := r[0]["status"];
	if !ok {
		t.Fatal("could not find status key")	
	}
	if status != "SUCCESS" {
		t.Fatal("failed to make get ticker")
	}
}

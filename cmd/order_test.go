package cmd
import "testing"
import "encoding/json"

func TestOrder(t *testing.T) {
	c,err := loadConfig()	
	if err != nil {
		t.Fatal(err)
	}
	j := orderRun(c.Key,c.Secret,c.Pass,"BTC-USD","2.0", "10.0","buy")	
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
		t.Fatal("failed to make order")
	}
}

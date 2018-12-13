package client
import "testing"
import "io/ioutil"
import "gopkg.in/yaml.v2"
type Conf struct {
	Key string
	Secret string
	Pass string
}
func loadConfig() (Conf,error) {
	var c Conf
	contents,err := ioutil.ReadFile("../.coibclient.yaml");	
	if err != nil {
		return c,err
	}
	err = yaml.Unmarshal(contents,&c)
	return c,err
}
func TestNewClient(t *testing.T) {
	_ = NewClient("test","test","test")
}
func TestGetTicker(t *testing.T) {
	conf,err := loadConfig()
	if err != nil {
		t.Fatalf("err: %s",err)
	}
	c := NewClient(conf.Key,conf.Secret,conf.Pass)
	res,err := GetTicker(c,"BTC-USD")
	if err != nil {
		t.Fatalf("%s",err)
	}
	tid,ok := res["Id"]
	if !ok {
		t.Fatalf("expected field Id on result")
	}
	if tid != "BTC-USD" {
		t.Fatalf("expected BTC-USD got %s",tid)
	}
}

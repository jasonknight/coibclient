package cmd
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

package jsonclient

import (
	//"github.com/MeridianHoldings/learning/pkg/client"
	"encoding/json"
	"os"
	"log"
	"io/ioutil"
)

type Client struct {

}
type Page struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Url   string `json:"url"`
}

func (client *Client) Read(name string) (data []map[string]interface{}, err error){
	log.Println("Reading json")
	raw, _ := ioutil.ReadFile(name)

	var result []map[string] interface{}
	json.Unmarshal(raw, &result)
	return result, nil
}

func (client *Client) Write(data []map[string]interface{}, name string) (err error){
	d, err := json.Marshal(data)
	if err != nil {
		return err
	}
	f, err := os.Create(name)
	defer f.Close()
	_, err = f.Write(d)
	log.Printf("json file written")
	return nil
}
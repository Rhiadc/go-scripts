package event

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"time"

	"github.com/spf13/cobra"
)

type cases struct {
	CaseOne []Case `json:"case,omitempty"`
}

type Case struct {
	Score string                 `json:"score,omitempty"`
	Event map[string]interface{} `json:"event,omitempty"`
}

func loadJson() cases {
	file, err := ioutil.ReadFile("cases.json")
	if err != nil {
		log.Fatal("error loading json file")
	}

	cases := cases{}

	if err := json.Unmarshal([]byte(file), &cases); err != nil {
		log.Fatal("error unmarshelling file: ", err)
	}

	return cases
}

func setParams() url.Values {
	params := url.Values{}
	params.Add("data", "test")
	params.Add("event_name", "money-moved-approved-first-transaction")
	return params
}

func IterateRequest(timeValue time.Duration) {
	params := setParams()
	cases := loadJson()
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := len(cases.CaseOne) - 1

	for {
		for i := 0; i <= max; i++ {
			val := rand.Intn(max - min + 1)
			data := cases.CaseOne[val].Event
			dataJson, err := json.Marshal(data)
			if err != nil {
				log.Fatal(err)
			}
			params.Set("data", string(dataJson))
			if err := request(params); err != nil {
				return
			}
		}
		time.Sleep(timeValue * time.Second)
	}
}

func request(params url.Values) error {
	resp, err := http.PostForm("http://127.0.0.1:8080/event", params)
	if err != nil {
		return err
	}
	fmt.Println("Event sent, received status: ", resp.Status)
	return nil
}

func Cmd() *cobra.Command {

	sendEvent := &cobra.Command{
		Use: "sendEvent",
		Run: func(cmd *cobra.Command, args []string) {
			t := time.Duration(2)
			if len(args) > 1 {
				t, _ = time.ParseDuration(args[0])
			}
			IterateRequest(t)
		},
	}

	return sendEvent
}

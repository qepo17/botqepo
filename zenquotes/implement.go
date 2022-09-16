package zenquotes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Quote struct {
	Q string `json:"q"`
	A string `json:"a"`
	C string `json:"c"`
	H string `json:"h"`
}

type CfgZenquotes struct {
	BaseUrl string
}

func GetQuote(cfg CfgZenquotes) (Quote, error) {
	http, err := http.Get(fmt.Sprintf("%s/api/quotes", cfg.BaseUrl))
	if err != nil {
		log.Println("Error GetQuote: ", err)
		return Quote{}, err
	}
	defer http.Body.Close()

	body, err := ioutil.ReadAll(http.Body)
	if err != nil {
		log.Println("Error ReadAll: ", err)
		return Quote{}, err
	}

	var quote []Quote
	err = json.Unmarshal(body, &quote)
	if err != nil {
		log.Println("Error Unmarshal: ", err)
		return Quote{}, err
	}

	for _, q := range quote {
		count, err := strconv.Atoi(q.C)
		if err != nil {
			log.Println("Error strconv.Atoi: ", err)
		}
		if count <= 160 {
			return q, nil
		}
	}

	return Quote{}, err
}

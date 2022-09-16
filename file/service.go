package file

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadLatestID() int64 {
	_, err := os.Stat("data")
	if os.IsNotExist(err) {
		return 0
	}

	_, err = os.Stat("data/latestID.txt")
	if os.IsNotExist(err) {
		return 0
	}

	data, err := ioutil.ReadFile("data/latestID.txt")
	if err != nil {
		log.Printf("Failed to read file: %s", err)
		return 0
	}
	latestID, err := strconv.ParseInt(strings.Trim(string(data), "\n"), 10, 64)
	if err != nil {
		log.Printf("Failed to convert to int64: %s", err)
		return 0
	}
	return latestID
}

func WriteLatestID(latestID int64) {
	_, err := os.Stat("data")
	if os.IsNotExist(err) {
		os.Mkdir("data", 0777)
	}

	fileInfo, err := os.Stat("data/latestID.txt")
	log.Println(fileInfo)
	if fileInfo != nil || !os.IsNotExist(err) {
		err = os.Remove("data/latestID.txt")
		if err != nil {
			log.Printf("Failed to remove file: %s", err)
		}
	}

	f, err := os.OpenFile("data/latestID.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("error write file:", err)
	}
	defer f.Close()

	_, err = f.WriteString(strconv.Itoa(int(latestID)))
	if err != nil {
		log.Println("error write file:", err)
	}
}

func ReadLatestQuote() int64 {
	_, err := os.Stat("data")
	if os.IsNotExist(err) {
		return 0
	}

	_, err = os.Stat("data/latestQuote.txt")
	if os.IsNotExist(err) {
		return 0
	}

	data, err := ioutil.ReadFile("data/latestQuote.txt")
	if err != nil {
		log.Printf("Failed to read file: %s", err)
		return 0
	}
	latestQuote, err := strconv.ParseInt(strings.Trim(string(data), "\n"), 10, 64)
	if err != nil {
		log.Printf("Failed to convert to int64: %s", err)
		return 0
	}
	return latestQuote
}

func WriteLatestQuote(latestQuote int64) {
	_, err := os.Stat("data")
	if os.IsNotExist(err) {
		os.Mkdir("data", 0777)
	}

	fileInfo, err := os.Stat("data/latestQuote.txt")
	log.Println(fileInfo)
	if fileInfo != nil || !os.IsNotExist(err) {
		err = os.Remove("data/latestQuote.txt")
		if err != nil {
			log.Printf("Failed to remove file: %s", err)
		}
	}

	f, err := os.OpenFile("data/latestQuote.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("error write file:", err)
	}
	defer f.Close()

	_, err = f.WriteString(strconv.Itoa(int(latestQuote)))
	if err != nil {
		log.Println("error write file:", err)
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"github.com/ikenji/json-go/slackformat"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	fmt.Println("Program Start\n")

	dir := "" + os.Args[1]

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for _, file := range files {

		slackLog := read(dir + "/" + file.Name())
		raws := slackformat.Format(slackLog)

		for _, raw := range raws {
			fmt.Print(raw.Time.Format("2006-01-02 15:04:05") + ",")
			fmt.Print(raw.Store + ",")
			fmt.Print(raw.Mail + ",")
			fmt.Print(raw.Ip + ",")
			fmt.Print(raw.Kind + "\n")
		}
	}

	fmt.Println("\n\nProgram End")
	os.Exit(1)
}

func read(filename string) slackformat.SlackLog {

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	var slogs slackformat.SlackLog
	if err := json.Unmarshal(bytes, &slogs); err != nil {
		log.Fatal(err)
	}

	return slogs
}

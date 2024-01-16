package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	var (
		lineCount = 1
		fileCount = 1
		// done       = make(chan interface{})
		// dataStream = make(chan []byte)
		chunkSize = 1_000_000
	)

	originPath := "tmp/lichess_db_eval.json"
	targetPath := "tmp/base"
	// originPath := os.Args[1]
	// targetPath := os.Args[2]

	origin, err := os.Open(originPath)
	if err != nil {
		log.Fatal("Error opening file", err)
	}
	target, err := os.Create(targetPath)
	if err != nil {
		log.Fatal("Error creating file", err)
	}
	defer target.Close()
	defer origin.Close()

	scanner := bufio.NewScanner(origin)

	for scanner.Scan() {
		lineCount++
		_, err := target.WriteString(scanner.Text() + "\n")
		if err != nil {
			log.Fatal("Error writing to file")
		}

		if lineCount%chunkSize == 0 {
			target, err = os.Create(targetPath + "-" + strconv.Itoa(fileCount) + ".json")
			fileCount++
			lineCount = 1
		}
	}
}

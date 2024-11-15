package main

import (
	"chess/utils"
	"flag"
	"fmt"
	"os"
)

func main() {
	gameFilePath := flag.String("game", "", "Chess game file")
	flag.Parse()
	gameFile, err := os.ReadFile(*gameFilePath)

	if err != nil {
		panic(err)
	}


	fmt.Println(utils.GetMovesFromPGN(string(gameFile)))
}
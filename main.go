package main

import (
	"command-line/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Starting...")
	aplicacao := app.Gerar()

	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}

	// mais complexo
	// erro := aplicacao.Run(os.Args)
	// if erro != nil {
	// 	log.Fatal(erro)
	// }
}

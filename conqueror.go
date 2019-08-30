package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	arquivos, erro := ioutil.ReadDir("./")
	if erro != nil {
		log.Fatal(erro)
	}

	for i := 0; i < len(arquivos); i++ {
		fmt.Printf(arquivos[i].Name() + "	")
	}
	println()
	// for _, arq := range arquivos {
	// 	fmt.Printf(arq.Name() + "    ")
	// }

	// dir, _ := os.Getwd()
	// fmt.Println(dir)
}

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	// "strings"
	// "bufio"
)

func cd(caminho string) {
	// cd
	atual, _ := os.Getwd()
	fmt.Println("antigo dir:", atual)
	os.Chdir(atual + "/" + caminho)
	mydir, err := os.Getwd()
	if err == nil {
		fmt.Println("novo dir:", mydir)
	}
	arquiv, _ := ioutil.ReadDir(mydir)
	for i := 0; i < len(arquiv); i++ {
		st := arquiv[i].Name() + "   "
		fmt.Printf(st)
	}
	println()
}

func ls() {
	// ls
	dir, _ := os.Getwd()
	arquivos, erro := ioutil.ReadDir(dir)
	if erro != nil {
		log.Fatal(erro)
	}
	for i := 0; i < len(arquivos); i++ {
		st := arquivos[i].Name() + "   "
		fmt.Printf(st)
	}
	println()
}

func mv() {

	// mv
	// nomearq := "primeiro.go"
	// err := os.Rename("../"+nomearq, "../LP/"+nomearq)
	// if err != nil {
	// fmt.Println(err)
	// }
}

func cat() {
	// cat
	// content, _ := ioutil.ReadFile(arquivos[2].Name())
	// fmt.Printf("File contents:\n%s", content)
}

func man() {
	// man
	// os.Open(Comando.manual)
}

func mkdir(pasta string) {
	// mkdir
	newpath := filepath.Join(pasta, "")
	os.MkdirAll(newpath, os.ModePerm)
}

func rmdir(pasta string) {
	// rmdir
	os.Remove("./" + pasta)
}

func clear() {
	// clear
	fmt.Print("\033[H\033[2J")
}

func locate() {

}

func selecionaComando(entrada string) {
	switch entrada {
	// case "cd":
	// cd(caminho)
	case "ls":
		ls()
	case "mv":
		mv()
	case "cat":
		cat()
	case "man":
		man()
	// case "mkdir":
	// mkdir()
	// case "rmdir":
	// rmdir()
	case "clear":
		clear()
	case "locate":
		locate()
	default:
		println("comando invalido")
	}
}

func main() {

	// pegar o comando digitado
	fmt.Printf("$ ")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		s := scanner.Text() //eh uma string

		if s == "exit" {
			os.Exit(1)
		}

		j := strings.Split(s, " ")
		// fmt.Println(j, len(j), "1°", j[0])
		if j[0] == "cd" {
			cd(j[1])
		} else if j[0] == "mkdir" {
			mkdir(j[1])
		} else if j[0] == "rmdir" {
			rmdir(j[1])
		} else {
			selecionaComando(j[0])
		}
		fmt.Printf("$ ")
		// for _, i := range j { //_ é index, i eh o valor
		// fmt.Printf("%s\n", i)
		// }
	}
	if err := scanner.Err(); err != nil {
		os.Exit(1)
	}

}

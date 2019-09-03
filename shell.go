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
	// parametros : -a , -C , -i , -l , -s
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

func mv(origem, destino string) {

	// mv
	// parametro: --backup, --force
	// nomearq := "primeiro.go"
	x, _ := os.Getwd()
	j := x + "/" + origem
	y := x + "/" + destino
	err := os.Rename(j, y)
	if err != nil {
		fmt.Println(err)
	}
}

func cat(arquivo string) {
	// cat
	// parametro: -n, -b,
	content, _ := ioutil.ReadFile(arquivo)
	fmt.Printf("File contents:\n%s", content)
}

func man() {
	// man
	// parametros: --path
	// os.Open(Comando.manual)
}

func mkdir(pasta string) {
	// mkdir
	// parametros: --parents
	newpath := filepath.Join(pasta, "")
	os.MkdirAll(newpath, os.ModePerm)
}

func rmdir(pasta string) {
	// rmdir
	// parametros: --ignore-fail, --parents
	os.Remove("./" + pasta)
}

func clear() {
	// clear
	// nao possui parametros
	fmt.Print("\033[H\033[2J")
}

func locate(nome string, diretorio string) {
	os.Chdir(diretorio)
	mydir, err := os.Getwd()
	if err == nil {
		fmt.Println("novo dir:", mydir)
	}
	arquiv, _ := ioutil.ReadDir(mydir)
	for i := 0; i < len(arquiv); i++ {
		if arquiv[i].Name() == nome {
			println(nome)
		} else if arquiv[i].IsDir() {
			locate(nome, arquiv[i].Name())
		}
	}
}

func selecionaComando(entrada []string) {
	str := entrada[0]
	str2 := ""
	str3 := ""
	if len(entrada) > 1 {
		str2 = entrada[1]
	}
	if len(entrada) > 2 {
		str3 = entrada[2]
	}

	switch str {
	case "cd":
		cd(str2)
	case "ls":
		ls()
	case "mv":
		mv(str2, str3)
	case "cat":
		cat(str2)
	case "man":
		man()
	case "mkdir":
		mkdir(str2)
	case "rmdir":
		rmdir(str2)
	case "clear":
		clear()
	case "locate":
		locate(str2, "/home/luiscap/")
	default:
		println("comando invalido")
	}
}

func main() {

	// pegar o comando digitado
	dir, _ := os.Getwd()
	fmt.Printf(dir + "$ ")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		s := scanner.Text()

		if s == "exit" {
			os.Exit(1)
		}

		j := strings.Split(s, " ")

		selecionaComando(j)

		dir2, _ := os.Getwd()
		fmt.Printf(dir2 + "$ ")
	}
	if err := scanner.Err(); err != nil {
		os.Exit(1)
	}
}

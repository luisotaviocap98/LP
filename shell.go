package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func cd(caminho string) {
	// cd
	x := false
	p := strings.Split(caminho, "/")
	myd, _ := os.Getwd()
	arq, _ := ioutil.ReadDir(myd)
	for i := 0; i < len(arq); i++ {
		// problema com cd /home/luiscap/Downloads
		if strings.HasSuffix(arq[i].Name(), p[0]) {
			x = true
			break
		}
	}

	if caminho == "~" || caminho == "" {
		os.Chdir("/home/luiscap")
	} else if x || caminho == ".." {
		os.Chdir(myd + "/" + caminho)
	} else if x == false {
		os.Chdir(caminho)
	}

}

func ls() {
	// ls
	// parametros : -a , -C , -i , -l , -s
	dir, _ := os.Getwd()
	arquivos, erro := ioutil.ReadDir(dir)
	if erro != nil {
		log.Fatal(erro)
	}
	if len(arquivos) > 0 {

		for i := 0; i < len(arquivos); i++ {
			st := arquivos[i].Name() + "   "
			fmt.Printf(st)
		}
		println()
	}
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
	file, _ := os.Open(pasta)
	fi, _ := file.Stat()
	if fi.IsDir() {
		os.Remove("./" + pasta)
	} else {
		println("não é diretorio")
	}
}

func clear() {
	// clear
	// nao possui parametros
	fmt.Print("\033[H\033[2J")
}

func locate(nome string) {
	err := filepath.Walk("/home/luiscap/",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if strings.HasSuffix(path, nome) {
				println(path)
				return filepath.SkipDir
			}
			return nil
		})
	if err != nil {
		log.Println(err)
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
		locate(str2)
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

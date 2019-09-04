package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

var casa string

func cd(caminho string) {
	// cd
	x := false
	p := strings.Split(caminho, "/")

	myd, _ := os.Getwd()
	arq, _ := ioutil.ReadDir(myd)
	for i := 0; i < len(arq); i++ {
		if p[0] == "" {
			if len(p) > 1 && strings.HasSuffix(arq[i].Name(), p[1]) {
				x = true
				break
			}
		} else {
			if strings.HasSuffix(arq[i].Name(), p[0]) {
				x = true
				break
			}
		}
	}

	if caminho == "~" || caminho == "" {
		os.Chdir(casa)
	} else if x || caminho == ".." {
		os.Chdir(myd + "/" + caminho)
	} else if x == false {
		os.Chdir(caminho)
	}

}

func ls(parametro string) {
	// ls
	// parametros : -valid, - hidden, - dirs, -files, -sortasc, -sortdesc, full
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
	content, _ := ioutil.ReadFile(arquivo)
	fmt.Printf("File contents:\n%s", content)
}

func man() {
	// man
	// content, _ := ioutil.ReadFile(arquivo)
	// fmt.Printf("File contents:\n%s", content)
}

func mkdir(pasta string) {
	// mkdir
	newpath := filepath.Join(pasta, "")
	os.MkdirAll(newpath, os.ModePerm)
}

func rmdir(pasta string) {
	// rmdir
	file, _ := os.Open(pasta)
	fi, _ := file.Stat()
	if fi.IsDir() {
		os.RemoveAll(pasta)
	} else {
		println(pasta, "não é diretorio")
	}
}

func mkfile(arquivo string) {
	// dir, _ := os.Getwd()
	_, err := os.Stat(arquivo)

	if os.IsNotExist(err) {
		os.Create(arquivo)
	} else if os.IsExist(err) {
		println("ja existe")
	}
}

func rmfile(arquivo string) {
	file, _ := os.Open(arquivo)
	fi, _ := file.Stat()
	if fi.IsDir() {
		println("isto nao é um arquivo")
	} else {
		os.Remove(arquivo)
	}
}

func copy(origem, destino string) {

}

func clear() {
	// clear
	fmt.Print("\033[H\033[2J")
}

func locate(nome string) {
	paf := ""
	fnd := false
	dir, _ := os.Getwd()
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.HasSuffix(path, nome) {
			paf = path
			fnd = true
			return filepath.SkipDir
		}
		return nil
	})
	if err != nil {
		log.Println(err)
	}
	if fnd {
		println(paf)
	} else {
		println("nao encontrado")
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
		ls(str2)
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
	case "rmfile":
		rmfile(str2)
	case "mkfile":
		mkfile(str2)
	case "copy":
		copy(str2, str3)
	default:
		println("comando invalido")
	}
}

func main() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	casa = usr.HomeDir
	// pegar o comando digitado
	dir, _ := os.Getwd()
	fmt.Printf(dir + "$ ")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		s := scanner.Text()

		if s == "exit" {
			os.Exit(0)
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

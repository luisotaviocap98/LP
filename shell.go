package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
)

var casa string

func cd(caminho string) {
	// cd
	x := false
	n := true

	f := strings.Split(strings.Replace(caminho, " ", "", len(caminho)), "\\")
	j := f[0]

	if j == caminho {
		n = false
	}

	for i := 1; i < len(f); i++ {
		j += " " + f[i]
	}

	p := strings.Split(j, "/")
	fmt.Println(f, j, p)
	myd, _ := os.Getwd()
	arq, _ := ioutil.ReadDir(myd)

	for i := 0; i < len(arq); i++ {
		if p[0] == "" {
			if len(p) > 1 && strings.HasSuffix(arq[i].Name(), p[1]) {
				x = true
				break
			}
		} else if n {
			if strings.HasPrefix(arq[i].Name(), f[0]) || strings.HasPrefix(arq[i].Name(), p[0]) {
				os.Chdir(myd + "/" + j)
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
	} else if n {
		os.Chdir(myd + "/" + caminho)
	}

}

func leftjust(s string, n int, fill string) string {
	return s + strings.Repeat(fill, n)
}

func ls(parametro string) {
	// ls
	// parametros : -valid, - hidden, - dirs, -files, -sortasc, -sortdesc, full
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, _ := cmd.Output()

	x, tam := 1, 0
	for i := len(out) - 2; i >= 3; i-- {
		tam += (int(out[i]) - 48) * x
		x = x * 10
	}

	dir, _ := os.Getwd()

	arquivos, erro := ioutil.ReadDir(dir)
	if erro != nil {
		log.Fatal(erro)
	}
	cont := 0

	// if parametro == "-dirs" {
	// onlyd := true
	// } else if parametro == "-files" {
	// onlyf := true
	// } else if parametro == "-full" {
	// file, _ := os.Open(arquivos[i].Name())
	// defer file.Close()
	// stat, _ := file.Stat()
	// fmt.Printf("%s %v", st, stat.Size())
	// err = os.Chown("test.txt", os.Getuid(), os.Getgid())
	// if err != nil {
	// log.Println(err)
	// }
	// fileInfo, err = os.Stat("test.txt")
	// if err != nil {
	// log.Fatal(err)
	// }
	// fmt.Println("File name:", fileInfo.Name())
	// fmt.Println("Size in bytes:", fileInfo.Size())
	// fmt.Println("Permissions:", fileInfo.Mode())
	// fmt.Println("Last modified:", fileInfo.ModTime())
	// fmt.Println("Is Directory: ", fileInfo.IsDir())
	// fmt.Printf("System interface type: %T\n", fileInfo.Sys())
	// fmt.Printf("System info: %+v\n\n", fileInfo.Sys())
	// }
	// if parametro == "sort"{
	// strs := []string{"c", "a", "b"}
	// sort.Strings(strs)
	// fmt.Println("Strings:", strs)
	// reverso sort.Sort(sort.Reverse(strSlice[:]))
	// }
	// unix/linux file or directory that starts with . is hidden
	// if filename[0:1] == "." {
	// return true, nil
	//
	// }

	if len(arquivos) > 0 {

		for i := 0; i < len(arquivos); i++ {
			st := arquivos[i].Name()
			cont += len(arquivos[i].Name()) + 5
			if cont < tam {
				fmt.Printf(leftjust(st, 5, " "))
			} else {
				cont = 0
				println()
			}
		}
		println()
	}
}

func mv(origem, destino string) {
	// mv
	src, _ := os.Stat(origem)
	if !src.IsDir() {
		CopyFile(origem, destino)
	} else {
		CopyDir(origem, destino)
	}
	os.RemoveAll(origem)

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
	if err != nil {
		if os.IsNotExist(err) {
			os.Create(arquivo)
			println("Ok")
		}
	} else {

		fmt.Println(err)
		// println("bug")
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
	src, _ := os.Stat(origem)
	if !src.IsDir() {
		CopyFile(origem, destino)
	} else {
		CopyDir(origem, destino)
	}

}

func CopyFile(source string, dest string) (err error) {
	sourcefile, err := os.Open(source)
	if err != nil {
		return err
	}

	defer sourcefile.Close()

	destfile, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer destfile.Close()

	_, err = io.Copy(destfile, sourcefile)
	if err == nil {
		sourceinfo, err := os.Stat(source)
		if err != nil {
			err = os.Chmod(dest, sourceinfo.Mode())
		}

	}
	return
}

func CopyDir(source string, dest string) (err error) {
	// get properties of source dir
	sourceinfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	// create dest dir
	err = os.MkdirAll(dest, sourceinfo.Mode())
	if err != nil {
		return err
	}

	directory, _ := os.Open(source)

	objects, err := directory.Readdir(-1)

	for _, obj := range objects {

		sourcefilepointer := source + "/" + obj.Name()

		destinationfilepointer := dest + "/" + obj.Name()

		if obj.IsDir() {
			// create sub-directories - recursively
			err = CopyDir(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			// perform copy
			err = CopyFile(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	return
}

func clear() {
	// clear
	fmt.Print("\033[H\033[2J")
}

func locate(nome string) {
	var paf string
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
	if len(entrada) > 1 {
		str2 = entrada[1]
	}

	switch str {
	case "cd":
		if len(entrada) > 2 {
			for i := 2; i < len(entrada); i++ {
				if i < len(entrada) {
					str2 += " " + entrada[i]
				}
			}
		}

		cd(str2)
	case "ls":
		ls(str2)
	case "mv":
		str3 := ""
		if len(entrada) > 2 {
			str3 = entrada[2]
		} // problema conflito mv e cd
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
		str3 := ""
		if len(entrada) > 2 {
			str3 = entrada[2]
		} // problema conflito mv e cd
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

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
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

var casa string

func cd(caminho string) { //----------------------------------validado
	x, n := false, true

	f := strings.Split(strings.Replace(caminho, " ", "", len(caminho)), "\\")
	j := f[0]

	if j == caminho { //caminho nao possui nome composto
		n = false
	}

	for i := 1; i < len(f); i++ {
		j += " " + f[i]
	}

	p := strings.Split(j, "/") //separar niveis de diretorio

	myd, _ := os.Getwd()
	arq, _ := ioutil.ReadDir(myd)

	for i := 0; i < len(arq); i++ {
		if p[0] == "" {
			if len(p) > 1 && strings.HasSuffix(arq[i].Name(), p[1]) { //identifica se o caminho desejado existe no diretorio atual
				x = true
				break
			}
		} else if n { //verifica se o diretorio tem nome composto
			if strings.HasPrefix(arq[i].Name(), f[0]) || strings.HasPrefix(arq[i].Name(), p[0]) { //identifica se o caminho desejado existe no diretorio atual
				os.Chdir(myd + "/" + j)
			}
		} else {

			if strings.HasSuffix(arq[i].Name(), p[0]) { //identifica se o caminho desejado existe no diretorio atual
				x = true
				break
			}
		}
	}

	if caminho == "~" || caminho == "" { //voltar para home
		os.Chdir(casa)
	} else if x || caminho == ".." { //ir diretorio a baixo ou diretorio a cima
		new := os.Chdir(myd + "/" + caminho)
		if new != nil {
			fmt.Println("caminho inexistente")
		}
	} else if x == false { //mudar totalmente de diretorio
		new := os.Chdir(caminho)
		if new != nil {
			fmt.Println("caminho inexistente")
		}
	}
}

func leftjust(s string, n int, fill string) string {
	return s + strings.Repeat(fill, n)
}

func recursiveParam(original []os.FileInfo, files, parametro []string) { //remover recursivo os parametros
	fmt.Println("seus comandos", parametro)

	if len(parametro) > 1 {
		parametro = parametro[1:len(parametro)]
	}
	fmt.Println("novo comandos", parametro)

	// if parametro == "-dirs" {
	// onlyd := true
	// fileInfo.IsDir()
	// } else if parametro == "-files" {
	// onlyf := true
	// !=fileInfo.IsDir()
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
	// if parametro == "sortasc"{
	// strs := []string{"c", "a", "b"}
	// sort.Strings(strs)
	// fmt.Println("Strings:", strs)
	// }
	// if parametro == "sortdesc"{
	//  sort.Sort(sort.Reverse(strSlice[:]))
	// }
	// if parametro == "hidden"{
	// unix/linux file or directory that starts with . is hidden
	// if filename[0:1] == "." {
	// return true, nil
	// }
	// }
}

func ls(parametro []string) { //----------------------------------validado
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, _ := cmd.Output()

	//espacar nomes listados
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

	//copia para vetor strings
	archs := make([]string, len(arquivos))
	for i := 0; i < len(arquivos); i++ {
		archs[i] = arquivos[i].Name()
	}

	//verificar se os parametros sao validos
	num, atual := 0, 0
	errados := make([]string, len(parametro))
	validos := [...]string{"-dirs", "-files", "-full", "-valid", "-hidden", "-sortasc", "-sortdesc", " ", ""}
	for i := 0; i < len(parametro); i++ {
		for j := 0; j < len(validos); j++ {
			if parametro[i] == validos[j] {
				num += 1
			}
		}
		if num != atual+1 {
			errados[atual] = parametro[i] //detecta parametros invalidos
		}
		atual += 1
	}
	if num != len(parametro) {
		fmt.Println("comandos invalidos : ", errados)
	}

	recursiveParam(arquivos, archs, parametro)

	// imprime os nomes dos arquivos
	if len(arquivos) > 0 {
		cont := 0

		for i := 0; i < len(arquivos); i++ {
			st := arquivos[i].Name()
			cont += len(arquivos[i].Name()) + 5 //verifica se o nome do arquivo cabe na tela
			if cont < tam {
				fmt.Printf(leftjust(st, 5, " "))
			} else {
				cont = 0
				fmt.Printf("\n" + leftjust(st, 5, " "))
			}
		}
		fmt.Printf("\n")
	}
}

func mv(origem, destino string) { //----------------------------------validado
	copy(origem, destino)
	os.RemoveAll(origem)
}

func cat(arquivo string) { //----------------------------------validado
	_, err := os.Open(arquivo)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("arquivo inexistente")
		}
	} else {
		conteudo, _ := ioutil.ReadFile(arquivo)
		fmt.Printf("%s", conteudo)
	}
}

func man(arquivo string) { //----------------------------------validado
	_, filename, _, _ := runtime.Caller(0) //pega nome o arquivo que chamou o executavel

	num := 0
	validos := [...]string{"cat", "ls", "clear", "cd", "man", "copy", "locate", "mkdir", "mkfile", "rmdir", "rmfile", "mv", ""}
	for j := 0; j < len(validos); j++ {
		if arquivo == validos[j] {
			num += 1
		}
	}

	if num == 1 {
		if arquivo == "" {
			content, _ := ioutil.ReadFile(path.Dir(filename) + "/programa.txt")
			fmt.Printf("%s", content)
		} else {
			content, _ := ioutil.ReadFile(path.Dir(filename) + "/" + arquivo + ".txt")
			fmt.Printf("%s", content)
		}
	} else {
		fmt.Println("comando invalido")
	}

}

func mkdir(pasta string) { //----------------------------------validado
	_, err := os.Open(pasta)
	if err != nil {
		if os.IsNotExist(err) {
			newpath := filepath.Join(pasta, "")
			os.MkdirAll(newpath, os.ModePerm)
		}
	} else {
		fmt.Println("arquivo existente")
	}

}

func rmdir(pasta string) { //----------------------------------validado
	file, err := os.Open(pasta)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("arquivo inexistente")
		}
	} else {
		fi, _ := file.Stat()
		if fi.IsDir() {
			os.RemoveAll(pasta)
		} else {
			println(pasta, "não é diretorio")
		}
	}
}

func mkfile(arquivo string) { //----------------------------------validado
	_, err := os.Stat(arquivo)
	if err != nil {
		if os.IsNotExist(err) {
			os.Create(arquivo)
		}
	} else {
		fmt.Println("arquivo existente")
	}
}

func rmfile(arquivo string) { //----------------------------------validado
	file, err := os.Open(arquivo)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("arquivo inexistente")
		}
	} else {
		fi, _ := file.Stat()
		if fi.IsDir() {
			fmt.Println("isto nao é um arquivo")
		} else {
			os.Remove(arquivo)
		}
	}
}

func copy(origem, destino string) { //----------------------------------validado
	src, _ := os.Stat(origem)
	_, err := os.Open(destino)
	if err != nil { //destino nao existe
		if !src.IsDir() {
			CopyFile(origem, destino)
		} else {
			CopyDir(origem, destino)
		}

	} else {
		next, _ := os.Stat(destino)
		if !src.IsDir() && !next.IsDir() {
			CopyFile(origem, destino)
		} else if !src.IsDir() && next.IsDir() {
			CopyFile(origem, destino+"/"+origem)
		} else {
			CopyDir(origem, destino)
		}
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
	fmt.Print("\033[H\033[2J")
}

func locate(nome string) {
	var paf string
	fnd := false
	dir, _ := os.Getwd()
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, nome) {
			paf = path
			fnd = true
			return io.EOF
		}
		return nil
	})

	if fnd {
		fmt.Println("arquivo encontrado em \n" + paf)
	} else {
		fmt.Println("arquivo nao encontrado")
	}
}

func validacao(palavras ...string) bool {
	for _, val := range palavras {
		if val == "" {
			fmt.Println("parametros invalidos")
			return false
		}
	}
	return true
}

func qntparams(num, tam int) bool {
	if tam > num { //len() > num
		fmt.Println("quantidade de parametros invalida")
		return false
	}
	return true
}

func selecionaComando(entrada []string) {
	str := entrada[0]
	str2, str3 := "", ""
	if len(entrada) > 1 {
		str2 = entrada[1]
	}
	if len(entrada) > 2 {
		str3 = entrada[2]
	}

	switch str {
	case "cd":
		if qntparams(2, len(entrada)) {
			cd(str2)
		}
	case "ls":
		if qntparams(9, len(entrada)) {
			ls(entrada[1:])
		}
	case "mv":
		if validacao(str2, str3) && qntparams(3, len(entrada)) {
			mv(str2, str3)
		}
	case "cat":
		if validacao(str2) && qntparams(2, len(entrada)) {
			cat(str2)
		}
	case "man":
		if qntparams(2, len(entrada)) {
			man(str2)
		}
	case "mkdir":
		if validacao(str2) && qntparams(2, len(entrada)) {
			mkdir(str2)
		}
	case "rmdir":
		if validacao(str2) && qntparams(2, len(entrada)) {
			rmdir(str2)
		}
	case "clear":
		if qntparams(1, len(entrada)) {
			clear()
		}
	case "locate":
		if validacao(str2) && qntparams(2, len(entrada)) {
			locate(str2)
		}
	case "rmfile":
		if validacao(str2) && qntparams(2, len(entrada)) {
			rmfile(str2)
		}
	case "mkfile":
		if validacao(str2) && qntparams(2, len(entrada)) {
			mkfile(str2)
		}
	case "copy":
		if validacao(str2, str3) && qntparams(3, len(entrada)) {
			copy(str2, str3)
		}
	case "":
		fmt.Printf("")
	default:
		fmt.Println("comando invalido")
	}
}

func main() {
	//descobrir nome da pasta Desktop / Area de Trabalho
	usr, _ := user.Current()
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

		cmd := strings.Split(strings.Replace(s, "\\ ", "\\", len(s)), " ") //separa os argumentos
		selecionaComando(cmd)

		dir2, _ := os.Getwd()
		fmt.Printf(dir2 + "$ ")
	}
	if err := scanner.Err(); err != nil {
		os.Exit(1)
	}
}

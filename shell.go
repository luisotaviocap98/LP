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
	"regexp"
	"runtime"
	"sort"
	"strings"
)

var casa string

func manipulate(rota string) ([]string, string, []string, bool, bool) {
	x, n := false, true

	f := strings.Split(strings.Replace(rota, " ", "", len(rota)), "\\")
	j := f[0]

	if j == rota { //caminho nao possui nome composto
		n = false
	}

	for i := 1; i < len(f); i++ {
		j += " " + f[i]
	}

	p := strings.Split(j, "/") //separar niveis de diretorio

	return f, j, p, x, n
}

func cd(caminho string) { //----------------------------------validado
	f, j, p, x, n := manipulate(caminho)
	muda := false
	myd, _ := os.Getwd()
	arq, _ := ioutil.ReadDir(myd)

	for i := 0; i < len(arq); i++ {
		if p[0] == "" { //diretorio "/dir", pode estar no diretorio atual ou ser outro diretorio
			if len(p) > 1 && strings.HasSuffix(arq[i].Name(), p[1]) { //identifica se o caminho desejado existe no diretorio atual
				muda = false
				x = true
				break
			} else {
				muda = true
			}
		} else if n { //verifica se o diretorio tem nome composto

			if strings.HasPrefix(arq[i].Name(), f[0]) || strings.HasPrefix(arq[i].Name(), p[0]) { //identifica se o caminho desejado existe no diretorio atual
				muda = false
			} else {
				muda = true
			}
		} else {
			if strings.HasSuffix(arq[i].Name(), p[0]) { //identifica se o caminho desejado existe no diretorio atual
				muda = false
				x = true
				break
			} else {
				muda = true
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
		if muda {
			if n {
				new := os.Chdir(j)
				if new != nil {
					fmt.Println("caminho inexistente")
				}
			} else {
				new := os.Chdir(caminho)
				if new != nil {
					fmt.Println("caminho inexistente")
				}
			}
		} else {
			if n {
				new := os.Chdir(myd + "/" + j)
				if new != nil {
					fmt.Println("caminho inexistente")
				}
			} else {
				new := os.Chdir(myd + "/" + caminho)
				if new != nil {
					fmt.Println("caminho inexistente")
				}
			}
		}
	}
}

func leftjust(s string, n int, fill string) string {
	return s + strings.Repeat(fill, n)
}

func recursiveParam(original []os.FileInfo, files, parametro []string) { //remover recursivo os parametros

	var str []string

	// verifico a quantidade de parâmetros útil
	qtdParamUtil := 0
	for i := 0; i < len(parametro); i++ {
		if parametro[i] != "" {
			qtdParamUtil += 1
		}
	}

	liberado := false

	if qtdParamUtil > 2 {
		fmt.Println("quantidade de parâmetros inválidos")
		return
	} else if qtdParamUtil == 2 {
		if parametro[0] == "-sortasc" || parametro[1] == "-sortasc" || parametro[0] == "-sortdesc" || parametro[1] == "-sortdesc" {
			liberado = true
		} else {
			fmt.Println("combinação de parametros não permitida")
		}
	} else {
		liberado = true
	}
	if liberado {
		// ls padrão ou ls com ordenação
		if qtdParamUtil == 0 || (qtdParamUtil == 1 && (parametro[0] == "-sortasc" || parametro[0] == "-sortdesc")) { // padrão, lista todos exceto ocultas
			if qtdParamUtil == 1 {
				if parametro[0] == "-sortasc" {
					sort.Slice(original, func(i, j int) bool { return strings.ToLower(original[i].Name()) < strings.ToLower(original[j].Name()) })
				}

				if parametro[0] == "-sortdesc" {
					sort.Slice(original, func(i, j int) bool { return strings.ToLower(original[i].Name()) > strings.ToLower(original[j].Name()) })
				}
			}
			for i := 0; i < len(original); i++ {
				match, _ := regexp.MatchString("^\\..*", original[i].Name())
				if !match {
					str = append(str, original[i].Name()) // adiciona oque ja tinha, mais os dados novos
				}
			}
		} else { // se não for somente ls básico, então faremos tratamento de ls complexo

			for i := 0; i < len(parametro); i++ {

				// valid simples e com ordenação
				if parametro[i] == "-valid" { //  não lista entradas implícitas (. e ..)

					for i := 0; i < len(original); i++ {
						if original[i].Name() != "." || original[i].Name() != ".." {
							str = append(str, original[i].Name()) // adiciona oque ja tinha, mais os dados novos
						}
					}

					for p := 0; p < len(parametro); p++ {
						if parametro[p] == "-sortasc" {
							sort.Slice(str, func(i, j int) bool { return strings.ToLower(str[i]) < strings.ToLower(str[j]) })
						}
						if parametro[p] == "-sortdesc" {
							sort.Slice(str, func(i, j int) bool { return strings.ToLower(str[i]) > strings.ToLower(str[j]) })
						}
					}
				} else if parametro[i] == "-hidden" { //  lista entradas ocultas

					str = append(str, ".")  // adiciona oque ja tinha, mais os dados novos
					str = append(str, "..") // adiciona oque ja tinha, mais os dados novos

					for i := 0; i < len(original); i++ {
						str = append(str, original[i].Name()) // adiciona oque ja tinha, mais os dados novos
					}

					for p := 0; p < len(parametro); p++ {
						if parametro[p] == "-sortasc" {
							sort.Slice(str, func(i, j int) bool { return strings.ToLower(str[i]) < strings.ToLower(str[j]) })
						}
						if parametro[p] == "-sortdesc" {
							sort.Slice(str, func(i, j int) bool { return strings.ToLower(str[i]) > strings.ToLower(str[j]) })
						}
					}
				} else if parametro[i] == "-dirs" { // lista somente diretórios

					for i := 0; i < len(original); i++ {
						if original[i].IsDir() {
							str = append(str, original[i].Name()) // adiciona oque ja tinha, mais os dados novos
						}
					}

					for p := 0; p < len(parametro); p++ {
						if parametro[p] == "-sortasc" {
							sort.Slice(str, func(i, j int) bool { return strings.ToLower(str[i]) < strings.ToLower(str[j]) })
						}
						if parametro[p] == "-sortdesc" {
							sort.Slice(str, func(i, j int) bool { return strings.ToLower(str[i]) > strings.ToLower(str[j]) })
						}
					}
				} else if parametro[i] == "-files" { //  lista somente arquivos

					for i := 0; i < len(original); i++ {
						if !original[i].IsDir() {
							str = append(str, original[i].Name()) // adiciona oque ja tinha, mais os dados novos
						}
					}

					for p := 0; p < len(parametro); p++ {
						if parametro[p] == "-sortasc" {
							sort.Slice(str, func(i, j int) bool { return strings.ToLower(str[i]) < strings.ToLower(str[j]) })
						}
						if parametro[p] == "-sortdesc" {
							sort.Slice(str, func(i, j int) bool { return strings.ToLower(str[i]) > strings.ToLower(str[j]) })
						}
					}
				} else if parametro[i] == "-full" { //  lista todas as propriedades das entradas (date, size, owner,. . . )

					for p := 0; p < len(parametro); p++ {
						if parametro[p] == "-sortasc" {
							sort.Slice(original, func(i, j int) bool { return strings.ToLower(original[i].Name()) < strings.ToLower(original[j].Name()) })
						}

						if parametro[p] == "-sortdesc" {
							sort.Slice(original, func(i, j int) bool { return strings.ToLower(original[i].Name()) > strings.ToLower(original[j].Name()) })
						}
					}

					for i := 0; i < len(original); i++ {
						match, _ := regexp.MatchString("^\\..*", original[i].Name())
						if !match {
							fmt.Print(original[i].Mode()) // permissões
							fmt.Print("      ")
							fmt.Print(original[i].ModTime()) // ultima modificação (tempo)
							fmt.Print("      ")
							fmt.Print(original[i].Size()) // tamanho do arquivo
							fmt.Print("      ")
							fmt.Println(original[i].Name()) // nome do arquivo
						}
					}
				}
			}
		}
	}
	// imprime os resultados
	if len(str) > 0 {
		imprimir(str)
	}
}

func imprimir(dados []string) {

	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, _ := cmd.Output()

	//espacar nomes listados
	x, tam := 1, 0
	for i := len(out) - 2; i >= 3; i-- {
		tam += (int(out[i]) - 48) * x
		x = x * 10
	}

	var copy []string

	alert := 0

	// elimina dados duplicados
	for i := 0; i < len(dados); i++ {
		alert = 0
		for j := 0; j < len(copy); j++ {
			if dados[i] == copy[j] {
				alert = 1
			}
		}
		if alert == 0 {
			copy = append(copy, dados[i])
		}
	}

	if len(copy) > 0 {
		cont := 0
		for i := 0; i < len(copy); i++ {
			st := copy[i]
			cont += len(copy[i]) + 5 //verifica se o nome do arquivo cabe na tela
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

func ls(parametro []string) { //----------------------------------validado

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
	} else {

		recursiveParam(arquivos, archs, parametro)
	}

}

func mv(origem, destino string) { //----------------------------------validado
	if copy(origem, destino) {
		os.RemoveAll(origem)
	}
}

func cat(arquivo string) { //----------------------------------validado
	_, arquivo, _, _, _ = manipulate(arquivo)
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
	_, pasta, _, _, _ = manipulate(pasta)
	_, err := os.Open(pasta)
	if err != nil {
		if os.IsNotExist(err) {
			newpath := filepath.Join(pasta, "")
			os.MkdirAll(newpath, os.ModePerm)
		}
	} else {
		fmt.Println("arquivo existente, deseja continuar? [s | n]")
		reader := bufio.NewReader(os.Stdin)
		escolha, _ := reader.ReadString('\n')
		escolha = strings.Replace(escolha, "\n", "", -1)
		if escolha == "n" {
			return
		} else if escolha == "s" {
			newpath := filepath.Join(pasta, "")
			os.MkdirAll(newpath, os.ModePerm)
		} else {
			fmt.Println("opcao invalida")
		}
	}

}

func rmdir(pasta string) { //----------------------------------validado
	_, pasta, _, _, _ = manipulate(pasta)
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
	_, arquivo, _, _, _ = manipulate(arquivo)
	_, err := os.Stat(arquivo)
	if err != nil {
		if os.IsNotExist(err) {
			os.Create(arquivo)
		}
	} else {
		fmt.Println("arquivo existente, deseja continuar? [s | n]")
		reader := bufio.NewReader(os.Stdin)
		escolha, _ := reader.ReadString('\n')
		escolha = strings.Replace(escolha, "\n", "", -1)
		if escolha == "n" {
			return
		} else if escolha == "s" {
			os.Create(arquivo)
		} else {
			fmt.Println("opcao invalida")
		}
	}
}

func rmfile(arquivo string) { //----------------------------------validado
	_, arquivo, _, _, _ = manipulate(arquivo)
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

func copy(origem, destino string) bool { //----------------------------------validado
	var sucesso error
	_, origem, _, _, _ = manipulate(origem)
	_, destino, _, _, _ = manipulate(destino)
	_, erro := os.Open(origem)
	if erro != nil {
		fmt.Println("origem inexistente")
	} else {
		src, _ := os.Stat(origem)

		_, err := os.Open(destino)
		if err != nil { //destino nao existe
			if !src.IsDir() {
				sucesso = CopyFile(origem, destino)
			} else {
				sucesso = CopyDir(origem, destino)
			}

		} else {
			fmt.Println("arquivo destino ja existente, deseja continuar? [s | n]")
			reader := bufio.NewReader(os.Stdin)
			escolha, _ := reader.ReadString('\n')
			escolha = strings.Replace(escolha, "\n", "", -1)
			if escolha == "n" {
				return false
			} else if escolha == "s" {
				next, _ := os.Stat(destino)
				if !src.IsDir() && !next.IsDir() {
					sucesso = CopyFile(origem, destino)
				} else if !src.IsDir() && next.IsDir() {
					sucesso = CopyFile(origem, destino+"/"+origem)
				} else {
					sucesso = CopyDir(origem, destino)
				}
			} else {
				fmt.Println("opcao invalida")
			}

		}
	}
	if sucesso == nil {
		return true
	} else {
		return false
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
	_, nome, _, _, _ = manipulate(nome)
	var paf string
	fnd := false
	dir, _ := os.Getwd()
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		_, x, _, _, _ := manipulate(path)
		j := strings.Split(x, "/")

		if j[len(j)-1] == nome {
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
	if len(os.Args) > 1 {
		arg := os.Args[1]
		cd(arg) //iniciar em um diretorio especifico
	}
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

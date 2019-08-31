package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	// "strings"
	// "bufio"
)

type Comando struct {
	nome       string
	operacao   func()
	parametros [10]string
	manual     string
}

// tatica: usar map contendo vetor de comandos e comparar com vetor de struct

func main() {
	/*
		// pegar o comando digitado
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			s := scanner.Text() //eh uma string

			if s == "exit" {
				os.Exit(1)
			}

			j := strings.Split(s, " ")
			fmt.Println(j, len(j), "1°", j[0])
			//%q outra opcao

			for _, i := range j { //_ é index, i eh o valor
				fmt.Printf("%s\n", i)
			}
		}
		if err := scanner.Err(); err != nil {
			os.Exit(1)
		}
	*/

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
	fmt.Println(dir)

	// cat
	// content, _ := ioutil.ReadFile(arquivos[2].Name())
	// fmt.Printf("File contents:\n%s", content)

	// mkdir
	// err := ioutil.WriteFile("testdata", []byte(""), 0644)
	// if err != nil {
	// log.Fatal(err)
	// }

	// mv
	// nomearq := "primeiro.go"
	// err := os.Rename("../"+nomearq, "../LP/"+nomearq)
	// if err != nil {
	// fmt.Println(err)
	// }

	// rmdir
	// os.Remove("./xuxu")

	// clear
	// for i := 0; i < 30; i++ {
	// fmt.Println("")
	// }

	// man
	// os.Open(Comando.manual)

	// locate

	// cd
	// os.Chdir("/home/luiscap")
	// mydir, err := os.Getwd()
	// if err == nil {
	// fmt.Println("novo dir:", mydir)
	// }
	// arquiv, erro := ioutil.ReadDir(mydir)
	// for i := 0; i < len(arquiv); i++ {
	// st := arquiv[i].Name() + "   "
	// fmt.Printf(st)
	// }

	/*
		// mapeamento dos comandos
		m := map[int]string{
			0: "cd",
			1: "ls",
			2: "mv",
			3: "cat",
			4: "man",
			5: "mkdir",
			6: "rmdir",
			7: "clear",
			8: "locate",
		}

		str := "ls"
		for _, j := range m {
			if str == j {
				println("bateu", j)
			}
		}
	*/

	/*
		//vetor dos comandos
		const cmds []Comando{
			Comando{
				nome: "cd"
				operacao:
				parametros:
				manual: "cd.txt"
			},
			Comando{
				nome: "ls"
				operacao:
				parametros:
				manual: "ls.txt"
			},
			Comando{
				nome: "mv"
				operacao:
				parametros:
				manual: "mv.txt"
			},
			Comando{
				nome: "cat"
				operacao:
				parametros:
				manual: "cat.txt"
			},
			Comando{
				nome: "man"
				operacao:
				parametros:
				manual: "man.txt"
			},
			Comando{
				nome: "mkdir"
				operacao:
				parametros:
				manual: "mkdir.txt"
			},
			Comando{
				nome: "rmdir"
				operacao:
				parametros:
				manual: "rmdir.txt"
			},
			Comando{
				nome: "clear"
				operacao:
				parametros:
				manual: "clear.txt"
			},
			Comando{
				nome: "locate"
				operacao:
				parametros:
				manual: "locate.txt"
			}
		}
	*/
}

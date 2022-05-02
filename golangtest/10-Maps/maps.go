package main

import "fmt"

func main() {
	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario)

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Joao",
			"ultimo":   "pedro",
		},
		"curso": {
			"nome":   "engenharia",
			"campos": "campos1",
		},
	}
	fmt.Println(usuario2)
}

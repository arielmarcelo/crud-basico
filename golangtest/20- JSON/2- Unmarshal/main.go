package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raça  string `json:"raça"`
	Idade uint8  `json:"idade"`
}

func main() {
	cachorroEmJSON := `{"nome":"Rex","raça":"Vira Lata","idade":3}`

	var c cachorro

	if erro := json.Unmarshal([]byte(cachorroEmJSON), &c); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c)

	cachorro2EmJSON := `{"nome":"Snoop","raça":"Shou Shou"}`

	c2 := make(map[string]string)

	if erro := json.Unmarshal([]byte(cachorro2EmJSON), &c2); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c2)

}

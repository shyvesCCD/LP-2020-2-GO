package main

import (
	"io"
	"log"
	"os"
	"time"
)

func main() {
	now := time.Now()

	teamsFile, err := os.OpenFile("teams.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Erro em abrir o arquivo teams.txt, retornando o seguinte erro: %v", err)
	}
	defer teamsFile.Close()

	write := io.MultiWriter(os.Stdout, teamsFile)

	log.SetOutput(write)

	teams, err := nbaApi.getAllTheTeams()
	if err != nil {
		log.Fatalf("Erro durante a requisicoes de todos os times: %v", err)
	}

	for _, team := range teams {
		log.Printf("Name %s", team.Name)
		log.Println("------------Divisão------------")
	}

	log.Printf("Levou %v para fazer todas as requisicoes.", time.Now().Sub(now).String())
}

package main

import (
	"LPProject/nhlApi"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

func main() {
	// ajudar a comparar o tempo de solicitação
	now := time.Now()

	rosterFile, err := os.OpenFile("rosters.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("error opening the file rosters.txt: %v", err)
	}
	defer rosterFile.Close()

	wrt := io.MultiWriter(os.Stdout, rosterFile)

	log.SetOutput(wrt)

	var wg sync.WaitGroup

	teams, err := nhlApi.GetAllTeams()
	if err != nil {
		log.Fatalf("error while getting all teams: %v", err)
	}

	wg.Add(len(teams))

	// canal sem buffer
	results := make(chan []nhlApi.Roster)

	for _, team := range teams {
		go func(team nhlApi.Team) {
			roster, err := nhlApi.GetRosters(team.ID)
			if err != nil {
				log.Fatalf("error getting roster: %v", err)
			}
			results <- roster
			wg.Done()
		}(team)

	}

	go func() {
		wg.Wait()
		close(results)
	}()

	display(results)

	log.Printf("took %v", time.Since(now).String())
}

func display(results chan []nhlApi.Roster) {
	var entrada string
	fmt.Println("POSIÇÕES EXISTENTES:")
	fmt.Println("[D,C,G,RW,LW]")
	fmt.Println("ESCOLHA A POSIÇÃO DO JOGADOR:")
	fmt.Scan(&entrada)
	for r := range results {
		for _, x := range r {
			if entrada == x.Position.Abbreviation {
				log.Println("----------------------")
				log.Printf("ID: %d\n", x.Person.ID)
				log.Printf("Name: %s\n", x.Person.FullName)
				log.Printf("Position: %s\n", x.Position.Abbreviation)
				log.Printf("Jersey: %s\n", x.JerseyNumber)
				log.Println("----------------------")
			}
		}
	}
}

package main

import (
	"fmt"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {
	err := tickets.SetTicketList("tickets.csv")
	if err != nil {
		fmt.Println(err)
	}

	count, err := tickets.GetTotalTickets("Brazil")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("El número de personas que viajan a Brazil es", count)

	count, err = tickets.GetEarlyMorningTickets()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("El número de personas que viajan de madrugada es", count)

	count, err = tickets.GetLateMorningTickets()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("El número de personas que viajan en la mañana es", count)

	count, err = tickets.GetAfternoonTickets()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("El número de personas que viajan en la tarde es", count)

	count, err = tickets.GetNightTickets()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("El número de personas que viajan de noche es", count)

}

package tickets

import (
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
)

type Ticket struct {
	ID                     int
	Nombre, Email, Destino string
	Hora                   string
	Precio                 float64
}

var TicketList = []Ticket{}

func SetTicketList(fileName string) error {

	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	bytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	content := string(bytes)

	for _, line := range strings.Split(content, "\r\n") {
		//Omitir lineas vacias
		if len(line) < 1 {
			break
		}

		data := strings.Split(line, ",")

		id, err := strconv.Atoi(data[0])
		if err != nil {
			return err
		}

		precio, err := strconv.ParseFloat(data[5], 64)
		if err != nil {
			return err
		}

		TicketList = append(TicketList, Ticket{
			ID:      id,
			Nombre:  data[1],
			Email:   data[2],
			Destino: data[3],
			Hora:    data[4],
			Precio:  precio,
		})
	}
	return nil
}

// ejemplo 1
func GetTotalTickets(destination string) (int, error) {
	var count int

	for _, ticket := range TicketList {
		if ticket.Destino == destination {
			count++
		}
	}
	return count, nil
}

// ejemplo 2
func GetEarlyMorningTickets() (int, error) {
	var count int

	for _, ticket := range TicketList {
		hora := strings.Split(ticket.Hora, ":")[0]
		numHora, err := strconv.Atoi(hora)
		if err != nil {
			return 0, err
		}
		if numHora >= 0 && numHora <= 6 {
			count++
		}
	}
	return count, nil
}

func GetLateMorningTickets() (int, error) {
	var count int

	for _, ticket := range TicketList {
		hora := strings.Split(ticket.Hora, ":")[0]
		numHora, err := strconv.Atoi(hora)
		if err != nil {
			return 0, err
		}
		if numHora >= 7 && numHora <= 12 {
			count++
		}
	}
	return count, nil
}

func GetAfternoonTickets() (int, error) {
	var count int

	for _, ticket := range TicketList {
		hora := strings.Split(ticket.Hora, ":")[0]
		numHora, err := strconv.Atoi(hora)
		if err != nil {
			return 0, err
		}
		if numHora >= 13 && numHora <= 19 {
			count++
		}
	}
	return count, nil
}

func GetNightTickets() (int, error) {
	var count int

	for _, ticket := range TicketList {
		hora := strings.Split(ticket.Hora, ":")[0]
		numHora, err := strconv.Atoi(hora)
		if err != nil {
			return 0, err
		}
		if numHora >= 20 && numHora <= 23 {
			count++
		}
	}
	return count, nil
}

func GetCountByPeriod(time string) (int, error) {
	switch time {
	case "madrugada":
		return GetEarlyMorningTickets()
	case "mañana":
		return GetLateMorningTickets()
	case "tarde":
		return GetAfternoonTickets()
	case "noche":
		return GetNightTickets()
	default:
		return 0, errors.New("No se encontró el periodo especificado")
	}
}

// ejemplo 3
func GetAverageDestination(destination string, total int) (float64, error) {
	destinationTickets, err := GetTotalTickets(destination)
	if err != nil {
		return 0, err
	}
	result := float64(destinationTickets) / float64(total)
	return result, nil
}

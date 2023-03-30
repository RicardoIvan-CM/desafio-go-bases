package tickets

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	SetTicketList("../../tickets.csv")
	code := m.Run()
	os.Exit(code)
}

func TestGetTotalTickets(t *testing.T) {
	t.Run("Brazil", func(t *testing.T) {
		//Arrange
		var destination string = "Brazil"
		var expectedResult int = 45

		//Act
		result, err := GetTotalTickets(destination)

		//Assert
		assert.NoError(t, err, "Existe un error")
		assert.Equal(t, expectedResult, result, "Los resultados no coinciden")
	})

	t.Run("Mexico", func(t *testing.T) {
		//Arrange
		var destination string = "Mexico"
		var expectedResult int = 11

		//Act
		result, err := GetTotalTickets(destination)

		//Assert
		assert.NoError(t, err, "Existe un error")
		assert.Equal(t, expectedResult, result, "Los resultados no coinciden")
	})
}

func TestGetCountByTimePeriod(t *testing.T) {
	t.Run("Madrugada", func(t *testing.T) {
		//Arrange
		var expectedResult int = 304

		//Act
		result, err := GetCountByPeriod("madrugada")

		//Assert
		assert.NoError(t, err, "Existe un error")
		assert.Equal(t, expectedResult, result, "Los resultados no coinciden")
	})

	t.Run("Mañana", func(t *testing.T) {
		//Arrange
		var expectedResult int = 256

		//Act
		result, err := GetCountByPeriod("mañana")

		//Assert
		assert.NoError(t, err, "Existe un error")
		assert.Equal(t, expectedResult, result, "Los resultados no coinciden")
	})

	t.Run("Tarde", func(t *testing.T) {
		//Arrange
		var expectedResult int = 289

		//Act
		result, err := GetCountByPeriod("tarde")

		//Assert
		assert.NoError(t, err, "Existe un error")
		assert.Equal(t, expectedResult, result, "Los resultados no coinciden")
	})

	t.Run("Noche", func(t *testing.T) {
		//Arrange
		var expectedResult int = 151

		//Act
		result, err := GetCountByPeriod("noche")

		//Assert
		assert.NoError(t, err, "Existe un error")
		assert.Equal(t, expectedResult, result, "Los resultados no coinciden")
	})

	t.Run("Dato inválido", func(t *testing.T) {
		//Arrange
		var expectedResult int = 0

		//Act
		result, err := GetCountByPeriod("medianoche")

		//Assert
		assert.EqualError(t, err, "No se encontró el periodo especificado", "Los errores no coinciden")
		assert.Equal(t, expectedResult, result, "Los resultados no coinciden")
	})
}

func TestAverageDestination(t *testing.T) {
	t.Run("Brazil", func(t *testing.T) {
		//Arrange
		var destination string = "Brazil"
		var total int = 1000
		var expectedResult float64 = 0.045

		//Act
		result, err := GetAverageDestination(destination, total)

		//Assert
		assert.NoError(t, err, "Existe un error")
		assert.Equal(t, expectedResult, result, "Los resultados no coinciden")
	})

	t.Run("Mexico", func(t *testing.T) {
		//Arrange
		var destination string = "Mexico"
		var total int = 1000
		var expectedResult float64 = 0.011

		//Act
		result, err := GetAverageDestination(destination, total)

		//Assert
		assert.NoError(t, err, "Existe un error")
		assert.Equal(t, expectedResult, result, "Los resultados no coinciden")
	})
}

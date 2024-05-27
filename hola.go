package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

func main() {

	var x_train, y_train []float64

	resp, err := http.Get("https://raw.githubusercontent.com/DiegoAlejandroGarciaArzate/Programaci-n-Concurrente/main/taxi_data.csv")
	if err != nil {
		fmt.Println("error descarga csv \n", err)
		return
	}
	defer resp.Body.Close()

	reader := csv.NewReader(resp.Body)

	data, err := reader.ReadAll()
	if err != nil {
		fmt.Println("error abrir csv \n", err)
		return
	}

	fmt.Println("tamaño data:  ", len(data))
	fmt.Println("Columns: ", data[0][0], "  ", data[0][1])

	n := len(data) - 1
	for i := 1; i < n; i++ {
		trip, _ := strconv.ParseFloat(data[i][0], 64)
		total, _ := strconv.ParseFloat(data[i][1], 64)

		x_train = append(x_train, trip)
		y_train = append(y_train, total)
	}

	fmt.Println("Tamaño arreglos: x = ", len(x_train), "; y = ", len(y_train))

	slope, intercept := calculateRegression(x_train, y_train)
	fmt.Printf("Slope: %.2f, Intercept: %.2f\n", slope, intercept)

	// Canal para resultados
	results := make(chan string, 1000)

	// Grupo de espera para goroutines
	var wg sync.WaitGroup

	// Calcular y mostrar los valores de y para 1000 diferentes valores de x usando goroutines
	initialX := 1.0
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			x := initialX + float64(i)
			y := slope*x + intercept
			results <- fmt.Sprintf("Para x = %.2f, el valor predicho de y es: %.2f\n", x, y)
		}(i)
	}

	// Cerrar el canal cuando todas las goroutines terminen
	go func() {
		wg.Wait()
		close(results)
	}()

	// Recibir y mostrar resultados del canal
	for result := range results {
		fmt.Print(result)
	}
}

func calculateRegression(x, y []float64) (float64, float64) {
	n := float64(len(x))
	var sumX, sumY, sumXY, sumX2 float64

	for i := 0; i < int(n); i++ {
		sumX += x[i]
		sumY += y[i]
		sumXY += x[i] * y[i]
		sumX2 += x[i] * x[i]
	}

	slope := (n*sumXY - sumX*sumY) / (n*sumX2 - sumX*sumX)
	intercept := (sumY - slope*sumX) / n

	return slope, intercept
}


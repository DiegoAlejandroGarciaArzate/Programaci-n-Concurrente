package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"
)

func main() {

	var x_train, y_train []float64

	resp, err := http.Get("https://raw.githubusercontent.com/DiegoAlejandroGarciaArzate/Programaci-n-Concurrente/main/taxi_data.csv")
	defer resp.Body.Close()

	if err != nil {
		fmt.Println("error descarga csv \n", err)
		return
	}

	reader := csv.NewReader(resp.Body)

	data, err := reader.ReadAll()

	if err != nil {
		fmt.Println("error abrir csv? \n", err)
		return
	}

	fmt.Println("tamaño data:  ", len(data))

	fmt.Println("Columns: ", data[0][0], "  ", data[0][1])

	n := len(data) - 1
	//Carga valores, si se hacen pruebas cambiar el len(data) a un número como 10
	for i := 1; i < n; i++ {
		trip, _ := strconv.ParseFloat(data[i][0], 64)
		total, _ := strconv.ParseFloat(data[i][1], 64)

		x_train = append(x_train, trip)
		y_train = append(y_train, total)
	}

	//solo imprimir si el el valor de n es pequeño
	/*for i := 0; i < len(x_train); i++ {
		fmt.Println(x_train[i], y_train[i])

	} */

	fmt.Println("Tamñano arreglos: x = ", len(x_train), ";  y = ", len(y_train))
}

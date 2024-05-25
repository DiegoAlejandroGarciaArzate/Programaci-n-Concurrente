package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("HOLA MUNDO")

	//resp, err := http.Get("https://raw.githubusercontent.com/iller15/TP-Concurrente/main/Secuencial.go")
	resp, err := http.Get("https://github.com/DiegoAlejandroGarciaArzate/Programaci-n-Concurrente/blob/main/data/AEP_hourly.csv")
	defer resp.Body.Close()

	if err != nil {
		fmt.Println("error \n", err)
		return
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("error \n", err)
		return
	}
	fmt.Println(body[0])
}

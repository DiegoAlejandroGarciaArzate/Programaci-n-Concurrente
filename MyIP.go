package main

import (
	"log"
	"net"
	"strings"
)

func localladdress() {
	ifaces, _ := net.Interfaces()

	for _, i := range ifaces { //pasa por los interfaces [todos en forma de arreglo]
		addrs, _ := i.Addrs()

		for _, a := range addrs { //direcciones ip
			log.Printf("%v %v\n", i.Name, a)
		}
	}

}

func descubrirIP() string {
	//interfas de red
	ifaces, _ := net.Interfaces()
	for _, i := range ifaces { //interfaces de red
		if strings.HasPrefix(i.Name, "Ethernet") {
			//solo si tiene "Ethernet" al inicio
			addrs, _ := i.Addrs()
			for _, addr := range addrs {
				switch t := addr.(type) {
				case *net.IPNet: //si el tipo es ip? o es en caso de emergencia
					if t.IP.To4() != nil {
						return t.IP.To4().String() //retornamos la ip ethernet version 4
					} //if el tipo de ip es ipv4
				}
			}
		}
	}
	return "127.0.0.1" //default ip de un localhost si no está conectado
}

func main() {
	//localladdress()
	println("Mi IP es: ", descubrirIP())
}

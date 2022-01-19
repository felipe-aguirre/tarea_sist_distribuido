package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	pb "github.com/felipe-aguirre/tarea_sist_distribuido/protos"
	"google.golang.org/grpc"
)

// IP Del Broker
const (
	BrokerAddress = "localhost:50050"
)
var vector = "None"
var ServerIP = "None"
// AddCity nombre_planeta nombre_ciudad [nuevo_valor]
// Esto creará una nueva línea en el registro planetario correspondiente.
// Si dicho planeta aún no posee un archivo de registro planetario debe crearse
// uno.
// Este comando puede o no ingresarse con el nuevo valor. En caso de no escribirse uno,
// debe guardarse esa ciudad con valor 0.

// FORMATO Registro Planetario (1 por planeta)
// nombre_planeta nombre_ciudad cantidad_soldados_rebeldes
// Ejemplo:
// Tatooine Mos_Eisley 5

func main() {


		exit := false
		loop := true
		for exit != loop {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Leia > ")
			text, _ := reader.ReadString('\n')
			text = strings.Replace(text, "\r\n", "", -1)
			respuesta := strings.Split(text, " ")
			if (respuesta[0] == "GetNumberRebelds"){
				// Conexión a Broker
				conn, err := grpc.Dial(BrokerAddress, grpc.WithInsecure(), grpc.WithBlock())
				if err != nil {
					log.Fatalf("did not connect: %v", err)
				}
				c := pb.NewManejoComunicacionClient(conn)

				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()
				r, _ := c.Comunicar(ctx, &pb.MessageRequest{Request: text, Autor: "Leia"})
				log.Printf(`Mensaje recibido del Broker: %s`, r.GetReply())
				log.Printf("Contactando al servidor %s con el comando '%s'", r.GetReply(), text)
				respuesta := r.GetReply()
				linea_leida := strings.Split(respuesta, ",")
				vector = linea_leida[0]
				log.Printf("Vector recibido: %s", vector)
				conn.Close()
				
			} else {
				if strings.Compare("exit", text) == 0 {
					exit = true
				} else {
					fmt.Println("Comando erroneo, intente con GetNumberRebelds")
				}
			}
		}
	
	

}

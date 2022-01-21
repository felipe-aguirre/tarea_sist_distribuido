package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"time"

	pb "github.com/felipe-aguirre/tarea_sist_distribuido/protos"
	"google.golang.org/grpc"
)

var Reset  = "\033[0m"
var Red    = "\033[31m"
var Green  = "\033[32m"
var Yellow = "\033[33m"
var Blue   = "\033[34m"
var Purple = "\033[35m"
var Cyan   = "\033[36m"
var Gray   = "\033[37m"
var White  = "\033[97m"

// IP Del Broker
const (
	BrokerAddress = "localhost:13370"
)

//Vector en el formato
// { 
// 	nombrePlaneta: vector
// }
var vector = map[string]string{}


// Server IP en formato
// {
// 	nombrePlaneta: IP
// }
var ServerIP = map[string]string{}



func main() {
	if runtime.GOOS == "windows" {
		Reset  = ""
		Red    = ""
		Green  = ""
		Yellow = ""
		Blue   = ""
		Purple = ""
		Cyan   = ""
		Gray   = ""
		White  = ""
	}

		exit := false
		loop := true
		for exit != loop {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print(Green + "Leia > "+ Reset)

			text, _ := reader.ReadString('\n')
			if runtime.GOOS == "windows" {
				text = strings.Replace(text, "\r\n", "", -1)
			} else {
				text = strings.Replace(text, "\n", "", -1)
			} 
			respuesta := strings.Split(text, " ")
			if (respuesta[0] == "GetNumberRebelds"){

				if (len(respuesta) != 3) {
					fmt.Println(Red + "ERROR: "+ Yellow + "GetNumberRebelds" + Reset +  " debe tener tener el formato:")
					fmt.Println("GetNumberRebelds nombre_planeta nombre_ciudad")
					fmt.Println("")
				} else{
					// Conexión a Broker
					log.Printf("")
					fmt.Println("Conectando a Broker en " + Yellow + BrokerAddress + Reset)

					conn, err := grpc.Dial(BrokerAddress, grpc.WithInsecure(), grpc.WithBlock())
					if err != nil {
						log.Fatalf("did not connect: %v", err)
					}
					c := pb.NewManejoComunicacionClient(conn)

					ctx, cancel := context.WithTimeout(context.Background(), time.Second)
					defer cancel()
					r, _ := c.Comunicar(ctx, &pb.MessageRequest{Request: text, Autor: "Leia"})
					
					respuestaBroker := r.GetReply()
					linea_leida := strings.Split(respuestaBroker, ",")
					RespuestaVector := linea_leida[0]
					respuestaPlanetas := linea_leida[1]
					IPrecibida := linea_leida[2]
					if respuestaPlanetas != "-1" {
						fmt.Println("La ciudad " + Yellow +  respuesta[2] + Reset + " del planeta " + Yellow + respuesta[1] + Reset + " tiene " + Yellow + respuestaPlanetas + Reset + " rebeldes.")
						fmt.Println("Vector recibido: " + Yellow + RespuestaVector + Reset)
						fmt.Println("IP de servidor usada: " + Yellow + IPrecibida + Reset)
						fmt.Println("")
						vector[respuesta[1]] = RespuestaVector
						ServerIP[respuesta[1]] = IPrecibida
					} else {
						fmt.Println("La ciudad " + Yellow +  respuesta[2] + Reset + " del planeta " + Yellow + respuesta[1] + Reset + " no se encuentra en el servidor.")
						fmt.Println("IP de servidor usada: " + Yellow + IPrecibida + Reset)
					}
					
					conn.Close()
				}

				
				
			} else {
				if strings.Compare("exit", text) == 0 {
					exit = true
				} else {
					fmt.Println(Red + "ERROR: " + Reset + "Comando erroneo.")
					fmt.Println("Intente con: ")
					fmt.Println(Yellow + "GetNumberRebelds" + Reset)
				}
			}
		}
	
	

}

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

// IP Del Broker
const (
	LOCAL = true
	autor = "Almirante Thrawn"
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
var BrokerAddress = ""

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
	if LOCAL {
		BrokerAddress = "localhost:13370"
	} else {
		BrokerAddress = "137.184.61.128:13370"
	}
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
			fmt.Print(Green + autor + " > "+ Reset)
			text, _ := reader.ReadString('\n')
			if runtime.GOOS == "windows" {
				text = strings.Replace(text, "\r\n", "", -1)
			} else {
				text = strings.Replace(text, "\n", "", -1)
			} 
			
			respuesta := strings.Split(text, " ")
			

			if (respuesta[0] == "AddCity" || respuesta[0] == "UpdateName" || respuesta[0] == "UpdateNumber" || respuesta[0] == "DeleteCity"){
				if (respuesta[0] == "AddCity" && len(respuesta) != 4 && len(respuesta) != 3) {
					fmt.Println(Red + "ERROR: "+ Yellow + "AddCity" + Reset +  " debe tener tener el formato:")
					fmt.Println("AddCity nombre_planeta nombre_ciudad [nuevo_valor]")
					fmt.Println("donde "+ Yellow +"[nuevo_valor]"+ Reset +" es opcional")

				} else if (respuesta[0] == "UpdateName" && len(respuesta) != 4 ) {
					fmt.Println(Red + "ERROR: "+ Yellow + "UpdateName" + Reset +  " debe tener tener el formato:")
					fmt.Println("UpdateName nombre_planeta nombre_ciudad nuevo_valor")

				} else if (respuesta[0] == "UpdateNumber" && len(respuesta) != 4 ) {
					fmt.Println(Red + "ERROR: "+ Yellow + "UpdateNumber" + Reset +  " debe tener tener el formato:")
					fmt.Println("AddCity nombre_planeta nombre_ciudad [nuevo_valor]")

				}else if (respuesta[0] == "DeleteCity" && len(respuesta) != 3 ) {
					fmt.Println(Red + "ERROR: "+ Yellow + "DeleteCity" + Reset +  " debe tener tener el formato:")
					fmt.Println("DeleteCity nombre_planeta nombre_ciudad")

				}else {
					// Conexión a Broker
					conn, err := grpc.Dial(BrokerAddress, grpc.WithInsecure(), grpc.WithBlock())
					if err != nil {
						fmt.Println(Red+"ERROR: No se pudo conectar al broker: %v"+Reset, err)
						log.Fatalf("")
					}
					c := pb.NewManejoComunicacionClient(conn)

					ctx, cancel := context.WithTimeout(context.Background(), time.Second)
					defer cancel()

					r, _ := c.Comunicar(ctx, &pb.MessageRequest{Request: text, Autor: autor, Ip : ServerIP[respuesta[1]],Reloj: vector[respuesta[1]]})
					log.Printf("")
					fmt.Println("IP recibida del Broker: "+Yellow + r.GetReply() + Reset)
					fmt.Println("Contactando al servidor " + Yellow + r.GetReply() + Reset)
					fmt.Println("    con el comando "+Yellow +text + Reset)

					ServerIP[respuesta[1]] = r.GetReply()
					conn.Close()
					

					// Conexión al Servidor Fulcrum
					conn, err = grpc.Dial(ServerIP[respuesta[1]], grpc.WithInsecure(), grpc.WithBlock())
					if err != nil {
						fmt.Println(Red+"ERROR: No se pudo conectar al servidor: %v"+Reset, err)
						log.Fatalf("")
					}
					c = pb.NewManejoComunicacionClient(conn)

					ctx, cancel = context.WithTimeout(context.Background(), time.Second)
					r, _ = c.Comunicar(ctx, &pb.MessageRequest{Request: text, Autor: autor})
					fmt.Println("Vector recibido del Servidor: "+ Yellow +  r.GetReply() + Reset)
					fmt.Println("Guardando vector " + Yellow + r.GetReply() + Reset + " en el planeta " + Yellow + respuesta[1] + Reset)
					fmt.Println("")
					vector[respuesta[1]] = r.GetReply()
					defer cancel()
				}
			} else {
				if strings.Compare("exit", text) == 0 {
					exit = true
				} else {
					fmt.Println(Red + "ERROR: " + Reset + "Comando erroneo.")
					fmt.Println("Intente con: ")
					fmt.Println(Yellow + "AddCity" + Reset + ", " + Yellow + "UpdateName" + Reset + ", " + Yellow + "UpdateNumber" + Reset + " o " + Yellow + "DeleteCity" +Reset)
				}
			}
		}
	
	

}

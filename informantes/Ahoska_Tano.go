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
var vector = map[string]string{}
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
			fmt.Print("Ahoska Tano > ")
			text, _ := reader.ReadString('\n')
			text = strings.Replace(text, "\r\n", "", -1)
			respuesta := strings.Split(text, " ")
			if (respuesta[0] == "AddCity" || respuesta[0] == "UpdateName" || respuesta[0] == "UpdateNumber" || respuesta[0] == "DeleteCity"){
				// Conexión a Broker
				conn, err := grpc.Dial(BrokerAddress, grpc.WithInsecure(), grpc.WithBlock())
				if err != nil {
					log.Fatalf("did not connect: %v", err)
				}
				c := pb.NewManejoComunicacionClient(conn)

				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()
				r, err := c.Comunicar(ctx, &pb.MessageRequest{Request: text, Autor: "Ahoska Tano"})
				log.Printf(`Mensaje recibido del Broker: %s`, r.GetReply())
				log.Printf("Contactando al servidor %s con el comando '%s'", r.GetReply(), text)
				ServerIP = r.GetReply()
				conn.Close()
				

				// Conexión al Servidor Fulcrum
				conn, err = grpc.Dial(ServerIP, grpc.WithInsecure(), grpc.WithBlock())
				if err != nil {
					log.Fatalf("did not connect: %v", err)
				}
				c = pb.NewManejoComunicacionClient(conn)

				ctx, cancel = context.WithTimeout(context.Background(), time.Second)
				r, err = c.Comunicar(ctx, &pb.MessageRequest{Request: text, Autor: "Ahoska Tano"})
				log.Printf(`Mensaje recibido del Servidor: %s`, r.GetReply())
				log.Printf("Guardando vector %s asociado al planeta %s", r.GetReply(), respuesta[1])
				vector[respuesta[1]] = r.GetReply()
				defer cancel()

			} else {
				if strings.Compare("exit", text) == 0 {
					exit = true
				} else {
					fmt.Println("Comando erroneo, intente con AddCity, UpdateName, UpdateNumber o DeleteCity")
				}
			}
		}
	
	

}

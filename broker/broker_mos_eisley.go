package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"runtime"
	"strconv"
	"strings"
	"time"

	pb "github.com/felipe-aguirre/tarea_sist_distribuido/protos"
	"google.golang.org/grpc"
)

const (
	port = ":13370"
	PLAN = "B"
	CANTIDAD_SERVIDORES = 3
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

var servidores = [...]string{"localhost:13371","localhost:13372","localhost:13373"}

var person = map[string]string{
	"Ahoska Tano": "None",
	"Almirante Thrawn": "None",
	"Leia": "None",
}

var seed = rand.NewSource(44)
var random = rand.New(seed)

type ManejoComunicacionServer struct {
	pb.UnimplementedManejoComunicacionServer
}
// Funcion ReceiveMessage debe tener el mismo nombre en informantes
func (s *ManejoComunicacionServer) Comunicar(ctx context.Context, in *pb.MessageRequest) (*pb.MessageReply, error) {
	log.Printf("")
	fmt.Println("Consulta recibida: ")
	fmt.Println(Yellow + in.GetRequest() + Reset)
	fmt.Println("Autor: "+ Yellow + in.GetAutor() + Reset)

	var IP = ""
	// Plan A
	// Recordar la IP de cada informante o Leia y redirigirlo al mismo, hasta los 2 minutos
	// Una conexión en limpio garantiza una asignación aleatoria
	if in.GetAutor()!= "FulcrumDELETE" {
		IP = person[in.GetAutor()]
		if PLAN=="A" {
			if IP == "None" {
				fmt.Println("No hay IP guardada, se elige una aleatoria")
				aleatorio:= random.Intn(CANTIDAD_SERVIDORES)
				IP = servidores[aleatorio]
				person[in.GetAutor()] = IP
				fmt.Println("Servidor: " + Yellow + person[in.GetAutor()] + Reset)
				fmt.Println("    al usuario " + Yellow +in.GetAutor() + Reset + "\n")
			} else {
				fmt.Println("IP en memoria")
				fmt.Println("Servidor: " + Yellow + person[in.GetAutor()] + Reset + "\n")
			}
		} else {
			// Plan B
			// Coordinar con los servidores fulcrum quien tiene un reloj de vector igual o más reciente
			// que el que envía el informante/leia
			
			respuesta := strings.Split(in.GetRequest(), " ")

			if in.GetIp() == "None" || in.GetIp() == "" {
				// Primera vez que se conecta el usuario y no tiene IP asignada
				// Asignar una aleatoria
				fmt.Println("No hay IP guardada, se elige una aleatoria")
				aleatorio:= random.Intn(CANTIDAD_SERVIDORES)
				IP = servidores[aleatorio]
				fmt.Println("Servidor: " + Yellow + IP + Reset)
				fmt.Println("    al usuario " + Yellow +in.GetAutor() + Reset + "\n")

			} else{
				// Ya hay IP Guardada - Comparar los relojes
				// Reconocer que servidor es el actual para identificar el indice
				fmt.Println("IP Guardada: " + Yellow + in.GetIp() + Reset)
				
				indice := 0
				if in.GetIp() == "localhost:13372"{
					indice = 1
				} else if in.GetIp() == "localhost:13373" {
					indice = 2
				}
				fmt.Println("Indice del vector correspondiente: " + Yellow + strconv.Itoa(indice) + Reset)
				// Extraer el número del servidor en el reloj de vectores recibido del informante
				VectorRecibidoInformante := in.GetReloj()
				fmt.Println("Vector recibido por consulta: " + Yellow + in.GetReloj() + Reset)
				vectorRecibidoInformanteLista := strings.Split(VectorRecibidoInformante, " ")
				vectorRecibidoInformanteIP, _ := strconv.Atoi(vectorRecibidoInformanteLista[indice])

				// Consultar servidor por servidor y comparar el número con el recibido
				for i := 0; i < CANTIDAD_SERVIDORES; i++ {
					conn, err := grpc.Dial(servidores[i], grpc.WithInsecure(), grpc.WithBlock())
					if err != nil {
						log.Fatalf("did not connect: %v", err)
					}
					c := pb.NewManejoComunicacionClient(conn)

					ctx, cancel := context.WithTimeout(context.Background(), time.Second)
					// Envía el nombre del planeta
					r, _ := c.ConsultarReloj(ctx, &pb.RelojRequest{Request: respuesta[1]})
					fmt.Println("Vector recibido del servidor " + Yellow + r.GetReply() + Reset)
					VectorRecibidoServidor := r.GetReply()
					vectorRecibidoServidorLista := strings.Split(VectorRecibidoServidor, " ")
					vectorRecibidoServidorIP, _ := strconv.Atoi(vectorRecibidoServidorLista[indice])
					cancel()
					if vectorRecibidoInformanteIP >= vectorRecibidoServidorIP {
						//El servidor si es válido, retornar la IP correspondiente
						fmt.Println("El vector es mayor")
						IP = servidores[indice]	
						fmt.Println("Asignando IP " + Yellow + IP + Reset)
						break
					}
				}
			}
	
		}
	}
	
	if in.GetAutor() == "FulcrumDELETE" {
		person = map[string]string{
			"Ahoska Tano": "None",
			"Almirante Thrawn": "None",
			"Leia": "None",
		}
		IP = "IPs Eliminadas"
		fmt.Println("IPs Reiniciadas")
		fmt.Println("")
	}
	if in.GetAutor() == "Leia"{
		// Conexión al Servidor Fulcrum
		conn, err := grpc.Dial(IP, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		c := pb.NewManejoComunicacionClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		r, _ := c.Comunicar(ctx, &pb.MessageRequest{Request: in.GetRequest(), Autor: "Broker"})
		IP = r.GetReply()+","+IP
		fmt.Println("Enviando data " + Yellow + IP + Reset +  " a Leia")
		defer cancel()
	}
	fmt.Println("")
	return &pb.MessageReply{Reply: IP}, nil
}

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
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterManejoComunicacionServer(s, &ManejoComunicacionServer{})
	fmt.Println("Broker escuchando en puerto " + Yellow + port + Reset + "\n")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

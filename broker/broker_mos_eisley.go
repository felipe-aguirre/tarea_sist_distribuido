package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"time"

	pb "github.com/felipe-aguirre/tarea_sist_distribuido/protos"
	"google.golang.org/grpc"
)

const (
	port = ":50050"
	
)
var servidores = [...]string{"localhost:50051","localhost:50052","localhost:50053"}

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
	log.Printf("Se recibió: %v de parte de %v" , in.GetRequest(), in.GetAutor())

	// Plan A
	// Recordar la IP de cada informante o Leia y redirigirlo al mismo, hasta los 2 minutos
	// Una conexión en limpio garantiza una asignación aleatoria

	// Desactivar Plan A: Comentar desde Acá
	// /*
	var IP = ""
	if in.GetAutor()!= "Fulcrum" {
		IP = person[in.GetAutor()]
	if IP == "None" {
		log.Printf("No hay IP guardada, se elige una aleatoria")
		// Usar rand int (3) para iterar entre [0,1,2] los 3 servidores fulcrum
		aleatorio:= random.Intn(2)
		//Usar aleatorio = 0 mientras para probar con el servidor 1
	
		IP = servidores[aleatorio]
		person[in.GetAutor()] = IP
		log.Printf("Se le ha asignado la IP %v al usuario %v - Retornando valor de IP", person[in.GetAutor()], in.GetAutor())
	} else {
		log.Printf("%v ya tiene la ip %v asignada. - Retornando valor de IP", in.GetAutor(), person[in.GetAutor()])
	}
	// */ Hasta acá


	// Plan B
	// Coordinar con los servidores fulcrum quien tiene un reloj de vector igual o más reciente
	// que el que envía el informante/leia

	}
	
	if in.GetAutor() == "Fulcrum" {
		person = map[string]string{
			"Ahoska Tano": "None",
			"Almirante Thrawn": "None",
			"Leia": "None",
		}
		IP = "IPs Eliminadas"
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
		log.Printf(`Mensaje recibido del Servidor: %s`, r.GetReply())
		IP = r.GetReply()
		defer cancel()
	}
	return &pb.MessageReply{Reply: IP}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterManejoComunicacionServer(s, &ManejoComunicacionServer{})
	log.Printf("Sevidor escuchando en puerto %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

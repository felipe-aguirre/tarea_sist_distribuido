package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	pb "github.com/felipe-aguirre/tarea_sist_distribuido/protos"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)
var vector = []int{0,0,0}

type ManejoComunicacionServer struct {
	pb.UnimplementedManejoComunicacionServer
}
// Funcion ReceiveMessage debe tener el mismo nombre en informantes
func (s *ManejoComunicacionServer) Comunicar(ctx context.Context, in *pb.MessageRequest) (*pb.MessageReply, error) {
	log.Printf("Se recibió: %v", in.GetRequest())

	respuesta := strings.Split(in.GetRequest(), " ")
	if strings.Compare("AddCity", respuesta[0]) == 0 {

		// Configuracion de linea a escribir
		nombre_planeta := respuesta[1]
		nombre_ciudad := respuesta[2]
		linea_a_escribir := nombre_planeta + " " + nombre_ciudad
		registro_log:= in.GetRequest()
		if len(respuesta) > 3 {
			// Se quiere agregar el nuevo valor
			linea_a_escribir = linea_a_escribir + " " + respuesta[3]

		} else {
			linea_a_escribir = linea_a_escribir +" 0"
			registro_log = registro_log + " 0"
		}

		// Open File
		f, err := os.Open("planeta_" + nombre_planeta + ".txt")
		// En caso de que no exista, se crea y se agrega la linea necesaria
		if err != nil {
			log.Println("No se encontró archivo de planeta, creando uno nuevo")
			f.Close()
			f, err := os.OpenFile("planeta_" + nombre_planeta + ".txt", os.O_CREATE|os.O_WRONLY, 0660)
			if err != nil {
				log.Fatalf("No se pudo crear el archivo nuevo: %v", err)
			}
			_, err = f.Write([]byte(linea_a_escribir + "\n"))
			log.Println("Se agrego la linea: ", linea_a_escribir)
			if err != nil {
				log.Fatal(err)
			}
			f.Close()
			logger, _ := os.OpenFile("planeta_"+nombre_planeta+".log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0660)
			logger.Write([]byte(registro_log + "\n"))
			logger.Close()
		} else {
			// El archivo si existe, se verifica primero si ya existe la ciudad
			scanner := bufio.NewScanner(f)
			NoExisteLinea := true
			for scanner.Scan() {
				linea_leida := strings.Split(scanner.Text(), " ")
				// Si la linea existe
				if strings.Compare(nombre_ciudad, linea_leida[1]) == 0 {
					log.Println("Ya existe la ciudad que se intenta agregar")
					NoExisteLinea = false
					break
				}
			}
			f.Close()
			if NoExisteLinea {
				f, _ := os.OpenFile("planeta_"+nombre_planeta+".txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0660)
				_, err = f.Write([]byte(linea_a_escribir + "\n"))
				log.Println("Se agrego la linea: ", linea_a_escribir)
				if err != nil {
					log.Fatal(err)
				}
				f.Close()
				logger, _ := os.OpenFile("planeta_"+nombre_planeta+".log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0660)
				logger.Write([]byte(registro_log + "\n"))
				logger.Close()
			}
		}

	} else if strings.Compare("UpdateName", respuesta[0]) == 0 {
		//TODO: Cuando se intenta cambiar a un nombre que ya existe
		log.Println("Se ejecuto UpdateName")
		nombre_planeta := respuesta[1]
		nombre_ciudad := respuesta[2]
		nuevo_valor := respuesta[3]
		f, _ := os.Open("planeta_" + nombre_planeta + ".txt")
		// Leer el archivo y pasarlo a array
		scanner := bufio.NewScanner(f)
		var textoCompleto []string
		for scanner.Scan() {
			linea_a_escribir := (scanner.Text())
			linea_leida := strings.Split(scanner.Text(), " ")
			// Si la linea existe
			if strings.Compare(nombre_ciudad, linea_leida[1]) == 0 {
				linea_a_escribir = nombre_planeta + " " + nuevo_valor + " " + linea_leida[2]
			}
			textoCompleto = append(textoCompleto, linea_a_escribir)
		}
		f.Close()
		file, _ := os.OpenFile("planeta_"+nombre_planeta+".txt", os.O_CREATE|os.O_WRONLY, 0660)
		for _, elem := range textoCompleto {
			file.Write([]byte(elem + "\n"))
		}
		file.Close()
		logger, _ := os.OpenFile("planeta_"+nombre_planeta+".log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0660)
		logger.Write([]byte(in.GetRequest() + "\n"))
		logger.Close()

	} else if strings.Compare("UpdateNumber", respuesta[0]) == 0 {
		log.Println("Se ejecuto UpdateNumber")
		nombre_planeta := respuesta[1]
		nombre_ciudad := respuesta[2]
		nuevo_valor := respuesta[3]
		f, _ := os.Open("planeta_" + nombre_planeta + ".txt")
		// Leer el archivo y pasarlo a array
		scanner := bufio.NewScanner(f)
		var textoCompleto []string
		for scanner.Scan() {
			linea_a_escribir := (scanner.Text())
			linea_leida := strings.Split(scanner.Text(), " ")
			// Si la linea existe
			if strings.Compare(nombre_ciudad, linea_leida[1]) == 0 {
				linea_a_escribir = nombre_planeta + " " + nombre_ciudad + " " + nuevo_valor
			}
			textoCompleto = append(textoCompleto, linea_a_escribir)
		}
		f.Close()
		file, _ := os.OpenFile("planeta_"+nombre_planeta+".txt", os.O_CREATE|os.O_WRONLY, 0660)
		for _, elem := range textoCompleto {
			file.Write([]byte(elem + "\n"))
		}
		file.Close()
		logger, _ := os.OpenFile("planeta_"+nombre_planeta+".log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0660)
		logger.Write([]byte(in.GetRequest() + "\n"))
		logger.Close()

	} else if strings.Compare("DeleteCity", respuesta[0]) == 0 {
		log.Println("Se ejecuto DeleteCity")
		nombre_planeta := respuesta[1]
		nombre_ciudad := respuesta[2]
		f, err := os.Open("planeta_" + nombre_planeta + ".txt")
		if err != nil {
			log.Fatalf("Linea 135 - Hubo un error al abrir el archivo: %v", err)
		}
		// Leer el archivo y pasarlo a array
		scanner := bufio.NewScanner(f)
		var textoCompleto []string
		for scanner.Scan() {
			linea_a_escribir := (scanner.Text())
			linea_leida := strings.Split(scanner.Text(), " ")
			// Si la linea existe
			if strings.Compare(nombre_ciudad, linea_leida[1]) != 0 {
				textoCompleto = append(textoCompleto, linea_a_escribir)
			}
		}
		f.Close()
		os.Remove("planeta_" + nombre_planeta + ".txt")
		file, _ := os.OpenFile("planeta_"+nombre_planeta+".txt", os.O_CREATE|os.O_WRONLY, 0660)
		for _, elem := range textoCompleto {
			file.Write([]byte(elem + "\n"))
		}
		file.Close()
		logger, _ := os.OpenFile("planeta_"+nombre_planeta+".log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0660)
		logger.Write([]byte(in.GetRequest() + "\n"))
		logger.Close()
	} 
	


	vector[0]++ 
	vectorReturn:="["+strings.Trim(strings.Join(strings.Fields(fmt.Sprint(vector)), " "), "[]")+"]"

	return &pb.MessageReply{Reply: vectorReturn}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterManejoComunicacionServer(s, &ManejoComunicacionServer{})
	log.Printf("Servidor escuchando en puerto %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

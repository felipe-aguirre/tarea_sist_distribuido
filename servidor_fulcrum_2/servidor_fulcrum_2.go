package main

import (
	"bufio"
	"context"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	pb "github.com/felipe-aguirre/tarea_sist_distribuido/protos"
	"google.golang.org/grpc"
)

const (
	port = ":50052"
	indiceServidor = 1
)

var vector = map[string]string{}

var listaPlanetas = []string{}
type ManejoComunicacionServer struct {
	pb.UnimplementedManejoComunicacionServer
}


func (s *ManejoComunicacionServer) Coordinar(ctx context.Context, in *pb.CoordinacionRequest) (*pb.CoordinacionReply, error) {
		// Planetas del servidor X
		// Ejemplo: "planeta1,planeta2,planetae"

		// Logs de cada planeta concatenado con ; del servidor X
		// Ejemplo: "logp11,logp12,logp13;logp21,logp22"

		// Relojes de vector de cada planeta
		// Ejemplo: "RelojPlaneta1, RelojPlaneta2,RelojPlaneta3"
		// Cada RelojPlaneta tiene el formato a b c

  // Paso 1: Recolectar lista planetas
	listaPlanetasSTR := strings.Join(listaPlanetas, ";")

	// Paso 2: Obtener logs de cada planeta
	listaLogs := []string{}
	for _, planeta := range listaPlanetas {
		logsDelPlaneta := []string{}
		fileCheck, _ := os.Open("planeta_" + planeta + ".log")
		// Leer el archivo y pasarlo a array
		scannerCheck := bufio.NewScanner(fileCheck)
		for scannerCheck.Scan() {
			logsDelPlaneta = append(logsDelPlaneta,scannerCheck.Text())
		}
		logsDelPlanetaSTR := strings.Join(logsDelPlaneta, ",")
		listaLogs = append(listaLogs, logsDelPlanetaSTR)
	}
	listaLogsSTR := strings.Join(listaLogs, ";")
	
	//Paso 3: Obtener los vectores de cada planeta
	listaVectores := []string{}
	for _, planeta := range listaPlanetas {
		listaVectores = append(listaVectores, vector[planeta])
	}
	listaVectoresSTR := strings.Join(listaVectores, ";")
	return &pb.CoordinacionReply{Planetas: listaPlanetasSTR, Logs: listaLogsSTR, Vector: listaVectoresSTR}, nil
}
// Funcion ReceiveMessage debe tener el mismo nombre en informantes
func (s *ManejoComunicacionServer) Comunicar(ctx context.Context, in *pb.MessageRequest) (*pb.MessageReply, error) {
	log.Printf("Se recibió: %v", in.GetRequest())
	respuesta := strings.Split(in.GetRequest(), " ")


	//Caso solicitud de Leia (Desde el broker)
	if in.GetAutor() == "Broker" {
		// Leer planeta y ciudad
		log.Println("Se ejecuto GetNumberRebelds")
		nombre_planeta := respuesta[1]
		nombre_ciudad := respuesta[2]
		cantidad_ciudad := "0"
		f, _ := os.Open("planeta_" + nombre_planeta + ".txt")
		// Leer el archivo y pasarlo a array
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			linea_leida := strings.Split(scanner.Text(), " ")
			// Si la linea existe
			if strings.Compare(nombre_ciudad, linea_leida[1]) == 0 {
				cantidad_ciudad = linea_leida[2]
			}
		}
		f.Close()
		posibleRespuesta := ""
 		if _, ok := vector[nombre_planeta]; ok {
			posibleRespuesta = vector[nombre_planeta]+ "," + cantidad_ciudad
		}else{ 
			posibleRespuesta = "0 0 0," + cantidad_ciudad
		}
		
		log.Println("Respuesta: ",posibleRespuesta)
		return &pb.MessageReply{Reply: posibleRespuesta}, nil



	}

	if strings.Compare("AddCity", respuesta[0]) == 0 {
		log.Println("Se ejecuto AddCity")
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
			log.Printf("No se encontró archivo de planeta, creando uno nuevo")
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
			
			//Add de planeta a listado
			listaPlanetas = append(listaPlanetas, nombre_planeta)

			//Add de vector
			if _, ok := vector[nombre_planeta]; ok {
				// Vector si existe en el planeta
				vectorRespuesta := strings.Split(vector[nombre_planeta], " ")
		
				x, _ := strconv.Atoi(vectorRespuesta[indiceServidor])
				vectorRespuesta[indiceServidor] = strconv.Itoa(x + 1)
				vector[nombre_planeta] = strings.Join(vectorRespuesta, " ")
			} else {
				// vector no existe, se crea nuevo con valor 1 0 0 | 0 1 0 | 0 0 1
				if indiceServidor == 0 {
					vector[nombre_planeta] = "1 0 0"
				}else if indiceServidor == 1{
					vector[nombre_planeta] = "0 1 0"
				}else {
					vector[nombre_planeta] = "0 0 1"
				}
			}
	
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

				//Add de vector
				if _, ok := vector[nombre_planeta]; ok {
					// Vector si existe en el planeta
					vectorRespuesta := strings.Split(vector[nombre_planeta], " ")
			
					x, _ := strconv.Atoi(vectorRespuesta[indiceServidor])
					vectorRespuesta[indiceServidor] = strconv.Itoa(x + 1)
					vector[nombre_planeta] = strings.Join(vectorRespuesta, " ")
				} else {
					// vector no existe, se crea nuevo con valor 1 0 0 | 0 1 0 | 0 0 1
					if indiceServidor == 0 {
						vector[nombre_planeta] = "1 0 0"
					}else if indiceServidor == 1{
						vector[nombre_planeta] = "0 1 0"
					}else {
						vector[nombre_planeta] = "0 0 1"
					}
				}
			}
		}

	} else if strings.Compare("UpdateName", respuesta[0]) == 0 {
		//TODO: Cuando se intenta cambiar a un nombre que ya existe
		log.Println("Se ejecuto UpdateName")
		nombre_planeta := respuesta[1]
		nombre_ciudad := respuesta[2]
		nuevo_valor := respuesta[3]

		//Revisión si valor nuevo es una ciudad existente. De ser así se cancela
		//la operación
		ciudadExiste := false
		fileCheck, _ := os.Open("planeta_" + nombre_planeta + ".txt")
		// Leer el archivo y pasarlo a array
		scannerCheck := bufio.NewScanner(fileCheck)
		for scannerCheck.Scan() {
			linea_leida := strings.Split(scannerCheck.Text(), " ")
			// Si la linea existe
			if strings.Compare(nuevo_valor, linea_leida[1]) == 0 {
				ciudadExiste = true
			}
		}
		fileCheck.Close()
		
		if (!ciudadExiste) {
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

			//Add de vector
			if _, ok := vector[nombre_planeta]; ok {
				// Vector si existe en el planeta
				vectorRespuesta := strings.Split(vector[nombre_planeta], " ")
		
				x, _ := strconv.Atoi(vectorRespuesta[indiceServidor])
				vectorRespuesta[indiceServidor] = strconv.Itoa(x + 1)
				vector[nombre_planeta] = strings.Join(vectorRespuesta, " ")
			} else {
				// vector no existe, se crea nuevo con valor 1 0 0 | 0 1 0 | 0 0 1
				if indiceServidor == 0 {
					vector[nombre_planeta] = "1 0 0"
				}else if indiceServidor == 1{
					vector[nombre_planeta] = "0 1 0"
				}else {
					vector[nombre_planeta] = "0 0 1"
				}
			}

		}
		

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

		//Add de vector
		if _, ok := vector[nombre_planeta]; ok {
			// Vector si existe en el planeta
			vectorRespuesta := strings.Split(vector[nombre_planeta], " ")
			
			x, _ := strconv.Atoi(vectorRespuesta[indiceServidor])
			vectorRespuesta[indiceServidor] = strconv.Itoa(x + 1)
			vector[nombre_planeta] = strings.Join(vectorRespuesta, " ")
		} else {
			// vector no existe, se crea nuevo con valor 1 0 0 | 0 1 0 | 0 0 1
			if indiceServidor == 0 {
				vector[nombre_planeta] = "1 0 0"
			}else if indiceServidor == 1{
				vector[nombre_planeta] = "0 1 0"
			}else {
				vector[nombre_planeta] = "0 0 1"
			}
		}

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
		if len(textoCompleto) != 0{
			file, _ := os.OpenFile("planeta_"+nombre_planeta+".txt", os.O_CREATE|os.O_WRONLY, 0660)
			for _, elem := range textoCompleto {
				file.Write([]byte(elem + "\n"))
			}
			file.Close()
			logger, _ := os.OpenFile("planeta_"+nombre_planeta+".log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0660)
			logger.Write([]byte(in.GetRequest() + "\n"))
			logger.Close()
		}

		//Add de vector
		if _, ok := vector[nombre_planeta]; ok {
			// Vector si existe en el planeta
			vectorRespuesta := strings.Split(vector[nombre_planeta], " ")
	
			x, _ := strconv.Atoi(vectorRespuesta[indiceServidor])
			vectorRespuesta[indiceServidor] = strconv.Itoa(x + 1)
			vector[nombre_planeta] = strings.Join(vectorRespuesta, " ")
		} else {
			// vector no existe, se crea nuevo con valor 1 0 0 | 0 1 0 | 0 0 1
			if indiceServidor == 0 {
				vector[nombre_planeta] = "1 0 0"
			}else if indiceServidor == 1{
				vector[nombre_planeta] = "0 1 0"
			}else {
				vector[nombre_planeta] = "0 0 1"
			}
		}
	} 
	
	vectorViejo := strings.Split(vector[respuesta[1]], " ")
	x, _ := strconv.Atoi(vectorViejo[indiceServidor])
	vectorViejo[indiceServidor] = strconv.Itoa(x - 1)
	vectorViejoStr := strings.Join(vectorViejo, " ")
	log.Printf("Vector nuevo: [%v] > [%v]",vectorViejoStr, vector[respuesta[1]])

	return &pb.MessageReply{Reply: vector[respuesta[1]]}, nil
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
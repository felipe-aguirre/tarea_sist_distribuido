package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	pb "github.com/felipe-aguirre/tarea_sist_distribuido/protos"
	"google.golang.org/grpc"
)

const (
	BrokerAddress = "localhost:50050"
	port = ":50051"
	indiceServidor = 0
	TIMER = 10
)

var vector = map[string]string{}
var listaPlanetas = []string{}
type ManejoComunicacionServer struct {
	pb.UnimplementedManejoComunicacionServer
}
func max(a, b int) int {
	if a > b {
			return a
	}
	return b
}
func deleter() {
	dirname := "." + string(filepath.Separator)
      d, err := os.Open(dirname)
      if err != nil {
          fmt.Println(err)
          os.Exit(1)
      }
      defer d.Close()

      files, err := d.Readdir(-1)
      if err != nil {
          fmt.Println(err)
          os.Exit(1)
      }
			
      for _, file := range files {

          if file.Mode().IsRegular() {
              if filepath.Ext(file.Name()) == ".txt" {
                os.Remove(file.Name())
              }
							if filepath.Ext(file.Name()) == ".log" {
                os.Remove(file.Name())
              }
          }
      }
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
	x, _ := strconv.Atoi(vectorViejo[0])
	vectorViejo[0] = strconv.Itoa(x - 1)
	vectorViejoStr := strings.Join(vectorViejo, " ")
	log.Printf("Vector nuevo: [%v] > [%v]",vectorViejoStr, vector[respuesta[1]])

	return &pb.MessageReply{Reply: vector[respuesta[1]]}, nil
}

func aplicarCambiosRecibidos(planetas string, logs string, vectores string) {
	log.Printf("Iniciando aplicación de cambios recibidos")
	listaPlanetasRecibida := strings.Split(planetas, ";")
	if len(listaPlanetasRecibida) == 1 {
		if strings.Compare(listaPlanetasRecibida[0], "") == 0{
			listaPlanetasRecibida = []string{}
		}
	}
	listaLogs := strings.Split(logs, ";")
	if len(listaLogs) == 1 {
		if strings.Compare(listaLogs[0], "") == 0{
			listaLogs = []string{}
		}
	}
	listaVectores := strings.Split(vectores, ";")
	if len(listaVectores) == 1 {
		if strings.Compare(listaVectores[0], "") == 0{
			listaVectores = []string{}
		}
	}
	


	for index, planeta := range listaPlanetasRecibida {
		log.Printf("Edición de planeta: %v",planeta)
		// Paso 1: Intentar abrir el archivo, si no existe crear uno nuevo
		f, err := os.Open("planeta_" + planeta + ".txt")
		// En caso de que no exista, se crea y se agrega la linea necesaria
		if err != nil {
			log.Printf("El planeta %v no se encuentra en el servidor Fulcrum 1, creando nuevo", planeta)
			f.Close()
			f, _ := os.OpenFile("planeta_" + planeta + ".txt", os.O_CREATE|os.O_WRONLY, 0660)
			f.Close()
			listaPlanetas = append(listaPlanetas, planeta)		
			
		} 
		// Archivo existe - continua a aplicación de log
		
		// Paso 2: Aplicar el log correspondiente al archivo
		if len(listaLogs) != 0 {
			listaLogsPlaneta := strings.Split(listaLogs[index], ",")
		for _, elem := range listaLogsPlaneta {
			log.Printf("Linea de log: %v",elem)
			logActual:= strings.Split(elem, ",")
			
			for _, comando := range logActual {
				log.Printf("Log actual: %v",comando)
				log.Printf("Largo del log: actual: %v\n",len(comando))
				respuesta := strings.Split(comando, " ")
				log.Printf("Largo del log: actual: %v\n",len(respuesta))
				nombre_planeta := respuesta[1]
				nombre_ciudad := respuesta[2]
				nuevo_valor:= ""
				if len(respuesta) == 4 {
					nuevo_valor= respuesta[3]
				}
				

				//Caso AddCity
				if strings.Compare("AddCity", respuesta[0]) == 0 {
					f, _ := os.Open("planeta_" + nombre_planeta + ".txt")
					scanner := bufio.NewScanner(f)
					var textoCompleto []string
					for scanner.Scan() {
						linea_a_escribir := (scanner.Text())
						linea_leida := strings.Split(scanner.Text(), " ")
						// Si la linea actual no corresponde a la ciudad actual, se
						// agrega al buffer textoCompleto
						if strings.Compare(nombre_ciudad, linea_leida[1]) != 0 {
							textoCompleto = append(textoCompleto, linea_a_escribir)
						}
					}
					linea_a_escribir := nombre_planeta + " " + nombre_ciudad + " " + nuevo_valor
					textoCompleto = append(textoCompleto, linea_a_escribir)
					f.Close()
					file, _ := os.OpenFile("planeta_"+nombre_planeta+".txt", os.O_CREATE|os.O_WRONLY, 0660)
					for _, elem := range textoCompleto {
						file.Write([]byte(elem + "\n"))
					}
					file.Close()


				
				// Caso UpdateName
				} else if strings.Compare("UpdateName", respuesta[0]) == 0 {
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
					}
				} else if strings.Compare("UpdateNumber", respuesta[0]) == 0 {
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
				} else if strings.Compare("DeleteCity", respuesta[0]) == 0 {
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
					}
					
				}
			}
		//Paso 3: Coordinar vectores dejando los valores más altos
		}
		}

		
		log.Printf("Revisión de vector para planeta: %v",planeta)
		// listaVectores = [1 1 0; 1 0 1; 2 0 1  etc] ya separados como array de strings
		// Comparar vector de fulcrum 1 con el entregado

		// Vector en blanco
		vector0 := 0
		vector1 := 0
		vector2 := 0
		vectorFulcrum1:= "0 0 0"
		// Vector de fulcrum 1:
		if _, ok := vector[planeta]; ok {
			// Vector si existe en el planeta
			vectorFulcrum1 = vector[planeta]
			log.Printf("Vector actual del planeta %v en Fulcrum 1: %v",planeta, vectorFulcrum1)
			vectorFulcrum1List := strings.Split(vectorFulcrum1, " ")
			vector0, _ = strconv.Atoi(vectorFulcrum1List[0])
			vector1, _ = strconv.Atoi(vectorFulcrum1List[1])
			vector2, _ = strconv.Atoi(vectorFulcrum1List[2])
		} else {
			log.Printf("No existe vector del planeta %v en Fulcrum 1", planeta)
		
		}
		
		
		// Vector nuevo
		vectorNuevo := listaVectores[index]
		log.Printf("Vector recibido para el planeta %v: %v",planeta, vectorNuevo)
		vectorNuevoFulcrum1List := strings.Split(vectorNuevo, " ")
		vectorNuevo0, _ := strconv.Atoi(vectorNuevoFulcrum1List[0])
		vectorNuevo1, _ := strconv.Atoi(vectorNuevoFulcrum1List[1])
		vectorNuevo2, _ := strconv.Atoi(vectorNuevoFulcrum1List[2])

		vectorFinal := []string{strconv.Itoa(max(vector0, vectorNuevo0)),
			strconv.Itoa(max(vector1, vectorNuevo1)), 
			strconv.Itoa(max(vector2, vectorNuevo2))}
		
		log.Printf("Vector nuevo: [%v] > [%v]",vectorFulcrum1, strings.Join(vectorFinal, " "))
		vector[planeta] = strings.Join(vectorFinal, " ")
	}
}
func ejecutarCoordinacion() {
	for range time.Tick(time.Second * TIMER) {
		log.Printf("Inicio de Merge")

		// Solicitar archivos, se recibe listado de:

		// Planetas del servidor X
		// Ejemplo: "planeta1,planeta2,planetae"

		// Logs de cada planeta concatenado con ; del servidor X
		// Ejemplo: "logp11,logp12,logp13;logp21,logp22"

		// Relojes de vector de cada planeta
		// Ejemplo: "RelojPlaneta1, RelojPlaneta2,RelojPlaneta3"
		// Cada RelojPlaneta tiene el formato a b c

		// Conexión a Fulcrum 2
		log.Printf("Solicitando información a Fulcrum 2")
		conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		c := pb.NewManejoComunicacionClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		r, _ := c.Coordinar(ctx, &pb.CoordinacionRequest{Request: "Coordinemos"})
		log.Printf("Se recibió respuesta de Fulcrum 2:")
		planetaFulcrum2 :=r.GetPlanetas()
		logFulcrum2 := r.GetLogs()
		vectorFulcrum2 := r.GetVector()
		cancel()
		log.Printf("Planetas: %v",planetaFulcrum2)
		log.Printf("Logs: %v",logFulcrum2)
		log.Printf("Vectores: %v",vectorFulcrum2)
		aplicarCambiosRecibidos(planetaFulcrum2, logFulcrum2, vectorFulcrum2)

		// Conexión a Fulcrum 3
		/*
		log.Printf("Solicitando información a Fulcrum 3")
		conn, err = grpc.Dial("localhost:50053", grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		c = pb.NewManejoComunicacionClient(conn)

		ctx, cancel = context.WithTimeout(context.Background(), time.Second)
		r, _ = c.Coordinar(ctx, &pb.CoordinacionRequest{Request: "Coordinemos"})
		log.Printf("Se recibió respuesta de Fulcrum 3")
		planetaFulcrum3 :=r.GetPlanetas()
		logFulcrum3 := r.GetLogs()
		vectorFulcrum3 := r.GetVector()
		cancel()
		log.Printf("Planetas: %v", planetaFulcrum3)
		log.Printf("Logs: %v", logFulcrum3)
		log.Printf("Vectores: %v", vectorFulcrum3)
		aplicarCambiosRecibidos(planetaFulcrum3, logFulcrum3, vectorFulcrum3)
		*/


		// Borrar los logs del servidor fulcrum 1
		log.Printf("Eliminando logs de Fulcrum 1")
		for _, planeta := range listaPlanetas {
			os.Remove("planeta_" + planeta + ".log")
		}
		
		// Preparando DATA para enviar a fulcrum 2 y 3

		// Paso 1: Recolectar lista planetas
		listaPlanetasSTR := strings.Join(listaPlanetas, ";")

		// Paso 2: Obtener TXT de cada planeta
		listaTxt := []string{}
		for _, planeta := range listaPlanetas {
			logsDelPlaneta := []string{}
			fileCheck, _ := os.Open("planeta_" + planeta + ".txt")
			// Leer el archivo y pasarlo a array
			scannerCheck := bufio.NewScanner(fileCheck)
			for scannerCheck.Scan() {
				logsDelPlaneta = append(logsDelPlaneta,scannerCheck.Text())
			}
			logsDelPlanetaSTR := strings.Join(logsDelPlaneta, ",")
			listaTxt = append(listaTxt, logsDelPlanetaSTR)
		}
		listaTxtSTR := strings.Join(listaTxt, ";")
		
		//Paso 3: Obtener los vectores de cada planeta
		listaVectores := []string{}
		for _, planeta := range listaPlanetas {
			listaVectores = append(listaVectores, vector[planeta])
		}
		listaVectoresSTR := strings.Join(listaVectores, ";")


		// Ordenar a los servidores Fulcrum2 y 3 acatar los nuevos cambios
		//Fulcrum 2
		log.Printf("Enviando nueva data a Fulcrum 2")
		conn, err = grpc.Dial("localhost:50052", grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		c = pb.NewManejoComunicacionClient(conn)

		ctx, cancel = context.WithTimeout(context.Background(), time.Second)
		rReestructuracion, _ := c.Reestructurar(ctx, &pb.ReestructuracionRequest{Planetas: listaPlanetasSTR, Vectores: listaVectoresSTR, Registrotxt: listaTxtSTR})
		log.Printf("Se recibió respuesta de Fulcrum 2: %v", rReestructuracion.GetReply())
		cancel()

		//Fulcrum 3
		/*
		log.Printf("Enviando nueva data a Fulcrum 2")
		conn, err = grpc.Dial("localhost:50053", grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		c = pb.NewManejoComunicacionClient(conn)

		ctx, cancel = context.WithTimeout(context.Background(), time.Second)
		rReestructuracion, _ = c.Reestructurar(ctx, &pb.ReestructuracionRequest{Planetas: listaPlanetasSTR, Vectores: listaVectoresSTR, Registrotxt: listaTxtSTR})
		log.Printf("Se recibió respuesta de Fulcrum 3: %v", rReestructuracion.GetReply())
		cancel()
		*/

		// Avisarle al broker que reinicie las IPs
		conn, err = grpc.Dial(BrokerAddress, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		c = pb.NewManejoComunicacionClient(conn)
		ctx, cancel = context.WithTimeout(context.Background(), time.Second)
		respBroker, _ := c.Comunicar(ctx, &pb.MessageRequest{Request: "Eliminar", Autor: "FulcrumDELETE"})
		log.Printf("Respuesta broker: %v", respBroker.GetReply())
		cancel()

	}
}

func main() {
	deleter()
	go ejecutarCoordinacion()
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

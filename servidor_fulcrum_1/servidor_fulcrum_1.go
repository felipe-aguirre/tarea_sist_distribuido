package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	pb "github.com/felipe-aguirre/tarea_sist_distribuido/protos"
	"google.golang.org/grpc"
)

const (
	port = ":13371"
	indiceServidor = 0
	TIMER = 120
	LOCAL = false
)

var BrokerAddress = ""
var Fulcrum2 = ""
var Fulcrum3 = ""

var Reset  = "\033[0m"
var Red    = "\033[31m"
var Green  = "\033[32m"
var Yellow = "\033[33m"
var Blue   = "\033[34m"
var Purple = "\033[35m"
var Cyan   = "\033[36m"
var Gray   = "\033[37m"
var White  = "\033[97m"


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

func (s *ManejoComunicacionServer) ConsultarReloj(ctx context.Context, in *pb.RelojRequest) (*pb.RelojReply, error) {
	vectorDelPlaneta := ""
	if _, ok := vector[in.GetRequest()]; ok {
		vectorDelPlaneta = vector[in.GetRequest()]
	}else{ 
		vectorDelPlaneta = "0 0 0"
	}
	return &pb.RelojReply{Reply: vectorDelPlaneta}, nil
}


// Funcion ReceiveMessage debe tener el mismo nombre en informantes
func (s *ManejoComunicacionServer) Comunicar(ctx context.Context, in *pb.MessageRequest) (*pb.MessageReply, error) {
	descuento := 1
	log.Printf("")
	fmt.Println("Se recibió: " + Yellow + in.GetRequest() + Reset)
	respuesta := strings.Split(in.GetRequest(), " ")


	//Caso solicitud de Leia (Desde el broker)
	if in.GetAutor() == "Broker" {
		// Leer planeta y ciudad
		fmt.Println("Ejecutando " + Yellow + "GetNumberRebelds" + Reset)
		nombre_planeta := respuesta[1]
		nombre_ciudad := respuesta[2]
		cantidad_ciudad := "-1"
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
		if cantidad_ciudad == "-1"{
			// No se encontró la ciudad
			fmt.Println("La ciudad " + Yellow + nombre_ciudad + Reset + " no se encuentra en este servidor")
		}else {
			fmt.Println("Ciudad: " + Yellow + nombre_ciudad + Reset)
			fmt.Println("Cantidad Rebeldes: " + Yellow + cantidad_ciudad + Reset)
		}
		fmt.Println("")
		return &pb.MessageReply{Reply: posibleRespuesta}, nil



	}

	if strings.Compare("AddCity", respuesta[0]) == 0 {
		fmt.Println("Ejecutando " + Yellow + "AddCity" + Reset)
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
			fmt.Println(Purple + "Warn: " + Reset + "No se encontró archivo de planeta.")
			fmt.Println("    Creando nuevos .txt y .log")

			f.Close()
			f, err := os.OpenFile("planeta_" + nombre_planeta + ".txt", os.O_CREATE|os.O_WRONLY, 0660)
			if err != nil {
				log.Fatalf("No se pudo crear el archivo nuevo: %v", err)
			}
			_, err = f.Write([]byte(linea_a_escribir + "\n"))
			fmt.Println("Agregado a .txt: " +  Yellow + linea_a_escribir + Reset)
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
					fmt.Println(Red + "ERROR: " + Reset + "Ya existe ciudad")
					fmt.Println("Ignorando comando recibido")
					NoExisteLinea = false
					descuento = 0
					break
				}
			}
			f.Close()
			if NoExisteLinea {
				f, _ := os.OpenFile("planeta_"+nombre_planeta+".txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0660)
				_, err = f.Write([]byte(linea_a_escribir + "\n"))
				fmt.Println("Agregado a .txt: " +  Yellow + linea_a_escribir + Reset)
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
		fmt.Println("Ejecutando " + Yellow + "UpdateName" + Reset)

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
				fmt.Println(Red + "ERROR: " + Reset + "Hay otra ciudad con el nombre " + Yellow + nuevo_valor + Reset)
				descuento = 0
			}
		}
		fileCheck.Close()
		
		if (!ciudadExiste) {
			fmt.Println("No hay ciudades con el nombre " + Yellow + nuevo_valor + Reset)
			fmt.Println("Inicia actualización de nombre:")

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
				file.Write([]byte(elem +"\n") )
			}
			file.Close()
			logger, _ := os.OpenFile("planeta_"+nombre_planeta+".log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0660)
			logger.Write([]byte(in.GetRequest() + "\n"))
			logger.Close()
			fmt.Println("Reemplazado: ")
			fmt.Println(Yellow + nombre_ciudad + Reset + " > " + Yellow + nuevo_valor + Reset)
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
		fmt.Println("Ejecutando " + Yellow + "UpdateNumber" + Reset)

		nombre_planeta := respuesta[1]
		nombre_ciudad := respuesta[2]
		nuevo_valor := respuesta[3]
		valor_antiguo := ""
		f, _ := os.Open("planeta_" + nombre_planeta + ".txt")
		// Leer el archivo y pasarlo a array
		scanner := bufio.NewScanner(f)
		var textoCompleto []string
		updateNumberEditado := false
		for scanner.Scan() {
			linea_a_escribir := (scanner.Text())
			linea_leida := strings.Split(scanner.Text(), " ")
			// Si la linea existe
			if strings.Compare(nombre_ciudad, linea_leida[1]) == 0 {
				linea_a_escribir = nombre_planeta + " " + nombre_ciudad + " " + nuevo_valor
				updateNumberEditado = true
				valor_antiguo = linea_leida[2]
			}
			textoCompleto = append(textoCompleto, linea_a_escribir)
		}
		f.Close()
		if updateNumberEditado {
			fmt.Println("Si existe la ciudad " + Yellow + nombre_ciudad + Reset)
			fmt.Println("Inicia actualización de cantidad:")
			file, _ := os.OpenFile("planeta_"+nombre_planeta+".txt", os.O_CREATE|os.O_WRONLY, 0660)
			for _, elem := range textoCompleto {
				file.Write([]byte(elem + "\n"))
			}
			file.Close()
			logger, _ := os.OpenFile("planeta_"+nombre_planeta+".log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0660)
			logger.Write([]byte(in.GetRequest() + "\n"))
			logger.Close()

			fmt.Println("Reemplazado: ")
			fmt.Println(Yellow + valor_antiguo + Reset + " > " + Yellow + nuevo_valor + Reset)
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
			descuento = 0
			fmt.Println(Red + "ERROR: " + Reset + "No hay ciudad con el nombre " + Yellow + nombre_ciudad + Reset)
		}
		

	} else if strings.Compare("DeleteCity", respuesta[0]) == 0 {
		fmt.Println("Ejecutando " + Yellow + "DeleteCity" + Reset)

		nombre_planeta := respuesta[1]
		nombre_ciudad := respuesta[2]
		f, err := os.Open("planeta_" + nombre_planeta + ".txt")
		if err != nil {
			log.Fatalf("Linea 135 - Hubo un error al abrir el archivo: %v", err)
		}
		// Leer el archivo y pasarlo a array
		scanner := bufio.NewScanner(f)
		var textoCompleto []string
		lineaExistia := false
		for scanner.Scan() {
			linea_a_escribir := (scanner.Text())
			linea_leida := strings.Split(scanner.Text(), " ")
			// Si la linea existe
			if strings.Compare(nombre_ciudad, linea_leida[1]) != 0 {
				// Linea no existe, se copia al buffer
				textoCompleto = append(textoCompleto, linea_a_escribir)
			} else {
				// Linea si existe
				lineaExistia = true
			}
		}
		f.Close()
		if lineaExistia {
			fmt.Println("Si existe la ciudad " + Yellow + nombre_ciudad + Reset)
			fmt.Println("Inicia borrado de ciudad.")
			os.Remove("planeta_" + nombre_planeta + ".txt")
			file, _ := os.OpenFile("planeta_"+nombre_planeta+".txt", os.O_CREATE|os.O_WRONLY, 0660)
			if len(textoCompleto) != 0{
				for _, elem := range textoCompleto {
					file.Write([]byte(elem + "\n"))
				}
				file.Close()	
			}
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
		} else {
			// Linea no existe, ignorar cambios
			descuento = 0
			fmt.Println(Red + "ERROR: " + Reset + "No hay ciudad con el nombre " + Yellow + nombre_ciudad + Reset)
		}
		
	} 
	
	vectorViejo := strings.Split(vector[respuesta[1]], " ")
	x, _ := strconv.Atoi(vectorViejo[indiceServidor])
	vectorViejo[indiceServidor] = strconv.Itoa(x - descuento)
	vectorViejoStr := strings.Join(vectorViejo, " ")
	if descuento == 1 {
	fmt.Println("Cambio vector: " + Yellow + vectorViejoStr + Reset + " > " + Yellow + vector[respuesta[1]] + Reset)
	} else{
		fmt.Println("Vector se mantiene en " + Yellow + vectorViejoStr + Reset )
	}
	fmt.Println("")

	return &pb.MessageReply{Reply: vector[respuesta[1]]}, nil
}

func aplicarCambiosRecibidos(planetas string, logs string, vectores string) {
	fmt.Println("Iniciando aplicación de cambios")
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
		fmt.Println("Edición de planeta: " + Yellow + planeta + Reset)
		// Paso 1: Intentar abrir el archivo, si no existe crear uno nuevo
		f, err := os.Open("planeta_" + planeta + ".txt")
		// En caso de que no exista, se crea y se agrega la linea necesaria
		if err != nil {
			fmt.Println("  El planeta " + Yellow + planeta + Reset + " no se encuentra en Fulcrum 1")
			fmt.Println("  Creando archivo .txt nuevo.")
			f.Close()
			f, _ := os.OpenFile("planeta_" + planeta + ".txt", os.O_CREATE|os.O_WRONLY, 0660)
			f.Close()
			listaPlanetas = append(listaPlanetas, planeta)		
			
		} 
		// Archivo existe - continua a aplicación de log
		
		// Paso 2: Aplicar el log correspondiente al archivo
		if listaLogs[index] != "-" {
			log.Printf("  Log del planeta: " + Yellow + listaLogs[index] + Reset)
			listaLogsPlaneta := strings.Split(listaLogs[index], ",")
			for _, elem := range listaLogsPlaneta {
				
				logActual:= strings.Split(elem, ",")
				
				for _, comando := range logActual {
					log.Printf("  Comando Aplicado: " + Yellow + comando + Reset)
					respuesta := strings.Split(comando, " ")
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
						file, _ := os.OpenFile("planeta_"+nombre_planeta+".txt", os.O_CREATE|os.O_WRONLY, 0660)
						if len(textoCompleto) != 0{
							for _, elem := range textoCompleto {
								file.Write([]byte(elem + "\n"))
							}			
						}
						file.Close()
						
					}
				}
			//Paso 3: Coordinar vectores dejando los valores más altos
			}
		} else {
			fmt.Println("  No hay comandos que aplicar.")
		}

		
		fmt.Println("  Revisión de vector: ")
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
			fmt.Println("  Vector actual del planeta en Fulcrum 1: " + Yellow + vectorFulcrum1 + Reset)
			vectorFulcrum1List := strings.Split(vectorFulcrum1, " ")
			vector0, _ = strconv.Atoi(vectorFulcrum1List[0])
			vector1, _ = strconv.Atoi(vectorFulcrum1List[1])
			vector2, _ = strconv.Atoi(vectorFulcrum1List[2])
		} else {
			fmt.Println("  No existe vector del planeta en Fulcrum 1")
		
		}
		
		
		// Vector nuevo
		vectorNuevo := listaVectores[index]
		fmt.Println("  Vector recibido : " + Yellow + vectorNuevo + Reset)
		vectorNuevoFulcrum1List := strings.Split(vectorNuevo, " ")
		vectorNuevo0, _ := strconv.Atoi(vectorNuevoFulcrum1List[0])
		vectorNuevo1, _ := strconv.Atoi(vectorNuevoFulcrum1List[1])
		vectorNuevo2, _ := strconv.Atoi(vectorNuevoFulcrum1List[2])

		vectorFinal := []string{strconv.Itoa(max(vector0, vectorNuevo0)),
			strconv.Itoa(max(vector1, vectorNuevo1)), 
			strconv.Itoa(max(vector2, vectorNuevo2))}
		
		fmt.Println("Vector nuevo: "+ Yellow + vectorFulcrum1 + Reset + " > " + Yellow + strings.Join(vectorFinal, " ") + Reset)
		vector[planeta] = strings.Join(vectorFinal, " ")
		

	}
	fmt.Println("Fin aplicación de cambios")
	fmt.Println("")
}
func ejecutarCoordinacion() {
	for range time.Tick(time.Second * TIMER) {
		fmt.Println(Red + "=== INICIO MERGE ===" + Reset)

		// Solicitar archivos, se recibe listado de:

		// Planetas del servidor X
		// Ejemplo: "planeta1,planeta2,planetae"

		// Logs de cada planeta concatenado con ; del servidor X
		// Ejemplo: "logp11,logp12,logp13;logp21,logp22"

		// Relojes de vector de cada planeta
		// Ejemplo: "RelojPlaneta1, RelojPlaneta2,RelojPlaneta3"
		// Cada RelojPlaneta tiene el formato a b c

		// Conexión a Fulcrum 2
		fmt.Println("Solicitando información a " + Yellow + "Fulcrum 2" + Reset)
		conn, err := grpc.Dial(Fulcrum2, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		c := pb.NewManejoComunicacionClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		r, _ := c.Coordinar(ctx, &pb.CoordinacionRequest{Request: "Coordinemos"})
		fmt.Println("Se recibió respuesta de "+ Yellow + "Fulcrum 2" + Reset)
		planetaFulcrum2 :=r.GetPlanetas()
		logFulcrum2 := r.GetLogs()
		vectorFulcrum2 := r.GetVector()
		cancel()
		fmt.Println("Planetas: " + Yellow +planetaFulcrum2 + Reset)
		fmt.Println("Logs: " + Yellow + logFulcrum2 + Reset)
		fmt.Println("Vectores: " + Yellow + vectorFulcrum2 + Reset)
		aplicarCambiosRecibidos(planetaFulcrum2, logFulcrum2, vectorFulcrum2)
		fmt.Println("")
		
		
		// Conexión a Fulcrum 3
		fmt.Println("Solicitando información a " + Yellow + "Fulcrum 3" + Reset)
		conn, err = grpc.Dial(Fulcrum3, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		c = pb.NewManejoComunicacionClient(conn)
		ctx, cancel = context.WithTimeout(context.Background(), time.Second)
		r, _ = c.Coordinar(ctx, &pb.CoordinacionRequest{Request: "Coordinemos"})
		fmt.Println("Se recibió respuesta de "+ Yellow + "Fulcrum 3" + Reset)
		planetaFulcrum3 :=r.GetPlanetas()
		logFulcrum3 := r.GetLogs()
		vectorFulcrum3 := r.GetVector()
		cancel()
		fmt.Println("Planetas: " + Yellow +planetaFulcrum3 + Reset)
		fmt.Println("Logs: " + Yellow + logFulcrum3 + Reset)
		fmt.Println("Vectores: " + Yellow + vectorFulcrum3 + Reset)
		aplicarCambiosRecibidos(planetaFulcrum3, logFulcrum3, vectorFulcrum3)
		fmt.Println("")


		// Borrar los logs del servidor fulcrum 1
		fmt.Println("Eliminando logs del servidor")
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
				//Cada linea
				logsDelPlaneta = append(logsDelPlaneta,scannerCheck.Text())
			}
			if len(logsDelPlaneta) == 0 {
				logsDelPlaneta = append(logsDelPlaneta,"-")
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
		fmt.Println("Enviando data a " + Yellow + "Fulcrum 2" + Reset)
		conn, err = grpc.Dial(Fulcrum2, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		c = pb.NewManejoComunicacionClient(conn)

		ctx, cancel = context.WithTimeout(context.Background(), time.Second)
		rReestructuracion, _ := c.Reestructurar(ctx, &pb.ReestructuracionRequest{Planetas: listaPlanetasSTR, Vectores: listaVectoresSTR, Registrotxt: listaTxtSTR})
		fmt.Println("Respuesta de Fulcrum 2: " + Yellow + rReestructuracion.GetReply() + Reset)
		cancel()

		
		//Fulcrum 3
		fmt.Println("Enviando data a " + Yellow + "Fulcrum 3" + Reset)
		conn, err = grpc.Dial(Fulcrum3, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		c = pb.NewManejoComunicacionClient(conn)

		ctx, cancel = context.WithTimeout(context.Background(), time.Second)
		rReestructuracion, _ = c.Reestructurar(ctx, &pb.ReestructuracionRequest{Planetas: listaPlanetasSTR, Vectores: listaVectoresSTR, Registrotxt: listaTxtSTR})
		fmt.Println("Respuesta de Fulcrum 2: " + Yellow + rReestructuracion.GetReply() + Reset)
		cancel()


		// Avisarle al broker que reinicie las IPs
		fmt.Println("Contactando a " + Yellow + "Broker" + Reset)

		conn, err = grpc.Dial(BrokerAddress, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		c = pb.NewManejoComunicacionClient(conn)
		ctx, cancel = context.WithTimeout(context.Background(), time.Second)
		respBroker, _ := c.Comunicar(ctx, &pb.MessageRequest{Request: "Eliminar", Autor: "FulcrumDELETE"})
		fmt.Println("Respuesta broker: " + Yellow + respBroker.GetReply() + Reset)
		cancel()
		fmt.Println(Red + "=== FIN MERGE ===" + Reset)
		fmt.Println("")


	}
}

func main() {
	if LOCAL {
		BrokerAddress = "localhost:13370"
		Fulcrum2 = "localhost:13372"
		Fulcrum3 = "localhost:13373"
	} else {
		BrokerAddress = "137.184.61.128:13370"
		Fulcrum2 = "159.89.231.241:13372"
		Fulcrum3 = "159.223.157.5:13373"
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
	deleter()
	go ejecutarCoordinacion()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterManejoComunicacionServer(s, &ManejoComunicacionServer{})
	fmt.Println("Sevidor escuchando en puerto " + Yellow + port + Reset + "\n")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

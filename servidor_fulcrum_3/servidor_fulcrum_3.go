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

	pb "github.com/felipe-aguirre/tarea_sist_distribuido/protos"
	"google.golang.org/grpc"
)

const (
	port = ":13373"
	indiceServidor = 2
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


var vector = map[string]string{}

var listaPlanetas = []string{}
type ManejoComunicacionServer struct {
	pb.UnimplementedManejoComunicacionServer
}


func deleter() {
	dirname := "." + string(filepath.Separator)
      d, err := os.Open(dirname)
      if err != nil {
          fmt.Println(err)
          os.Exit(1)
      }
      
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
			d.Close()
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

func (s *ManejoComunicacionServer) Reestructurar(ctx context.Context, 
	in *pb.ReestructuracionRequest) (*pb.ReestructuracionReply, error) {
		fmt.Println(Red + "=== INICIO MERGE ===" + Reset)
		fmt.Println("Se recibió nueva data desde Fulcrum 1:")
		planetas :=in.GetPlanetas()
		texts := in.GetRegistrotxt()
		vectores := in.GetVectores()
		fmt.Println("Planetas: " + Yellow + planetas + Reset)
		fmt.Println("Logs: " + Yellow + texts + Reset)
		fmt.Println("Vectores: " + Yellow + vectores + Reset)
		fmt.Println("Borrando data actual . . .")
		deleter()
		
		fmt.Println("Data Borrada")
		listaPlanetasRecibida := strings.Split(planetas, ";")
		if len(listaPlanetasRecibida) == 1 {
			if strings.Compare(listaPlanetasRecibida[0], "") == 0{
				listaPlanetasRecibida = []string{}
			}
		}
		listaTexts := strings.Split(texts, ";")
		if len(listaTexts) == 1 {
			if strings.Compare(listaTexts[0], "") == 0{
				listaTexts = []string{}
			}
		}
		listaVectores := strings.Split(vectores, ";")
		if len(listaVectores) == 1 {
			if strings.Compare(listaVectores[0], "") == 0{
				listaVectores = []string{}
			}
		}
		vector = map[string]string{}
		fmt.Println("Escribiendo data nueva")
		for index, planeta := range listaPlanetasRecibida {
		fmt.Println("Planeta "+Yellow + planeta + Reset)
			listaTextsPlaneta := strings.Split(listaTexts[index], ",")
			file, _ := os.OpenFile("planeta_"+planeta+".txt", os.O_CREATE|os.O_WRONLY, 0660)
			if !(len(listaTextsPlaneta) == 1 && listaTextsPlaneta[0] == "-") {
				for _, elem := range listaTextsPlaneta {
					file.Write([]byte(elem + "\n"))
				}
			}
			
			file.Close()
			vector[planeta] = listaVectores[index]
		}
		
	fmt.Println(Red + "=== FIN MERGE ===" + Reset)
	fmt.Println("")
	return &pb.ReestructuracionReply{Reply: "Recibido"}, nil
}
func (s *ManejoComunicacionServer) Coordinar(ctx context.Context, 
	in *pb.CoordinacionRequest) (*pb.CoordinacionReply, error) {
		// Planetas del servidor X
		// Ejemplo: "planeta1,planeta2,planetae"

		// Logs de cada planeta concatenado con ; del servidor X
		// Ejemplo: "logp11,logp12,logp13;logp21,logp22"

		// Relojes de vector de cada planeta
		// Ejemplo: "RelojPlaneta1, RelojPlaneta2,RelojPlaneta3"
		// Cada RelojPlaneta tiene el formato a b c

  // Paso 1: Recolectar lista planetas
	listaPlanetasSTR := strings.Join(listaPlanetas, ";")

	// Paso 2: Obtener txts de cada planeta
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
		if len(logsDelPlanetaSTR) == 0 {
			logsDelPlanetaSTR = "-"
		}
		listaLogs = append(listaLogs, logsDelPlanetaSTR)
		fileCheck.Close()
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
	deleter()
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

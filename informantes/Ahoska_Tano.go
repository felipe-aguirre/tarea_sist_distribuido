package main

import (
	"context"
	"log"
	"time"

	pb "github.com/felipe-aguirre/tarea_sist_distribuido/protos"
	"google.golang.org/grpc"
)

// IP Del Broker
const (
	address = "localhost:50051"
)

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

	// Consola
	// Codigo que va a Servidor fulcrum
	/*
		exit := false
		loop := true
		for exit != loop {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Ahoska Tano > ")
			text, _ := reader.ReadString('\n')
			text = strings.Replace(text, "\r\n", "", -1)
			respuesta := strings.Split(text, " ")

			// Caso 1: AddCity
			if strings.Compare("AddCity", respuesta[0]) == 0 {

				// Configuracion de linea a escribir
				nombre_planeta := respuesta[1]
				nombre_ciudad := respuesta[2]
				linea_a_escribir := nombre_planeta + " " + nombre_ciudad

				if len(respuesta) > 3 {
					// Se quiere agregar el nuevo valor
					linea_a_escribir = linea_a_escribir + " " + respuesta[3]

				} else {
					linea_a_escribir = linea_a_escribir + " " + "0"
				}

				// Open File
				f, err := os.Open("../registro_planetario/" + nombre_planeta + ".txt")
				// En caso de que no exista, se crea y se agrega la linea necesaria
				if err != nil {
					f.Close()
					f, _ := os.OpenFile("../registro_planetario/"+nombre_planeta+".txt", os.O_CREATE|os.O_WRONLY, 0660)
					_, err = f.Write([]byte(linea_a_escribir + "\n"))
					fmt.Println("Se agrego la linea: ", linea_a_escribir)
					if err != nil {
						log.Fatal(err)
					}
					f.Close()
				} else {
					// El archivo si existe, se verifica primero si ya existe la ciudad
					scanner := bufio.NewScanner(f)
					NoExisteLinea := true
					for scanner.Scan() {
						linea_leida := strings.Split(scanner.Text(), " ")
						// Si la linea existe
						if strings.Compare(nombre_ciudad, linea_leida[1]) == 0 {
							fmt.Println("Ya existe la ciudad que se intenta agregar")
							NoExisteLinea = false
							break
						}
					}
					f.Close()
					if NoExisteLinea {
						f, _ := os.OpenFile("../registro_planetario/"+nombre_planeta+".txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0660)
						_, err = f.Write([]byte(linea_a_escribir + "\n"))
						fmt.Println("Se agrego la linea: ", linea_a_escribir)
						if err != nil {
							log.Fatal(err)
						}
					}
				}

			} else if strings.Compare("UpdateName", respuesta[0]) == 0 {
				fmt.Println("Se ejecuto updateName")
				nombre_planeta := respuesta[1]
				nombre_ciudad := respuesta[2]
				nuevo_valor := respuesta[3]
				f, _ := os.Open("../registro_planetario/" + nombre_planeta + ".txt")
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
				file, _ := os.OpenFile("../registro_planetario/"+nombre_planeta+".txt", os.O_CREATE|os.O_WRONLY, 0660)
				for _, elem := range textoCompleto {
					file.Write([]byte(elem + "\n"))
				}
				file.Close()

			} else if strings.Compare("UpdateNumber", respuesta[0]) == 0 {
				fmt.Println("Se ejecuto updateName")
				nombre_planeta := respuesta[1]
				nombre_ciudad := respuesta[2]
				nuevo_valor := respuesta[3]
				f, _ := os.Open("../registro_planetario/" + nombre_planeta + ".txt")
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
				file, _ := os.OpenFile("../registro_planetario/"+nombre_planeta+".txt", os.O_CREATE|os.O_WRONLY, 0660)
				for _, elem := range textoCompleto {
					file.Write([]byte(elem + "\n"))
				}
				file.Close()

			} else if strings.Compare("DeleteCity", respuesta[0]) == 0 {
				fmt.Println("Se ejecuto updateName")
				nombre_planeta := respuesta[1]
				nombre_ciudad := respuesta[2]
				f, _ := os.Open("../registro_planetario/" + nombre_planeta + ".txt")
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
				os.Remove("../registro_planetario/" + nombre_planeta + ".txt")
				file, _ := os.OpenFile("../registro_planetario/"+nombre_planeta+".txt", os.O_CREATE|os.O_WRONLY, 0660)
				for _, elem := range textoCompleto {
					file.Write([]byte(elem + "\n"))
				}
				file.Close()
			} else {
				if strings.Compare("exit", text) == 0 {
					exit = true
				} else {
					fmt.Println("Comando erroneo, intente con AddCity, UpdateName, UpdateNumber o DeleteCity")
				}
			}
		}
	*/
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	comando := "Prueba de comando 1"
	r, err := c.CreateNewUser(ctx, &pb.NewUser{Name: name, Age: age})
	for name, age := range new_users {
		r, err := c.CreateNewUser(ctx, &pb.NewUser{Name: name, Age: age})
		if err != nil {
			log.Fatalf("could not create user: %v", err)
		}
		log.Printf(`User Details:
			NAME: %s
			AGE: %d
			ID: %d`, r.GetName(), r.GetAge(), r.GetId())

	}

}

package panics

import "fmt"

/* defer espera a que concluya una acción para iniciar la realización de otra */

func ConnectDB(){
	fmt.Println("Conectando base de datos")
}

func ConsultData(){
	fmt.Println("Consultando datos")
}

func CloseConsultData(){
	fmt.Println("Cerrando conexión set de datos...")
}

func DisconnectDB(){
	fmt.Println("Cerrando la base de datos...")
}
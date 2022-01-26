package panics

func MainPanics(){

	// todos los diferidos se ejecutan al final
	// el primer defer declarado sera el ultimo en ejecutarse
	// los diferidos se ejecutan aunque la aplicación entre en pánico


	//conecta y desconecta la base de datos
	ConnectDB()					//1
	defer DisconnectDB()			//4

	// consulta la información
	ConsultData()				//2
	defer CloseConsultData()		//3

}
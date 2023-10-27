package main

import (
	"github.com/Grimmore/Server/modelos"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Función principal que ejecuta el resto de funciones del programa
func main() {
	//crea la base da datos y sus tablas
	modelos.Migrate()
	//Añade al array directorios que sera recorrido por el watcher los directorios de la base de datos

	/*
		modelos.NewForm("F1")
		modelos.NewQuestion("Q1", "1")
		modelos.NewAnswer("A1", "1")
		modelos.NewAnswer("A2", "1")
		modelos.NewArea("Ar1", 1, "1")
		modelos.NewArea("Ar2", 2, "1")
		modelos.NewArea("Ar3", 3, "1")
		modelos.NewArea("Ar4", 5, "2")
	*/

	//crea un server en el puerto especificado en el json y en el comparte via una api el contenido de la base de datos usando el formato json, a su vez mantiene en ejecución el resto del programa
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	router.Use(cors.New(config))
	Init(router)
	err := router.Run(":5447")
	if err != nil {
		panic(err)
	}
}

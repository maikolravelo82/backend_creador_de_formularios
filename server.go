package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Grimmore/Server/database"
	"github.com/Grimmore/Server/modelos"

	"github.com/gin-gonic/gin"
)

var AreaID string = ""
var AnswerID string = ""
var QuestionID string = ""
var FormID string = ""

func Init(router *gin.Engine) {

	web := router.Group("/web")

	//APIS DE LAS AREAS
	web.POST("/api/area/del", func(ctx *gin.Context) {
		data, _ := ctx.GetRawData()
		modelos.DeleteArea(string(data))

	})

	web.GET("/api/area/getID", func(ctx *gin.Context) {
		a, _ := modelos.GetAreaByPath(AreaID)
		if a != nil {
			//especifica que el contenido de la api es un json
			ctx.Request.Header.Add("Content-Type", "application/json")
			//escribe en la ruta de la api el json
			ctx.JSON(http.StatusOK, a)
		}
	})

	web.GET("/api/area/answer", func(ctx *gin.Context) {
		a, _ := modelos.GetAreaByAnswer(AreaID)
		if a != nil {
			//especifica que el contenido de la api es un json
			ctx.Request.Header.Add("Content-Type", "application/json")
			//escribe en la ruta de la api el json
			ctx.JSON(http.StatusOK, a)
		}
	})

	web.POST("/api/area/getID", func(ctx *gin.Context) {
		data, _ := ctx.GetRawData()
		AreaID = string(data)
	})

	web.POST("/api/area/add", func(ctx *gin.Context) {

		data, _ := ctx.GetRawData()

		type Area struct {
			Nombre   string
			Valor    int
			AnswerID string
		}
		b := Area{}

		json.Unmarshal(data, &b)

		_, s := modelos.NewArea(b.Nombre, b.Valor, b.AnswerID)
		if s != nil {
			log.Printf("Error al crear: %v", s)
		}

	})

	web.GET("/api/area/get", func(ctx *gin.Context) {
		db := database.GetDB()
		area := []modelos.Area{}
		//busca en la base de datos el modelo Evento
		err := db.Find(&area).Error
		if err != nil {
			log.Println(err)
			return
		}
		if err == nil {
			//especifica que el contenido de la api es un json
			ctx.Request.Header.Add("Content-Type", "application/json")
			//escribe en la ruta de la api el json
			ctx.JSON(http.StatusOK, area)
		}
	})

	//API DE LAS ANSWERS
	web.POST("/api/answer/del", func(ctx *gin.Context) {
		data, _ := ctx.GetRawData()
		modelos.DeleteAnswer(string(data))

	})

	web.GET("/api/answer/getID", func(ctx *gin.Context) {
		a, _ := modelos.GetAnswerByPath(AnswerID)
		if a != nil {
			//especifica que el contenido de la api es un json
			ctx.Request.Header.Add("Content-Type", "application/json")
			//escribe en la ruta de la api el json
			ctx.JSON(http.StatusOK, a)
		}
	})

	web.GET("/api/answer/question", func(ctx *gin.Context) {
		a, _ := modelos.GetAnswerByQuestion(AnswerID)
		if a != nil {
			//especifica que el contenido de la api es un json
			ctx.Request.Header.Add("Content-Type", "application/json")
			//escribe en la ruta de la api el json
			ctx.JSON(http.StatusOK, a)
		}
	})

	web.POST("/api/answer/getID", func(ctx *gin.Context) {
		data, _ := ctx.GetRawData()
		AnswerID = string(data)
	})

	web.POST("/api/answer/add", func(ctx *gin.Context) {

		data, _ := ctx.GetRawData()

		type Answer struct {
			Context    string
			Questionid string
		}
		b := Answer{}

		json.Unmarshal(data, &b)

		_, s := modelos.NewQuestion(b.Context, b.Questionid)
		if s != nil {
			log.Printf("Error al crear: %v", s)
		}

	})

	web.GET("/api/answer/get", func(ctx *gin.Context) {
		db := database.GetDB()
		answer := []modelos.Answer{}
		//busca en la base de datos el modelo Evento
		err := db.Find(&answer).Error
		if err != nil {
			log.Println(err)
			return
		}
		if err == nil {
			//especifica que el contenido de la api es un json
			ctx.Request.Header.Add("Content-Type", "application/json")
			//escribe en la ruta de la api el json
			ctx.JSON(http.StatusOK, answer)
		}
	})

	//API DE LAS QUESTIONS
	web.POST("/api/question/del", func(ctx *gin.Context) {
		data, _ := ctx.GetRawData()
		modelos.DeleteQuestion(string(data))

	})

	web.GET("/api/question/getID", func(ctx *gin.Context) {
		a, _ := modelos.GetQuestionByPath(QuestionID)
		if a != nil {
			//especifica que el contenido de la api es un json
			ctx.Request.Header.Add("Content-Type", "application/json")
			//escribe en la ruta de la api el json
			ctx.JSON(http.StatusOK, a)
		}
	})

	web.GET("/api/question/form", func(ctx *gin.Context) {
		a, _ := modelos.GetAnswerByQuestion(QuestionID)
		if a != nil {
			//especifica que el contenido de la api es un json
			ctx.Request.Header.Add("Content-Type", "application/json")
			//escribe en la ruta de la api el json
			ctx.JSON(http.StatusOK, a)
		}
	})

	web.POST("/api/question/getID", func(ctx *gin.Context) {
		data, _ := ctx.GetRawData()
		QuestionID = string(data)
	})

	web.POST("/api/question/add", func(ctx *gin.Context) {

		data, _ := ctx.GetRawData()

		type Question struct {
			Context string
			Formid  string
		}
		b := Question{}

		json.Unmarshal(data, &b)

		_, s := modelos.NewQuestion(b.Context, b.Formid)
		if s != nil {
			log.Printf("Error al crear: %v", s)
		}

	})

	web.GET("/api/question/get", func(ctx *gin.Context) {
		db := database.GetDB()
		question := []modelos.Question{}
		//busca en la base de datos el modelo Evento
		err := db.Find(&question).Error
		if err != nil {
			log.Println(err)
			return
		}
		if err == nil {
			//especifica que el contenido de la api es un json
			ctx.Request.Header.Add("Content-Type", "application/json")
			//escribe en la ruta de la api el json
			ctx.JSON(http.StatusOK, question)
		}
	})

	//API DE LOS FORM
	web.POST("/api/form/del", func(ctx *gin.Context) {
		data, _ := ctx.GetRawData()
		modelos.DeleteForm(string(data))

	})

	web.GET("/api/form/getID", func(ctx *gin.Context) {
		a, _ := modelos.GetFormByPath(FormID)
		if a != nil {
			//especifica que el contenido de la api es un json
			ctx.Request.Header.Add("Content-Type", "application/json")
			//escribe en la ruta de la api el json
			ctx.JSON(http.StatusOK, a)
		}
	})

	web.POST("/api/form,/getID", func(ctx *gin.Context) {
		data, _ := ctx.GetRawData()
		FormID = string(data)
	})

	web.POST("/api/form/add", func(ctx *gin.Context) {

		data, _ := ctx.GetRawData()

		type Form struct {
			Nombre string
		}
		b := Form{}

		json.Unmarshal(data, &b)

		_, s := modelos.NewForm(b.Nombre)
		if s != nil {
			log.Printf("Error al crear: %v", s)
		}

	})

	web.GET("/api/form/get", func(ctx *gin.Context) {
		db := database.GetDB()
		form := []modelos.Form{}
		//busca en la base de datos el modelo Evento
		err := db.Find(&form).Error
		if err != nil {
			log.Println(err)
			return
		}
		if err == nil {
			//especifica que el contenido de la api es un json
			ctx.Request.Header.Add("Content-Type", "application/json")
			//escribe en la ruta de la api el json
			ctx.JSON(http.StatusOK, form)
		}
	})

}

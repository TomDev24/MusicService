package main

import (
	"net/http"
	"time"
	"fmt"
	"log"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"

)

type Song struct {
	Id			int			`json:"id" fake:"{skip}"`
	Group		string		`json:"group" fake:"{name}"`
	Song		string		`json:"song" fake:"{name}"`
	ReleaseDate	time.Time	`json:"releaseDate"`
	Text		string		`json:"text" fake:"{sentence:10}"`
	Link		string		`json:"link"`
}

type Config struct {
	Port	string `envconfig:"FAKE_API_PORT" required:"true"`
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	cfg := Config{}
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	r.GET("/info", getInfo)

	r.Run(cfg.Port)
}

func getInfo(c *gin.Context) {
	// group and song query args
	var songObject Song

	group := c.Query("group")
	song := c.Query("song")

	_ = gofakeit.Struct(&songObject)

	songObject.Group = group
	songObject.Song = song

	c.IndentedJSON(http.StatusOK, songObject)
}

func debug() {
	var song Song

	err := gofakeit.Struct(&song)
	log.Fatal(err)

	fmt.Println("------")
	fmt.Println(song)
	fmt.Println("------")
}

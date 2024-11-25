package server

import (
	"log"
	"fmt"
	"net/http"
	"encoding/json"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/TomDev24/MusicServiceGo/internal/model"
	"github.com/TomDev24/MusicServiceGo/pkg/pagination"
)

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func (s *Server) getAll (c *gin.Context) {
	var filterFields model.Song

	filterFields.Group = c.Query("group")
	filterFields.Song = c.Query("song")

	log.Println(filterFields)
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 10
	}

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}

	pagination := &pagination.Pagination{
		Limit: limit,
		Page: page,
	}
	pagination, err = model.AllSong(&filterFields, s.db, pagination)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "")
		return
	}

	c.JSON(http.StatusOK, pagination)
}

func (s *Server) deleteById (c *gin.Context) {
	var song model.Song

	id := c.Params.ByName("id")
	err := model.DeleteSong(s.db, id, &song)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "")
		return
	}
	c.JSON(http.StatusOK, song)
}

func (s *Server) getById (c *gin.Context) {
	var song model.Song

	id := c.Params.ByName("id")
	err := model.GetSongById(s.db, id, &song)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "")
		return
	}

	c.JSON(http.StatusOK, song)
}

func (s *Server) editById (c *gin.Context) {
	var song model.Song

	id := c.Params.ByName("id")
	err := model.GetSongById(s.db, id, &song)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "")
		return
	}

	err = c.BindJSON(&song)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "")
		return
	}

	model.SaveSong(s.db, &song)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "")
		return
	}

	c.JSON(http.StatusOK, song)
}

func (s *Server) create (c *gin.Context) {
	var song model.Song

	err := c.BindJSON(&song)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "")
		return
	}

	r, err := http.Get(fmt.Sprintf("%s?group=%s&song=%s", s.cfg.Serv.Endpoint, song.Group, song.Song))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, "")
		return
	}

	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(&song)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "")
		return
	}

	model.SaveSong(s.db, &song)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "")
		return
	}

	c.JSON(http.StatusOK, song)
}

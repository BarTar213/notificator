package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/BarTar213/notificator/models"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

const (
	invalidRequestBodyErr = "invalid request body"
	invalidIdParamErr     = "invalid param - ID"
	invalidNameParamErr   = "invalid param - name"
	invalidReadParamErr   = "invalid query param - read"
)

func handlePostgresError(c *gin.Context, l *log.Logger, err error, resource string) {
	if err == pg.ErrNoRows {
		c.JSON(http.StatusNotFound, models.Response{Error: fmt.Sprintf("%s with given information doesn't exists", resource)})
		return
	}
	l.Println(err)

	msg := ""
	pgErr, ok := err.(pg.Error)
	if ok {
		status := http.StatusBadRequest

		switch pgErr.Field('C') {
		case "23503":
			status = http.StatusNotFound
			msg = fmt.Sprintf("%s with given information doesn't exists", resource)
		case "23505":
			msg = fmt.Sprintf("%s with given information already exists", resource)
		}
		if len(msg) > 0 {
			c.JSON(status, models.Response{Error: msg})
			return
		}
	}

	c.JSON(http.StatusInternalServerError, models.Response{Error: "storage error"})
}

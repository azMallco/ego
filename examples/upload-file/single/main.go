package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/go-ego/ego"
)

func main() {
	router := ego.Default()
	router.Static("/", "./public")
	router.POST("/upload", func(c *ego.Context) {
		name := c.PostForm("name")
		email := c.PostForm("email")

		// Source
		file, _ := c.FormFile("file")
		src, _ := file.Open()
		defer src.Close()

		// Destination
		dst, _ := os.Create(file.Filename)
		defer dst.Close()

		// Copy
		io.Copy(dst, src)

		c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully with fields name=%s and email=%s.", file.Filename, name, email))
	})
	router.Run(":8080")
}

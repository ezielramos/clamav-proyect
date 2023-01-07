package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/api/ping", PingResponse)
	r.POST("api/scan", ScanArchive)

	if err := r.Run(":8080"); nil != err {
		panic(err.Error())
	}
}

func ScanArchive(c *gin.Context) {

	err := c.Request.ParseForm()
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	} else {

		path := c.Request.FormValue("file")

		form := new(bytes.Buffer)
		writer := multipart.NewWriter(form)
		fw, err := writer.CreateFormFile(path, filepath.Base(path))
		if err != nil {
			log.Fatal(err)
		}
		fd, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		defer fd.Close()
		_, err = io.Copy(fw, fd)
		if err != nil {
			log.Fatal(err)
		}

		writer.Close()

		client := &http.Client{}
		req, err := http.NewRequest("POST", "http://localhost:9000/scan", form)
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Set("Content-Type", writer.FormDataContentType())
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		bodyText, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%s\n", bodyText)

	}
	c.String(http.StatusOK, "scan succes")
}

func PingResponse(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

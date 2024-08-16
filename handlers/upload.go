package handlers

import (
	"fmt"
	"fofa-derp/models"
	"fofa-derp/utils"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func serveHomePage(c *gin.Context) {
	c.Header("Content-Type", "text/html")
	c.String(http.StatusOK, `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Upload File</title>
		</head>
		<body>
			<h1>Upload JSON File</h1>
			<form enctype="multipart/form-data" action="/upload" method="post">
				<input type="file" name="file" />
				<input type="submit" value="Upload" />
			</form>
		</body>
		</html>
	`)
}

func handleFileUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Get form err: %s", err.Error()))
		return
	}

	data, err := file.Open()
	defer data.Close()
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to open uploaded file: %s", err.Error()))
		return
	}

	bytes, err := ioutil.ReadAll(data)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to read file: %s", err.Error()))
		return
	}

	records, err := models.UnmarshalRecords(bytes)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to unmarshal JSON: %s", err.Error()))
		return
	}

	outputData := utils.ProcessRecords(records)            // Make sure this function call matches the utils package
	err = utils.WriteJSONToFile("result.json", outputData) // Make sure this function call matches the utils package
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to write file result.json: %s", err.Error()))
		return
	}

	c.FileAttachment("result.json", "result.json")
}

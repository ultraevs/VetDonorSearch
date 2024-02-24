package controller

import (
	"app/internal/model"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

// Message Сообщение ассистенту.
// @Summary Сообщение ассистенту
// @Description Сообщение ассистенту.
// @Accept json
// @Produce json
// @Param request body model.MessageBody true "Сообщение ассистенту"
// @Success 200 {object} model.CodeResponse "Сообщене получено"
// @Failure 400 {object} model.ErrorResponse "Не удалось получить сообщение"
// @Tags Assistant
// @Router /v1/message [post]
func Message(context *gin.Context) {
	var mes model.MessageBody
	if err := context.ShouldBindJSON(&mes); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Can't read the body"})
		return
	}
	// Script
	pythonScript := "./internal/api/python-assistant/text_processing.py"
	cmd := exec.Command("py", "-m", "pip", "install", "-r", "go-backend/internal/api/python-assistant/requirements.txt")
	output, _ := cmd.Output()
	fmt.Println(output)
	cmd = exec.Command("py", pythonScript, mes.MessageText)
	output, _ = cmd.Output()
	resultPath := filepath.Join(filepath.Dir(pythonScript), "result.json")
	file, err := os.Open(resultPath)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	var data interface{}
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		fmt.Println("Ошибка при декодировании JSON:", err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"callback": data})
}

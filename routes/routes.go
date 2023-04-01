package wwf

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type FileAction struct {
	OldName string `json:"old_name"`
	NewName string `json:"new_name"`
}

type FileEntry struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type PathRequest struct {
	Path string `json:"path"`
}

type FileNameRequest struct {
	FileName string `json:"file_name"`
}

func RegisterRoutes(api *gin.RouterGroup) {
	api.POST("/files", GetFiles)
	api.POST("/read_file", ReadFile)
	api.POST("/rename_file", RenameFile)
	api.POST("/delete_file", DeleteFile)
	api.POST("/create_file", CreateFile)
}

// * Получение файлов
func GetFiles(c *gin.Context) {
	var pathRequest PathRequest
	if err := c.ShouldBindJSON(&pathRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	files, err := ioutil.ReadDir(pathRequest.Path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var fileEntries []FileEntry
	for _, f := range files {
		fileType := "file"
		if f.IsDir() {
			fileType = "directory"
		}
		fileEntries = append(fileEntries, FileEntry{Name: f.Name(), Type: fileType})
	}

	c.JSON(http.StatusOK, gin.H{"files": fileEntries})
}

// * Чтение файлов
func ReadFile(c *gin.Context) {
	var fileNameRequest FileNameRequest
	if err := c.ShouldBindJSON(&fileNameRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	content, err := ioutil.ReadFile(fileNameRequest.FileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"content": string(content)})
}

// * Переменовать файл
func RenameFile(c *gin.Context) {
	var fileAction FileAction
	if err := c.ShouldBindJSON(&fileAction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := os.Rename(fileAction.OldName, fileAction.NewName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File renamed successfully"})
}

// * Удалить файл
func DeleteFile(c *gin.Context) {
	var fileNameRequest FileNameRequest
	if err := c.ShouldBindJSON(&fileNameRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := os.Remove(fileNameRequest.FileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File deleted successfully"})
}

// * Создать файл
func CreateFile(c *gin.Context) {
	type CreateFileRequest struct {
		FileName string `json:"file_name"`
		Content  string `json:"content"`
	}

	var createFileRequest CreateFileRequest
	if err := c.ShouldBindJSON(&createFileRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ioutil.WriteFile(createFileRequest.FileName, []byte(createFileRequest.Content), 0644)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File created successfully"})
}

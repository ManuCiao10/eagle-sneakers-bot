package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"net/http"
	"os/user"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jaypipes/ghw"
)

func newSHA256(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

func generateHWID() string {
	block, _ := ghw.Block()
	var disks []string
	for _, disk := range block.Disks {
		disks = append(disks, disk.SerialNumber)
	}

	userStruct, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}

	username := userStruct.Username

	return newSHA256(strings.Join(disks, ",") + "," + username)
}

func api() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

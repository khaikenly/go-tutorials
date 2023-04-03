package main

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Data  any  `json:"data"`
	Error bool `json:"error"`
}

var (
	ACCOUNT gin.Accounts
	SALT    []byte = []byte("saltsecret")
)

func main() {
	AccountsInit()
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	cli.NegotiateAPIVersion(ctx)
	client.IsErrConnectionFailed(err)

	// default Gin fw
	router := gin.Default()
	//static path
	router.Static("/assets", "./assets")
	//load html file
	router.LoadHTMLGlob("./templates/**/*")
	// basic auth
	auth := router.Group("/", gin.BasicAuth(ACCOUNT), func(c *gin.Context) {
		fmt.Println("1213213213")
		user, password, hasAuth := c.Request.BasicAuth()
		fmt.Println(user, password, hasAuth)
	})

	// static files
	auth.GET("/", func(c *gin.Context) {
		cmd := exec.Command("docker", "version")
		output, _ := cmd.CombinedOutput()
		s := fmt.Sprintf("%q", string(output))
		s = strings.Replace(s, `\n`, " <br>", -1)
		i := strings.Index(s, "Server")
		var client_version, server_version string
		if i < 0 {
			i = strings.Index(s, "Client")
			server_version, client_version = "Docker daemon is not running!", s[i:]
		} else {
			client_version, server_version = s[:i], s[i:]
		}
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":          "Docker version",
			"client_version": template.HTML(strings.ReplaceAll(client_version, `"`, "")),
			"server_version": template.HTML(strings.ReplaceAll(server_version, `"`, "")),
		})
	})

	// handlers
	auth.GET("/containers", func(c *gin.Context) {
		containers, err := getContainersStart(ctx, cli)
		if err != nil {
			c.JSON(http.StatusOK, Response{Data: nil, Error: true})
			panic(err)
		}
		c.JSON(http.StatusOK, Response{Data: containers, Error: false})
	})

	// Start the HTTP server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Unable to start HTTP server: %v\n", err)
	}
}

func getContainersStart(ctx context.Context, cli *client.Client) ([]types.Container, error) {
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	return containers, err
}

func AccountsInit() {
	ACCOUNT = gin.Accounts{"admin": "admin123"}
}

func ValidMAC(message, messageMAC, key []byte) string {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	expectedMAC := mac.Sum(nil)
	return string(expectedMAC)
}

package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// RequestInfo represents the structure of the response
type RequestInfo struct {
	Host struct {
		Hostname string   `json:"hostname"`
		IP       string   `json:"ip"`
		IPs      []string `json:"ips"`
	} `json:"host"`
	HTTP struct {
		Method      string `json:"method"`
		BaseURL     string `json:"baseUrl"`
		OriginalURL string `json:"originalUrl"`
		Protocol    string `json:"protocol"`
	} `json:"http"`
	Request struct {
		Params  map[string]string      `json:"params"`
		Query   map[string]string      `json:"query"`
		Cookies map[string]string      `json:"cookies"`
		Body    map[string]interface{} `json:"body"`
		Headers map[string]string      `json:"headers"`
	} `json:"request"`
	Environment struct {
		Path     string `json:"PATH"`
		Hostname string `json:"HOSTNAME"`
		Home     string `json:"HOME"`
	} `json:"environment"`
}

func main() {
	e := echo.New()
	e.HideBanner = true

	// Add middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route handler
	e.Any("/*", handleRequest)

	// Add a route to set a test cookie
	e.GET("/setcookie", func(c echo.Context) error {
		cookie := new(http.Cookie)
		cookie.Name = "testcookie"
		cookie.Value = "testvalue"
		c.SetCookie(cookie)
		return c.String(http.StatusOK, "Cookie set!")
	})

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	serverAddr := ":" + port
	log.Printf("Server listening on port %s", port)
	if err := e.Start(serverAddr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func handleRequest(c echo.Context) error {
	var requestInfo RequestInfo

	// Host information
	requestInfo.Host.Hostname = c.Request().Host
	if host, _, err := net.SplitHostPort(requestInfo.Host.Hostname); err == nil {
		requestInfo.Host.Hostname = host
	}
	ip, _ := getOutboundIP()
	requestInfo.Host.IP = ip
	requestInfo.Host.IPs = []string{}

	// HTTP information
	requestInfo.HTTP.Method = c.Request().Method
	requestInfo.HTTP.BaseURL = ""
	requestInfo.HTTP.OriginalURL = c.Request().URL.String()
	requestInfo.HTTP.Protocol = "http"

	// Request information
	requestInfo.Request.Params = make(map[string]string)
	pathParts := strings.Split(c.Request().URL.Path, "/")
	var params []string
	for _, part := range pathParts {
		if part != "" {
			params = append(params, part)
		}
	}
	// Add parameters with numeric indices
	for i, param := range params {
		requestInfo.Request.Params[strconv.Itoa(i)] = param
	}

	requestInfo.Request.Query = make(map[string]string)
	for key, values := range c.QueryParams() {
		if len(values) > 0 {
			requestInfo.Request.Query[key] = values[0]
		}
	}

	// Enhanced cookie handling
	requestInfo.Request.Cookies = make(map[string]string)
	cookies := c.Cookies()
	for _, cookie := range cookies {
		requestInfo.Request.Cookies[cookie.Name] = cookie.Value
	}

	requestInfo.Request.Body = make(map[string]interface{})
	requestInfo.Request.Headers = make(map[string]string)
	for k, v := range c.Request().Header {
		if len(v) > 0 {
			requestInfo.Request.Headers[strings.ToLower(k)] = v[0]
		}
	}

	// Environment information
	requestInfo.Environment.Path = os.Getenv("PATH")
	requestInfo.Environment.Hostname = os.Getenv("HOSTNAME")
	requestInfo.Environment.Home = os.Getenv("HOME")

	return c.JSON(200, requestInfo)
}

// getOutboundIP gets the preferred outbound IP of this machine
func getOutboundIP() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String(), nil
}

package main

import "fmt"


////////////////// Consts ////////////////////////

	const (
		appStatusURL  = "/app/status"
		createuserURL = "/create/user"
	)

///////////////// Server Interface ///////////////

type server interface {
	handleRequest(string, string) (int, string)
}

////////////////// Implementations ///////////////

/////////////////// Nginx Server /////////////////

type nginx struct {
	application       *application
	maxAllowedRequest int
	rateLimiter       map[string]int
}

func newNginxServer() *nginx {
	return &nginx{
		application:       &application{}, // embed application server
		maxAllowedRequest: 2,
		rateLimiter:       make(map[string]int),
	}
}

func (n *nginx) handleRequest(url, method string) (int, string) {
  if allowed := n.checkRateLimiting(url); !allowed {
		return 403, "Not Allowed"
	}
	return n.application.handleRequest(url, method) // request is forwarded to application server
}

func (n *nginx) checkRateLimiting(url string) bool {
	if n.rateLimiter[url] == 0 {
		n.rateLimiter[url] = 1
	}
	if n.rateLimiter[url] > n.maxAllowedRequest {
		return false
	}
	n.rateLimiter[url] = n.rateLimiter[url] + 1
	return true
}

///////////////// Application Server /////////////

type application struct {
}

func (a *application) handleRequest(url, method string) (int, string) {
	if url == "/app/status" && method == "GET" {
		return 200, "Ok"
	}
	if url == "/create/user" && method == "POST" {
		return 201, "User Created"
	}
	return 404, "Not Ok"
}

/////////////////////// Main /////////////////////

func main() {
	// Nginx Server
	nginxServer := newNginxServer()
	httpCode, body := nginxServer.handleRequest(appStatusURL, "GET")
  printHTTPStatus(appStatusURL, httpCode, body)
	httpCode, body = nginxServer.handleRequest(appStatusURL, "GET")
  printHTTPStatus(appStatusURL, httpCode, body)
	httpCode, body = nginxServer.handleRequest(appStatusURL, "GET")
  printHTTPStatus(appStatusURL, httpCode, body)
	httpCode, body = nginxServer.handleRequest(createuserURL, "POST")
  printHTTPStatus(createuserURL, httpCode, body)
	httpCode, body = nginxServer.handleRequest(createuserURL, "GET")
  printHTTPStatus(createuserURL, httpCode, body)
}

///////////////////// Helper /////////////////////

func printHTTPStatus(url string, httpCode int, body string) {
	fmt.Printf("Url: %s\nHttpCode: %d\nBody: %s\n\n", url, httpCode, body)
}

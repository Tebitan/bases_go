package webserver

import "net/http"

func MiWebServer() {
	//routing
	http.HandleFunc("/", home)
	//Config port webserver
	http.ListenAndServe(":3000", nil)
}

func home(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "./webserver/index.html")
}

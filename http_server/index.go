package http_server

import (
	"net/http"
	"rnd-surajan-go-http/env"
)

func HttpServer() {
	// 💡 Note: This will serve our "index.html" file inside "static" folder on path: "http://localhost:8080/".
	// fs := http.FileServer(http.Dir("static/"))
	// http.Handle("/", fs)

	// 💡 Note: This will serve our "index.html" file inside "static" folder on path: "http://localhost:8080/static".
	fs := http.FileServer(http.Dir("static/"))
	/*
		Before, our file server handled all requests from "/" path.
		We can modify our application so that it only handles requests with the path pattern "/static/".
	*/
	/*
		Since our static directory is the root of the file system, we must remove the "/static/" prefix from the request path before searching for
		the specified file.
	*/
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(env.GetBaseUrl(), nil)
}

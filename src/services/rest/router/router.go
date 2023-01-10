package router

import (
	"cupcake-core/src/logger"
	"cupcake-core/src/repository"
	"cupcake-core/src/services/rest/api"
	"net/http"
)

/*
The "router" package provides a "Router" structure that allows handling HTTP requests in a structured way.
The "Router" structure has two fields: "Logger" and "Repository".
The "Logger" field is of type "logger.HTTPRequestLogger" and is used to log information about HTTP requests.
The "Repository" field is of type "repository.Repository" and is used to access a data repository.
*/

type Router struct {
	Logger     logger.HTTPRequestLogger
	Repository repository.Repository
}

/*
The "Router" structure provides a "New" function that is used to create
a new instance of "Router" and configure the "Logger" and "Repository" fields.
*/

func New(l logger.HTTPRequestLogger, repo repository.Repository) *Router {
	return &Router{
		Logger:     l,
		Repository: repo,
	}
}

/*
The "Router" structure implements the "http.Handler" interface, which allows assigning
an instance of "Router" as HTTP request handler through the "ServeHTTP" function.
This function parses the "Path" of the HTTP request URL to determine what action to take.
Currently, it has two instances, one for the "/backoffice" and one for the "/front-office".

For each case, depending on the endpoint approach, it will receive as a parameter a repository,
a logger, a http.ResponseWriter and a http.Request to process the response and log the request.
*/

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/backoffice":
		// TODO: ...
	case "/frontoffice":
		// TODO: ...
	default:
		api.Index(router.Logger, w, r)
	}
}

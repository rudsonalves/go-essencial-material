package main

import (
	"go-list/internal/shared/middleware"
	"net/http"
)

type routeRegister interface {
	HandleFunc(pattern string, handler http.HandlerFunc)
}

type routeGroup struct {
	mux         *http.ServeMux
	middlewares []middleware.Middleware
}

func (g routeGroup) HandleFunc(pattern string, handler http.HandlerFunc) {
	finalHandler := middleware.Chain(
		handler,
		g.middlewares...,
	)

	g.mux.Handle(pattern, finalHandler)
}

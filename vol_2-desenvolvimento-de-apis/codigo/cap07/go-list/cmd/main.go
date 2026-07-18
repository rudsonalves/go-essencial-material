package main

import (
	"go-list/internal/shared/middleware"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	common := []middleware.Middleware{
		middleware.RequestID,
		middleware.Logging,
		middleware.Recover,
		middleware.CORS,
	}

	protected := []middleware.Middleware{
		middleware.Auth,
	}

	publicRoutes := routeGroup{
		mux: mux,
	}

	authenticatedRoutes := routeGroup{
		mux:         mux,
		middlewares: protected,
	}

	registerPublicRoutes(publicRoutes)
	registerAuthenticatedRoutes(authenticatedRoutes)

	handler := middleware.Chain(mux, common...)

	if err := http.ListenAndServe(":8080", handler); err != nil {
		panic(err)
	}
}

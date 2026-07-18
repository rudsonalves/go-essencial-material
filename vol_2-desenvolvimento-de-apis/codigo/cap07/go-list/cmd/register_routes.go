package main

func registerPublicRoutes(router routeRegister) {
	router.HandleFunc("POST /users", createUser)
}

func registerAuthenticatedRoutes(router routeRegister) {
	registerUserRoutes(router)
	registerListRoutes(router)
	registerItemRoutes(router)
	// registerTestsRoutes(router)
}

func registerUserRoutes(router routeRegister) {
	router.HandleFunc("GET /users", listUsers)
	router.HandleFunc("GET /users/{id}", getUser)
	router.HandleFunc("PUT /users/{id}", updateUser)
	router.HandleFunc("DELETE /users/{id}", deleteUser)
}

func registerListRoutes(router routeRegister) {
	router.HandleFunc("GET /lists", listLists)
	router.HandleFunc("POST /lists", createList)
	router.HandleFunc("GET /lists/{id}", getList)
	router.HandleFunc("PUT /lists/{id}", updateList)
	router.HandleFunc("DELETE /lists/{id}", deleteList)
}

func registerItemRoutes(router routeRegister) {
	router.HandleFunc("POST /lists/{id}/items", createItem)
	router.HandleFunc("GET /items/{id}", getItem)
	router.HandleFunc("PATCH /items/{id}", updateItem)
	router.HandleFunc("DELETE /items/{id}", deleteItem)
}

// func registerTestsRoutes(router routeRegister) {
// 	router.HandleFunc("/debug/panic", simulatePanic)
// }

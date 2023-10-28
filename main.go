package main

import (
	// "echo/biz"
	// "echo/rest"
	"echo/utils"
	"net/http"
)

// TODO: remove this
// func hello(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Hello World"))
// }

func main() {
	migrationDbConfig := utils.GetMigrationDbConfig()
	err := utils.Migrate(migrationDbConfig, "./migrations")
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	// TODO: remove this
	// mux.HandleFunc("/health", hello)
	http.ListenAndServe(":8080", mux)

	// userRepository := &biz.UserRepositoryFake{}
	// rest.SetupServicesControllers(userRepository).Run()
}

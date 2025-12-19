package main

// Librairie pour exécuter des requêtes SQL

// func main() {
// 	database, openError := sql.Open("sqlite", "./database.sqlite")

// 	if openError != nil {
// 		log.Fatal("Error while opening the database:", openError)
// 	}
// 	initializeDataModel(database)

// 	// Les routes HTTP
// 	http.HandleFunc("POST /signup", routes.Signup(database))

// 	fmt.Println("Starting http server on http://localhost:8080...")

// 	// On démarre le serveur HTTP et on vérifie si aucune erreur ne survient
// 	if listenError := http.ListenAndServe(":8080", nil); listenError != nil {
// 		// Si une erreur survient
// 		log.Fatal("Error while listening for http requests on http://localhost:8080", listenError)
// 	}
// }

// //// Début de chaque fonction
// func Signin(database *sql.DB) http.HandlerFunc {
// 	return func(response http.ResponseWriter, request *http.Request) {

// /// Creation de la table
// func initializeDataModel(database *sql.DB) {
//     _, execError := database.Exec(`

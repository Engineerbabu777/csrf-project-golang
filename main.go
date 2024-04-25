package main

import "log"

var host = "localhost"
var port = "9000"

func main() {

	db.InitDb()

	
	jwtErr := myJwt.InitJWT()

	if jwtErr != nil {
		log.Println("Error initializing JWT:", jwtErr)
		log.Fatal(jwtErr)
	}

	


}

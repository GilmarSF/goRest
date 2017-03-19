package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)
type Payload struct {
	Stuff Data
}

type Data struct {
	Fruit Fruits
	Veggies Vegetables
}
type Fruits map[string]int
type Vegetables map[string]int

func serveRest(w http.ResponseWriter, r *http.Request) {
	response, err := getJsonResponse()

	if err != nil {
		panic (err)
	}

	fmt.Fprintf(w,string(response))
}
func main() {
	log.Println("Inicio!!!")
     http.HandleFunc("/", serveRest)
     http.ListenAndServe("localhost:1337", 	nil)
}
 func getJsonResponse()([]byte, error) {
      fruits := make(map[string]int)
	 fruits["Macas"] = 25
	 fruits["Laranjas"] = 11
	 fruits["Peras"] = 13

	 vegetables := make(map[string]int)
	 vegetables["Abobrinhas"] = 21
	 vegetables["Xuxu"] = 3
	 vegetables["Berinjela"] = 18
	 vegetables["Vagem"] = 9

	 d:= Data{fruits, vegetables}
	 p:= Payload{d}

	 retorno, errr := json.MarshalIndent(p, "", "  ")
	 log.Println("Retorno: " + string(retorno))
	 return retorno, errr
 }
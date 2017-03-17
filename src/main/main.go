package main

import (

	//"fmt"
	data "srjc_database"
	"net/http"
	"log"
)


func main(){
	http.HandleFunc("/", schedulerHandler)
	//
	//
	//http.HandleFunc("/penis",func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "sending this json ")
	//})
	data.InitDB()
	//if err != nil {
	//	log.Panic("SONOFABITCH")
	//}else{
	//	println("global database initialized ")
	//	data.AllSections()
	//}
	defer log.Fatal(http.ListenAndServe(":8084", nil))

}
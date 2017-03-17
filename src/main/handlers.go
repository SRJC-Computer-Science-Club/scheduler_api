package main

import (
	"net/http"
	"fmt"
)

func schedulerHandler (w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "sending the other json")
}

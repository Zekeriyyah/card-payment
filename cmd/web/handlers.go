package main

import (
	"fmt"
	"net/http"
)

func (app *application) VirtualTerminal(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HIT THE HANDLER")
}

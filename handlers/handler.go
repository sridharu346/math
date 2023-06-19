package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sridharu346/calculations/schema"
	"strconv"
)

func Math(w http.ResponseWriter, r *http.Request) {
	
	//calling method
	if r.Method != "GET"{
		ResponseFormat(w, "nil", 0, 0, "StatusMethodNotAllowed", http.StatusMethodNotAllowed)
		return 
	}

	//reading the body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ResponseFormat(w, "nil", 0, 0, "StatusMethodNotAllowed", http.StatusMethodNotAllowed)
	}

	// unmarshal the body
	var operations schema.MathOperations
	err = json.Unmarshal(body, operations)
	if err != nil {
		ResponseFormat(w, "nil", 0, 0, "StatusMethodNotAllowed", http.StatusMethodNotAllowed)
	}

	//validating the code
	validate, err := validator.IsInteger(operations)
	if err != nil {
		if err != nil {
			ResponseFormat(w, "nil", 0, 0, "StatusMethodNotAllowed", http.StatusMethodNotAllowed)
		}

	}

	




	

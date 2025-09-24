package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type Result struct {
	Answer string
}

func calculatorHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))

	if r.Method == http.MethodPost {
		num1, err1 := strconv.ParseFloat(r.FormValue("num1"), 64)
		num2, err2 := strconv.ParseFloat(r.FormValue("num2"), 64)
		op := r.FormValue("op")

		if err1 != nil || err2 != nil {
			tmpl.Execute(w, Result{Answer: "Invalid input"})
			return
		}

		var result string
		switch op {
		case "add":
			result = fmt.Sprintf("%f", num1+num2)
		case "sub":
			result = fmt.Sprintf("%f", num1-num2)
		case "mul":
			result = fmt.Sprintf("%f", num1*num2)
		case "div":
			if num2 != 0 {
				result = fmt.Sprintf("%f", num1/num2)
			} else {
				result = "Error: Divide by zero"
			}
		default:
			result = "Unknown operation"
		}

		tmpl.Execute(w, Result{Answer: result})
		return
	}

	// GET request (first load)
	tmpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", calculatorHandler)
	fmt.Println("Calculator running on :8080")
	http.ListenAndServe(":8080", nil)
}

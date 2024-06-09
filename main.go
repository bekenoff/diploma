package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type Coffee struct {
	Name string  `json:"name"`
	Kwt  float64 `json:"kwt"`
}

type KwtResponse struct {
	Kwt float64 `json:"kwt"`
}

func main() {

	http.HandleFunc("/Coffee-machine", getCoffee)
	http.HandleFunc("/Fridge", getFridge)
	http.HandleFunc("/Freezer", getFreezer)
	http.HandleFunc("/Plate", getPlate)
	http.HandleFunc("/Washing", getWasher)
	http.HandleFunc("/Toaster", getToaster)
	http.HandleFunc("/Tech", getTechnic)

	// BY MODEL

	http.HandleFunc("/Coffee-machine-by-model", getCoffeeByModel)
	http.HandleFunc("/Fridge-by-model", getFridgeByModel)
	http.HandleFunc("/Freezer-by-model", getFreezerByModel)
	http.HandleFunc("/Plate-by-model", getPlateByModel)
	http.HandleFunc("/Washer-by-model", getWasherByModel)
	http.HandleFunc("/Toaster-by-model", getToasterByModel)

	log.Println("Server is running on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

// BY MODEL

func getPlateByModel(w http.ResponseWriter, r *http.Request) {
	// Extracting the name query parameter from the request
	name := r.URL.Query().Get("name")

	db, err := sql.Open("mysql", "root:zikRerSPppEEPJZUeawwtpMpyCmpOmtK@tcp(monorail.proxy.rlwy.net:22986)/railway")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// SQL query with WHERE clause using LIKE for partial matching
	query := `
        SELECT kwt 
        FROM Plate
        WHERE name LIKE ?
        ORDER BY CAST(REPLACE(kwt, ' кВт', '') AS UNSIGNED) ASC
        LIMIT 1
    `

	// Adjusting the search term to include the wildcard character
	name = name + "%"

	row := db.QueryRow(query, name)

	var kwtStr string
	err = row.Scan(&kwtStr)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "No matching records found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	kwt, err := parseKwt(kwtStr)
	if err != nil {
		http.Error(w, "Error converting kwt to float64", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(KwtResponse{Kwt: kwt})
}

func getWasherByModel(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	db, err := sql.Open("mysql", "root:zikRerSPppEEPJZUeawwtpMpyCmpOmtK@tcp(monorail.proxy.rlwy.net:22986)/railway")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	query := `
        SELECT kwt 
        FROM WashingMachine
        WHERE name LIKE ?
        ORDER BY CAST(REPLACE(kwt, ' кВт/ч', '') AS UNSIGNED) ASC
        LIMIT 1
    `
	name = name + "%"
	row := db.QueryRow(query, name)

	var kwtStr string
	err = row.Scan(&kwtStr)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "No matching records found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	kwt, err := parseKwt(kwtStr)
	if err != nil {
		http.Error(w, "Error converting kwt to float64", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(KwtResponse{Kwt: kwt})
}

func getToasterByModel(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	db, err := sql.Open("mysql", "root:zikRerSPppEEPJZUeawwtpMpyCmpOmtK@tcp(monorail.proxy.rlwy.net:22986)/railway")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	query := `
        SELECT kwt 
        FROM Toaster
        WHERE name LIKE ?
        ORDER BY CAST(REPLACE(kwt, ' кВтч', '') AS UNSIGNED) ASC
        LIMIT 1
    `
	name = name + "%"
	row := db.QueryRow(query, name)

	var kwtStr string
	err = row.Scan(&kwtStr)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "No matching records found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	kwt, err := parseKwt(kwtStr)
	if err != nil {
		http.Error(w, "Error converting kwt to float64", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(KwtResponse{Kwt: kwt})
}

func getFreezerByModel(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	db, err := sql.Open("mysql", "root:zikRerSPppEEPJZUeawwtpMpyCmpOmtK@tcp(monorail.proxy.rlwy.net:22986)/railway")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	query := `
        SELECT kwt 
        FROM Freezer
        WHERE name LIKE ?
        ORDER BY CAST(REPLACE(kwt, ' кВтч', '') AS UNSIGNED) ASC
        LIMIT 1
    `
	name = name + "%"
	row := db.QueryRow(query, name)

	var kwtStr string
	err = row.Scan(&kwtStr)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "No matching records found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	kwt, err := parseKwt(kwtStr)
	if err != nil {
		http.Error(w, "Error converting kwt to float64", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(KwtResponse{Kwt: kwt})
}

func getCoffeeByModel(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	db, err := sql.Open("mysql", "root:zikRerSPppEEPJZUeawwtpMpyCmpOmtK@tcp(monorail.proxy.rlwy.net:22986)/railway")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	query := `
        SELECT kwt 
        FROM Coffee
        WHERE name LIKE ?
        ORDER BY CAST(REPLACE(kwt, ' Вт', '') AS UNSIGNED) ASC
        LIMIT 1
    `
	name = name + "%"
	row := db.QueryRow(query, name)

	var kwtStr string
	err = row.Scan(&kwtStr)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "No matching records found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	kwt, err := parseKwt(kwtStr)
	if err != nil {
		http.Error(w, "Error converting kwt to float64", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(KwtResponse{Kwt: kwt})
}

func getFridgeByModel(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	db, err := sql.Open("mysql", "root:zikRerSPppEEPJZUeawwtpMpyCmpOmtK@tcp(monorail.proxy.rlwy.net:22986)/railway")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	query := `
        SELECT kwt 
        FROM Fridge
        WHERE name LIKE ?
        ORDER BY CAST(REPLACE(kwt, ' кВтч', '') AS UNSIGNED) ASC
        LIMIT 1
    `
	name = name + "%"
	row := db.QueryRow(query, name)

	var kwtStr string
	err = row.Scan(&kwtStr)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "No matching records found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	kwt, err := parseKwt(kwtStr)
	if err != nil {
		http.Error(w, "Error converting kwt to float64", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(KwtResponse{Kwt: kwt})
}

// Helper function to parse kwt string to float64
func parseKwt(kwtStr string) (float64, error) {
	// Using regular expression to extract float value from string
	re := regexp.MustCompile(`[0-9]+\.?[0-9]*`)
	kwtMatches := re.FindStringSubmatch(kwtStr)

	if len(kwtMatches) < 1 {
		return 0, fmt.Errorf("Failed to parse kwt value")
	}

	// Converting extracted string to float64
	kwt, err := strconv.ParseFloat(kwtMatches[0], 64)
	if err != nil {
		return 0, err
	}

	return kwt, nil
}

func getCoffee(w http.ResponseWriter, r *http.Request) {
	// Sample Coffee data
	coffee := Coffee{Name: "Espresso Machine", Kwt: 1200}

	// Setting response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Encoding data to JSON and writing it to response
	json.NewEncoder(w).Encode(coffee)
}

func getFridge(w http.ResponseWriter, r *http.Request) {
	// Sample Fridge data
	fridge := Coffee{Name: "Smart Fridge", Kwt: 250}

	// Setting response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Encoding data to JSON and writing it to response
	json.NewEncoder(w).Encode(fridge)
}

func getFreezer(w http.ResponseWriter, r *http.Request) {
	// Sample Freezer data
	freezer := Coffee{Name: "Deep Freezer", Kwt: 500}

	// Setting response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Encoding data to JSON and writing it to response
	json.NewEncoder(w).Encode(freezer)
}

func getPlate(w http.ResponseWriter, r *http.Request) {
	// Sample Plate data
	plate := Coffee{Name: "Induction Cooktop", Kwt: 1500}

	// Setting response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Encoding data to JSON and writing it to response
	json.NewEncoder(w).Encode(plate)
}

func getWasher(w http.ResponseWriter, r *http.Request) {
	// Sample Washing Machine data
	washer := Coffee{Name: "Front Load Washer", Kwt: 800}

	// Setting response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Encoding data to JSON and writing it to response
	json.NewEncoder(w).Encode(washer)
}

func getToaster(w http.ResponseWriter, r *http.Request) {
	// Sample Toaster data
	toaster := Coffee{Name: "Pop-up Toaster", Kwt: 900}

	// Setting response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Encoding data to JSON and writing
	json.NewEncoder(w).Encode(toaster)
}

func getTechnic(w http.ResponseWriter, r *http.Request) {
	// Sample Technic data
	technic := []Coffee{
		{Name: "Espresso Machine", Kwt: 1200},
		{Name: "Smart Fridge", Kwt: 250},
		{Name: "Deep Freezer", Kwt: 500},
		{Name: "Induction Cooktop", Kwt: 1500},
		{Name: "Front Load Washer", Kwt: 800},
		{Name: "Pop-up Toaster", Kwt: 900},
	}

	// Setting response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Encoding data to JSON and writing it to response
	json.NewEncoder(w).Encode(technic)
}

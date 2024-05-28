package main

import (
	"database/sql"
	"encoding/json"
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

func main() {
	http.HandleFunc("/get-coffee-machine", getCoffee)
	http.HandleFunc("/get-fridge", getFridge)
	http.HandleFunc("/get-freezer", getFreezer)
	http.HandleFunc("/get-plate", getPlate)
	http.HandleFunc("/get-washer", getWasher)
	http.HandleFunc("/get-toaster", getToaster)

	// BY MODEL

	http.HandleFunc("/get-coffee-machine-by-model", getCoffeeByModel)
	http.HandleFunc("/get-fridge-by-model", getFridgeByModel)
	http.HandleFunc("/get-freezer-by-model", getFreezerByModel)
	http.HandleFunc("/get-plate-by-model", getPlateByModel)
	http.HandleFunc("/get-washer-by-model", getWasherByModel)
	http.HandleFunc("/get-toaster-by-model", getToasterByModel)

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
        SELECT name, kwt 
        FROM plita
        WHERE name LIKE ?
        ORDER BY CAST(REPLACE(kwt, ' кВт', '') AS UNSIGNED) ASC
        LIMIT 1
    `

	// Adjusting the search term to include the wildcard character
	name = name + "%"

	rows, err := db.Query(query, name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var coffees []Coffee
	for rows.Next() {
		var name string
		var kwtStr string
		err := rows.Scan(&name, &kwtStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		kwt, err := parseKwt(kwtStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		coffees = append(coffees, Coffee{Name: name, Kwt: kwt})
	}

	if err = rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(coffees)
}
func getWasherByModel(w http.ResponseWriter, r *http.Request) {
	// Extracting the name query parameter from the request
	name := r.URL.Query().Get("name")

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/diploma")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// SQL query with WHERE clause using LIKE for partial matching
	query := `
        SELECT name, kwt 
        FROM stiralka
        WHERE name LIKE ?
        ORDER BY CAST(REPLACE(kwt, ' кВт/ч', '') AS UNSIGNED) ASC
        LIMIT 1
    `

	// Adjusting the search term to include the wildcard character
	name = name + "%"

	rows, err := db.Query(query, name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var coffees []Coffee
	for rows.Next() {
		var name string
		var kwtStr string
		err := rows.Scan(&name, &kwtStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		kwt, err := parseKwt(kwtStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		coffees = append(coffees, Coffee{Name: name, Kwt: kwt})
	}

	if err = rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(coffees)
}
func getToasterByModel(w http.ResponseWriter, r *http.Request) {
	// Extracting the name query parameter from the request
	name := r.URL.Query().Get("name")

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/diploma")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// SQL query with WHERE clause using LIKE for partial matching
	query := `
        SELECT name, kwt 
        FROM toster
        WHERE name LIKE ?
        ORDER BY CAST(REPLACE(kwt, ' кВтч', '') AS UNSIGNED) ASC
        LIMIT 1
    `

	// Adjusting the search term to include the wildcard character
	name = name + "%"

	rows, err := db.Query(query, name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var coffees []Coffee
	for rows.Next() {
		var name string
		var kwtStr string
		err := rows.Scan(&name, &kwtStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		kwt, err := parseKwt(kwtStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		coffees = append(coffees, Coffee{Name: name, Kwt: kwt})
	}

	if err = rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(coffees)
}

func getFreezerByModel(w http.ResponseWriter, r *http.Request) {
	// Extracting the name query parameter from the request
	name := r.URL.Query().Get("name")

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/diploma")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// SQL query with WHERE clause using LIKE for partial matching
	query := `
        SELECT name, kwt 
        FROM moroz
        WHERE name LIKE ?
        ORDER BY CAST(REPLACE(kwt, ' кВтч', '') AS UNSIGNED) ASC
        LIMIT 1
    `

	// Adjusting the search term to include the wildcard character
	name = name + "%"

	rows, err := db.Query(query, name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var coffees []Coffee
	for rows.Next() {
		var name string
		var kwtStr string
		err := rows.Scan(&name, &kwtStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		kwt, err := parseKwt(kwtStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		coffees = append(coffees, Coffee{Name: name, Kwt: kwt})
	}

	if err = rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(coffees)
}

func getCoffeeByModel(w http.ResponseWriter, r *http.Request) {
	// Extracting the name query parameter from the request
	name := r.URL.Query().Get("name")

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/diploma")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// SQL query with WHERE clause using LIKE for partial matching
	query := `
        SELECT name, kwt 
        FROM coffee
        WHERE name LIKE ?
        ORDER BY CAST(REPLACE(kwt, ' Вт', '') AS UNSIGNED) ASC
        LIMIT 1
    `

	// Adjusting the search term to include the wildcard character
	name = name + "%"

	rows, err := db.Query(query, name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var coffees []Coffee
	for rows.Next() {
		var name string
		var kwtStr string
		err := rows.Scan(&name, &kwtStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		kwt, err := parseKwt(kwtStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		coffees = append(coffees, Coffee{Name: name, Kwt: kwt})
	}

	if err = rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(coffees)
}

func getFridgeByModel(w http.ResponseWriter, r *http.Request) {
	// Extracting the name query parameter from the request
	name := r.URL.Query().Get("name")

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/diploma")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// SQL query with WHERE clause using LIKE for partial matching
	query := `
        SELECT name, kwt 
        FROM holod
        WHERE name LIKE ?
        ORDER BY CAST(REPLACE(kwt, ' кВтч', '') AS UNSIGNED) ASC
        LIMIT 1
    `

	// Adjusting the search term to include the wildcard character
	name = name + "%"

	rows, err := db.Query(query, name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var coffees []Coffee
	for rows.Next() {
		var name string
		var kwtStr string
		err := rows.Scan(&name, &kwtStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		kwt, err := parseKwt(kwtStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		coffees = append(coffees, Coffee{Name: name, Kwt: kwt})
	}

	if err = rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(coffees)
}

// ALL

func parseKwt(kwtStr string) (float64, error) {
	re := regexp.MustCompile(`[^\d.]`)
	cleanedStr := re.ReplaceAllString(kwtStr, "")
	if cleanedStr == "" {
		return 0, nil
	}
	return strconv.ParseFloat(cleanedStr, 64)
}
func getCoffee(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/diploma")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, kwt FROM coffee ORDER BY kwt ASC")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var coffees []Coffee
	for rows.Next() {
		var name string
		var kwtStr string
		err := rows.Scan(&name, &kwtStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		kwt, err := parseKwt(kwtStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		coffees = append(coffees, Coffee{Name: name, Kwt: kwt})
	}

	if err = rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(coffees)
}

func getFridge(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/diploma")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, kwt FROM holod ORDER BY kwt ASC")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var fridges []Coffee
	for rows.Next() {
		var name string
		var kwtStr string
		err := rows.Scan(&name, &kwtStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		kwt, err := parseKwt(kwtStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fridges = append(fridges, Coffee{Name: name, Kwt: kwt})
	}

	if err = rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fridges)
}

func getFreezer(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/diploma")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, kwt FROM holod ORDER BY CAST(REPLACE(kwt, ' Вт', '') AS UNSIGNED) ASC")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var freezers []Coffee
	for rows.Next() {
		var name string
		var kwtStr string
		err := rows.Scan(&name, &kwtStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		kwt, err := parseKwt(kwtStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		freezers = append(freezers, Coffee{Name: name, Kwt: kwt})
	}

	if err = rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(freezers)
}

func getPlate(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/diploma")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, kwt FROM plita ORDER BY CAST(REPLACE(kwt, ' Вт', '') AS UNSIGNED) ASC")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var plates []Coffee
	for rows.Next() {
		var name string
		var kwtStr string
		err := rows.Scan(&name, &kwtStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		kwt, err := parseKwt(kwtStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		plates = append(plates, Coffee{Name: name, Kwt: kwt})
	}

	if err = rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(plates)
}

func getWasher(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/diploma")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Query to select name and kwt, and order by numeric value of kwt
	rows, err := db.Query(`
		SELECT name, kwt 
		FROM stiralka 
		ORDER BY CAST(REPLACE(kwt, ' Вт', '') AS UNSIGNED) ASC
	`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var washers []Coffee
	for rows.Next() {
		var name string
		var kwtStr string
		err := rows.Scan(&name, &kwtStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		kwt, err := parseKwt(kwtStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		washers = append(washers, Coffee{Name: name, Kwt: kwt})
	}

	if err = rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(washers)
}

func getToaster(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/diploma")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, kwt FROM toster ORDER BY CAST(REPLACE(kwt, ' Вт', '') AS UNSIGNED) ASC")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var toasters []Coffee
	for rows.Next() {
		var name string
		var kwtStr string
		err := rows.Scan(&name, &kwtStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		kwt, err := parseKwt(kwtStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		toasters = append(toasters, Coffee{Name: name, Kwt: kwt})
	}

	if err = rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(toasters)
}

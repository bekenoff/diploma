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
        SELECT name, kwt 
        FROM Plate
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

	db, err := sql.Open("mysql", "root:zikRerSPppEEPJZUeawwtpMpyCmpOmtK@tcp(monorail.proxy.rlwy.net:22986)/railway")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// SQL query with WHERE clause using LIKE for partial matching
	query := `
        SELECT name, kwt 
        FROM Washing Machine
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

	db, err := sql.Open("mysql", "root:zikRerSPppEEPJZUeawwtpMpyCmpOmtK@tcp(monorail.proxy.rlwy.net:22986)/railway")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// SQL query with WHERE clause using LIKE for partial matching
	query := `
        SELECT name, kwt 
        FROM Toaster
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

	db, err := sql.Open("mysql", "root:zikRerSPppEEPJZUeawwtpMpyCmpOmtK@tcp(monorail.proxy.rlwy.net:22986)/railway")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// SQL query with WHERE clause using LIKE for partial matching
	query := `
        SELECT name, kwt 
        FROM Freezer
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

	db, err := sql.Open("mysql", "root:zikRerSPppEEPJZUeawwtpMpyCmpOmtK@tcp(monorail.proxy.rlwy.net:22986)/railway")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// SQL query with WHERE clause using LIKE for partial matching
	query := `
        SELECT name, kwt 
        FROM Coffee
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

	db, err := sql.Open("mysql", "root:zikRerSPppEEPJZUeawwtpMpyCmpOmtK@tcp(monorail.proxy.rlwy.net:22986)/railway")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// SQL query with WHERE clause using LIKE for partial matching
	query := `
        SELECT name, kwt 
        FROM Freezer
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
	db, err := sql.Open("mysql", "root:zikRerSPppEEPJZUeawwtpMpyCmpOmtK@tcp(monorail.proxy.rlwy.net:22986)/railway")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, kwt FROM Coffee ORDER BY kwt ASC")
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
	db, err := sql.Open("mysql", "root:zikRerSPppEEPJZUeawwtpMpyCmpOmtK@tcp(monorail.proxy.rlwy.net:22986)/railway")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, kwt FROM Freezer ORDER BY kwt ASC")
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
	db, err := sql.Open("mysql", "root:zikRerSPppEEPJZUeawwtpMpyCmpOmtK@tcp(monorail.proxy.rlwy.net:22986)/railway")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, kwt FROM Freezer ORDER BY CAST(REPLACE(kwt, ' Вт', '') AS UNSIGNED) ASC")
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
	db, err := sql.Open("mysql", "root:zikRerSPppEEPJZUeawwtpMpyCmpOmtK@tcp(monorail.proxy.rlwy.net:22986)/railway")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, kwt FROM Plate ORDER BY CAST(REPLACE(kwt, ' Вт', '') AS UNSIGNED) ASC")
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
	db, err := sql.Open("mysql", "root:zikRerSPppEEPJZUeawwtpMpyCmpOmtK@tcp(monorail.proxy.rlwy.net:22986)/railway")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Query to select name and kwt, and order by numeric value of kwt
	rows, err := db.Query(`
		SELECT name, kwt 
		FROM Washing Machine 
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
	db, err := sql.Open("mysql", "root:zikRerSPppEEPJZUeawwtpMpyCmpOmtK@tcp(monorail.proxy.rlwy.net:22986)/railway")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT name, kwt FROM Toaster ORDER BY CAST(REPLACE(kwt, ' Вт', '') AS UNSIGNED) ASC")
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

func getTechnic(w http.ResponseWriter, r *http.Request) {
	// Парсинг параметров запроса
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Получение имени таблицы и модели техники из параметров запроса
	tableName := r.Form.Get("table")
	name := r.Form.Get("name")

	// Проверка наличия обязательных параметров
	if tableName == "" || name == "" {
		http.Error(w, "Необходимо указать название таблицы и неполное название модели техники", http.StatusBadRequest)
		return
	}

	// Открытие соединения с базой данных
	db, err := sql.Open("mysql", "root:zikRerSPppEEPJZUeawwtpMpyCmpOmtK@tcp(monorail.proxy.rlwy.net:22986)/railway")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Подготовка SQL запроса с плейсхолдером и оператором LIKE
	query := "SELECT name, kwt FROM " + tableName + " WHERE name LIKE ? ORDER BY kwt ASC"
	rows, err := db.Query(query, "%"+name+"%")
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

	// Установка заголовка Content-Type и кодирование данных в JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fridges)
}

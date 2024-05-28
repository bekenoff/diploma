# diploma
# def
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

package main

import "C"
import (
	"strings"
)

//export HandleRVInput
func HandleRVInput(input *C.char) string {
	var s string = C.GoString(input)

	function, params := parseFunctionCall(s)

	response := ""

	switch function {
	case "version":
		response = version()
	case "prepLoadout":
		response = prepLoadout(params)
	case "getStatus":
		response = getStatus()
	case "getPrimaryWeapon":
		response = getPrimaryWeapon()
	}

	return response
}

func parseFunctionCall(input string) (string, string) {

	i := strings.Index(input, " ")

	if i == -1 {
		return input, ""
	} else {
		return input[:i], input[i+1:]
	}

}

func version() string {
	return "1.0.0"
}

func getStatus() string {
	if IsLoading.IsSet(){
		return "loading"
	} else if IsErrored.IsSet() {
		return "error"
	} else {
		return "ready"
	}
}

func getPrimaryWeapon() string {
	return ActiveLoadout.PrimaryWeaponSystems[0]
}

func prepLoadout(steamid string) string {

	SyncLoadout(steamid)

	return "Initialized loading of steamid: " + steamid

}

func main() {
	// This function is required to export a C Binary but will never be used
}
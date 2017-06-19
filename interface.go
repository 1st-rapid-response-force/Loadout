package main

import "C"
import (
	"strings"
)

//export HandleRVInput
func HandleRVInput(input *C.char) *C.char {
	var s string = C.GoString(input)

	function, params := parseFunctionCall(s)

	response := runRVTask(function, params)
	return C.CString(response)
}

func runRVTask(function string, params string) string {
	defer func() {
		if r := recover(); r != nil {
			return
		}
	}()

	switch function {
	case "version":
		return version()
	case "prepLoadout":
		return prepLoadout(params)
	case "getStatus":
		return getStatus()
	case "getPrimaryWeapon":
		return getPrimaryWeapon()
	}

	return "Function not Found"
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
	return "1.0.11"
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
	if len(ActiveLoadout.PrimaryWeaponSystems) > 0 {
		return ActiveLoadout.PrimaryWeaponSystems[0]
	} else {
		return ""
	}
}

func prepLoadout(steamid string) string {

	SyncLoadout(steamid)


	return "loading"

}

func main() {
	// This function is required to export a C Binary but will never be used
}
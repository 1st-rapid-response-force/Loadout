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
	case "getPrimaryAttachments":
		return getPrimaryWeaponAttachments()
	case "getSecondaryWeapon":
		return getSecondaryWeapon()
	case "getSecondaryAttachments":
		return getSecondaryAttachments()
	case "getLauncher":
		return getLauncherWeapon()
	case "getHelmet":
		return getHelmets()
	case "getUniform":
		return getUniform()
	case "getVest":
		return getVest()
	case "getBackpack":
		return getBackpack()
	case "getNightvision":
		return getNightvision()
	case "getGoggles":
		return getGoggles()
	case "getAdditionalEquipment":
		return getAdditionalEquipment()
	case "getRole":
		return getRole()
	case "getSpecialisms":
		return getSpecialisms()
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
	if IsLoading.IsSet() {
		return "loading"
	} else if IsErrored.IsSet() {
		return "error"
	} else {
		return "ready"
	}
}

func main() {
	// This function is required to export a C Binary but will never be used
}

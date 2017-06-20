package main

func serializeArrayToRV(items []string) string {
	response := ""

	for i := 0; i < len(items); i++ {
		response += items[i]
		response += " "
	}

	return response
}

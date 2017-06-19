package main

import (
	"net/http"
	"encoding/json"
)

type RRFLoadoutEndpoint struct {
	Position string `json:"member_position"`
	Loadout []struct {
		Category string `json:"category"`
		ClassName string `json:"class_name"`
	} `json:"loadout"`
}

// Synchronously set the loading flag
func SyncLoadout(steamid string) {

	if IsLoading.IsSet() {
		return
	}

	IsLoading.Set()
	IsErrored.UnSet()

	InitLoadout()

	go makeRequest(steamid)

}

func makeRequest(steamid string) {
	defer IsLoading.UnSet()

	resp, err := http.Get("https://1st-rrf.com/api/loadout/" + steamid)

	if err != nil || resp.StatusCode != http.StatusOK {
		IsErrored.Set()
		IsLoading.UnSet()
		return
	}

	defer resp.Body.Close()

	// Parse the JSON response to fill out the loadout section
	parsedResponse := RRFLoadoutEndpoint{}
	json.NewDecoder(resp.Body).Decode(&parsedResponse)

	parsedLoadout := Loadout{}

	for i := 0; i < len(parsedResponse.Loadout); i++ {

		item := parsedResponse.Loadout[i]

		switch item.Category {

		case "primary":
			parsedLoadout.PrimaryWeaponSystems = append(parsedLoadout.PrimaryWeaponSystems, item.ClassName)

		case "items":
			parsedLoadout.AdditionalEquipment = append(parsedLoadout.AdditionalEquipment, item.ClassName)

		}

	}

	ActiveLoadout = parsedLoadout

}
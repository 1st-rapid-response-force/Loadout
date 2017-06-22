package main

import (
	"encoding/json"
	"net/http"
)

type RRFLoadoutEndpoint struct {
	Position string `json:"member_position"`
	Loadout  []struct {
		Category  string `json:"category"`
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
		case "secondary":
			parsedLoadout.SidearmType = item.ClassName
		case "launcher":
			parsedLoadout.LauncherType = item.ClassName
		case "thrown":
			parsedLoadout.AdditionalEquipment = append(parsedLoadout.AdditionalEquipment, item.ClassName)
		case "uniform":
			parsedLoadout.Uniform = item.ClassName
		case "vest":
			parsedLoadout.Vest = item.ClassName
		case "backpack":
			parsedLoadout.Backpack = item.ClassName
		case "helmet":
			parsedLoadout.Helmets = append(parsedLoadout.Helmets, item.ClassName)
		case "goggles":
			parsedLoadout.Goggles = item.ClassName
		case "nightvision":
			parsedLoadout.Nightvision = item.ClassName
		case "primary_attachments":
			parsedLoadout.PrimaryWeaponAttachments = append(parsedLoadout.PrimaryWeaponAttachments, item.ClassName)
		case "secondary_attachments":
			parsedLoadout.SidearmAttachments = append(parsedLoadout.SidearmAttachments, item.ClassName)
		case "items":
			parsedLoadout.AdditionalEquipment = append(parsedLoadout.AdditionalEquipment, item.ClassName)
		}

	}

	parsedLoadout.Role = parsedResponse.Position

	ActiveLoadout = parsedLoadout

}

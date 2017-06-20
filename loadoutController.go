package main

import (
	"github.com/tevino/abool"
)

type Loadout struct {
	PrimaryWeaponSystems     []string
	PrimaryWeaponAttachments []string
	SidearmType              string
	SidearmAttachments       []string
	LauncherType             string
	Helmets                  []string
	Uniform                  string
	Vest                     string
	Backpack                 string
	Nightvision              string
	Goggles                  string
	AdditionalEquipment      []string
	Role                     string
	IsMedicQualified         bool
	IsEODQualified           bool
	IsEngineerQualified      bool
}

// There is a clusterfuck of thread safety here. Rather than flood this with mutex locks I have decided to handle it on
// the SQF side. No client should concurrently attempt a read that would compromise thread safety based on the SQF
// logic so this can be sloppy.

var ActiveLoadout Loadout
var IsLoading *abool.AtomicBool = abool.New()
var IsErrored *abool.AtomicBool = abool.New()

func InitLoadout() {
	ActiveLoadout = Loadout{}
}

func prepLoadout(steamid string) string {
	SyncLoadout(steamid)

	return "loading"
}

func getPrimaryWeapon() string {
	return serializeArrayToRV(ActiveLoadout.PrimaryWeaponSystems)
}

func getPrimaryWeaponAttachments() string {
	return serializeArrayToRV(ActiveLoadout.PrimaryWeaponAttachments)
}

func getSecondaryWeapon() string {
	return ActiveLoadout.SidearmType
}

func getSecondaryAttachments() string {
	return serializeArrayToRV(ActiveLoadout.SidearmAttachments)
}

func getLauncherWeapon() string {
	return ActiveLoadout.LauncherType
}

func getHelmets() string {
	return serializeArrayToRV(ActiveLoadout.Helmets)
}

func getUniform() string {
	return ActiveLoadout.Uniform
}

func getVest() string {
	return ActiveLoadout.Vest
}

func getBackpack() string {
	return ActiveLoadout.Backpack
}

func getNightvision() string {
	return ActiveLoadout.Nightvision
}

func getGoggles() string {
	return ActiveLoadout.Goggles
}

func getAdditionalEquipment() string {
	return serializeArrayToRV(ActiveLoadout.AdditionalEquipment)
}

func getRole() string {
	return ActiveLoadout.Role
}

func getSpecialisms() string {
	response := ""

	if ActiveLoadout.IsMedicQualified {
		response += "medic "
	}

	if ActiveLoadout.IsEODQualified {
		response += "eod "
	}

	if ActiveLoadout.IsEngineerQualified {
		response += "engineer "
	}

	return response
}

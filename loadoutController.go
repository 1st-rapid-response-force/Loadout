package main

import (
	"github.com/tevino/abool"
)

type Loadout struct {
	PrimaryWeaponSystems []string
	PrimaryWeaponAttachments []string
	SidearmType string
	SidearmAttachments []string
	LauncherType string
	Helmets []string
	Uniform string
	Vest string
	Backpack string
	Nightvision string
	Goggles string
	AdditionalEquipment []string
	BaseRole string
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
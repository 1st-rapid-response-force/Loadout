package main

import (
	"testing"
	"fmt"
)

func TestFunctionCall(t *testing.T) {

	c := []struct{
		in string
		function string
		params string
	}{
		{"version", "version", ""},
		{"getStatus", "getStatus", ""},
		{"prepLoadout steamid", "prepLoadout", "steamid"},
	}

	for i := 0; i < len(c); i++  {

		function, param := parseFunctionCall(c[i].in)

		if function != c[i].function || param != c[i].params {
			t.Error("Function call test failed")
		}

	}
}

func TestVersionExecution(t *testing.T) {

	fmt.Println(runRVTask("version", ""))

}

func TestPrimaryWeaponRetrievalExecution(t *testing.T) {

	fmt.Println(runRVTask("getPrimaryWeapon", ""))

}

func TestSyncLoadoutExecution(t *testing.T) {

	fmt.Println(runRVTask("prepLoadout", "76561198021531457"))

	fmt.Println(IsLoading.IsSet())
	for runRVTask("getStatus", "") != "ready" {

	};
	fmt.Println(runRVTask("getPrimaryWeapon", ""))

	fmt.Println(runRVTask("prepLoadout", "bullshit"))

	for runRVTask("getStatus", "") == "loading" {

	};
	fmt.Println(runRVTask("getPrimaryWeapon", ""))
}
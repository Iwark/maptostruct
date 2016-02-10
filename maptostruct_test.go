package maptostruct

import "testing"

type Person struct {
	Name   string
	Age    int
	Gender string `mts:"gender"`
}

func TestDo(t *testing.T) {
	person := map[string]interface{}{
		"Name":   "Iwark",
		"Age":    24,
		"gender": "man",
		"hobby":  "playing the piano",
	}

	result := &Person{}
	err := Do(person, result)
	if err != nil {
		t.Error("Convert Error:", err)
	}
	if result.Name != "Iwark" || result.Age != 24 {
		t.Error("Map to struct failed, got:", result)
	}
	if result.Gender != "man" {
		t.Error("Could not handle mts tag:", result)
	}
}

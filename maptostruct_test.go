package maptostruct

import "testing"

type MyStruct struct {
	Name string
	Age  int64
}

func TestDo(t *testing.T) {
	myData := make(map[string]interface{})
	myData["Name"] = "Tony"
	myData["Age"] = int64(23)
	result := &MyStruct{}
	err := Do(myData, result)
	if err != nil {
		t.Error("Convert Error:", err)
	}
	if result.Name != "Tony" || result.Age != int64(23) {
		t.Error("Map to struct failed, got:", result)
	}
}

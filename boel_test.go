package gojs_test

import (
	"github.com/simfatic/gojs"
	"io/ioutil"
	"testing"
)

type ValidationError struct {
	Validation string `json:"validation"`
	Message    string `json:"message"`
}
type ValidationResult struct {
	HasErrors bool                       `json:"has_errors"`
	ErrorMap  map[string]ValidationError `json:"error_map"`
}

//Test loading the boel JS library (https://github.com/simfatic/boel) and passing data to validate
//The library is budled in iife format, minimized
func RunJSFormDataTest(settings string, data string) (*ValidationResult, error) {
	bc, err := ioutil.ReadFile("./boel.min.js")
	if err != nil {
		return nil, err
	}
	code := string(bc)

	custom := `
	var bv = boel.makeBoel()
	function validate(strSettings, strFormData)
	{
		var settings = JSON.parse(strSettings)
		var fd = JSON.parse(strFormData)
		return bv.validateFields(settings.fields, fd)
	} 
	`
	js, err := gojs.New(code + custom)
	if err != nil {
		return nil, err
	}

	_, err = js.Run()
	if err != nil {
		return nil, err
	}

	res, err := js.GetGlobalObject().Call("validate", settings, data)

	if err != nil {
		return nil, err
	}
	var resX ValidationResult

	err = js.VM().ExportTo(res, &resX)
	if err != nil {
		return nil, err
	}

	return &resX, nil

}
func TestSimpleBoel(t *testing.T) {
	settings := `{"fields":[{"name":"name","type":"text","validations":[{"_vtype":"Required","condition":"","enabled":true,"message":""}]},{"name":"email","type":"text","validations":[{"_vtype":"Required","condition":"","enabled":true,"message":""}]},{"name":"Age","type":"number","validations":[{"_vtype":"GreaterThan","condition":"","message":"","num":"15"},{"_vtype":"LessThan","condition":"","message":"","num":"60"}]}]}`
	resX, err := RunJSFormDataTest(settings, "{}")
	if err != nil {
		t.Fatalf("Error running JS code %v", err)
	}
	t.Logf("Received result %v", resX)

	if resX.HasErrors != true {
		t.Fatalf("Expected the empty data return errors ")
	}

	for f, e := range resX.ErrorMap {
		t.Logf("Field %s Error %s ", f, e.Message)
	}

	if resX.ErrorMap["name"].Message != "name is required" {
		t.Error("The name field required error is not present in the result")
	}

	if resX.ErrorMap["email"].Message != "email is required" {
		t.Error("The email field required error is not present in the result")
	}
}

func TestBoelValidData(t *testing.T) {

	settings := `{"fields":[{"name":"name","type":"text","validations":[{"_vtype":"Required","condition":"","enabled":true,"message":""}]},{"name":"email","type":"text","validations":[{"_vtype":"Required","condition":"","enabled":true,"message":""}]},{"name":"Age","type":"number","validations":[{"_vtype":"GreaterThan","condition":"","message":"","num":"15"},{"_vtype":"LessThan","condition":"","message":"","num":"60"}]}]}`
	data := `{"name":"Robert", "email":"robert@tdd.com", "Age":18 }`
	resX, err := RunJSFormDataTest(settings, data)
	if err != nil {
		t.Fatalf("Error running JS code %v", err)
	}
	t.Logf("Received result %v", resX)

	if resX.HasErrors != false {
		t.Error("Expected the data to validate to true. ")
	}
}

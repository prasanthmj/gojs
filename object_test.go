package gojs_test

import (
	"github.com/prasanthmj/gojs"
	"testing"
)

func TestCallingObjectMethod(t *testing.T) {
	code := `
	function someObj()
	{
		this.greet=function(vv)
		{
			return "hello "+vv;
		}
		return this;
	}
	var someobj = new someObj()
	`

	js, err := gojs.New(code)
	if err != nil {
		t.Fatalf("Error loading JS code %v", err)
	}
	_, err = js.Run()
	if err != nil {
		t.Fatalf("Error running JS code %v", err)
	}

	obj, err := js.GetObject("someobj")
	if err != nil {
		t.Fatalf("Error getting object. %v", err)
	}

	res, err := obj.Call("greet", "me")
	if err != nil {
		t.Fatalf("Error calling object method %v", err)
	}
	restr := res.String()
	if restr != "hello me" {
		t.Fatalf(`Expected "Hello me" received "%s"`, restr)
	}
	t.Logf("Received method result %s ", restr)
}

func TestGettingNestedObject(t *testing.T) {
	code := `
	var someobj = {
		anotherobj:{
			v1:"value1",
			num:123
		},
		v1:"mvalue1"
	}
	`

	js, err := gojs.New(code)
	if err != nil {
		t.Fatalf("Error loading JS code %v", err)
		return
	}
	_, err = js.Run()
	if err != nil {
		t.Fatalf("Error running JS code %v", err)
	}
	sobj, err := js.GetObject("someobj")
	if err != nil {
		t.Fatalf("Error getting object. %v", err)
	}

	sv1, err := sobj.GetString("v1")
	if err != nil {
		t.Fatalf("Error calling object method %v", err)
	}
	t.Logf("sobj v1 is %s", sv1)

	aobj, err := sobj.GetObject("anotherobj")
	if err != nil {
		t.Fatalf("Error calling object method %v", err)
	}

	sval, err := aobj.GetString("v1")
	if err != nil {
		t.Fatalf("Error getting prop %v", err)
	}

	t.Logf("received str value %s", sval)
	if sval != "value1" {
		t.Fatalf("Expected value value1 got %s", sval)
	}

	nval, err := aobj.GetNumber("num")
	if err != nil {
		t.Fatalf("Error getting prop num. %v", err)
	}
	t.Logf("received num value %d", nval)

	if nval != 123 {
		t.Fatalf("Expected value 123 got %d", nval)
	}
}

func TestCallingMethodReturningObject(t *testing.T) {
	code := `
	function AnotherObj()
	{
		this.v1 ="myvalue1";
		this.greet = function(who)
		{
			return "hello "+who;
		}
		return this;
	}
	var someobj = {
		getObj:function()
		{
			return new AnotherObj()
		}
	}
	`

	js, err := gojs.New(code)
	if err != nil {
		t.Fatalf("Error loading JS code %v", err)
		return
	}
	_, err = js.Run()
	if err != nil {
		t.Fatalf("Error running JS code %v", err)
	}
	sobj, err := js.GetObject("someobj")
	if err != nil {
		t.Fatalf("Error getting object. %v", err)
	}

	aobj, err := sobj.CallReturningObj("getObj")
	if err != nil {
		t.Fatalf("Error getting object. %v", err)
	}
	sval, err := aobj.GetString("v1")
	if err != nil {
		t.Fatalf("Error getting object prop %v", err)
	}
	t.Logf("received str value %s", sval)
	if sval != "myvalue1" {
		t.Fatalf("Expected value value1 got %s", sval)
	}

	mval, err := aobj.CallReturningStr("greet", "me")
	if err != nil {
		t.Fatalf("Error calling object method %v", err)
	}
	t.Logf("received mval value %s", mval)

	if mval != "hello me" {
		t.Fatalf("Expected value hello me got %s", mval)
	}
}

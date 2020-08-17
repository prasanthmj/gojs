package gojs_test

import (
	"fmt"
	"github.com/prasanthmj/gojs"
	"testing"
)

func myGoApiFn() (string, error) {
	fmt.Printf("callback received from JS")
	return "result", nil
}

func TestInjectingFunction(t *testing.T) {
	code := `
	var vv="call me"
	var res = yourGoFunction(vv)
	`
	js, err := gojs.New(code)

	if err != nil {
		t.Fatalf("Error loading JS code %v", err)
		return
	}

	js.InjectFn("yourGoFunction", myGoApiFn)

	_, err = js.Run()
	if err != nil {
		t.Fatalf("Error running JS code %v", err)
	}
	res, err := js.GetGlobalObject().GetString("res")
	if err != nil {
		t.Fatalf("Error Getting global obj value %v", err)
	}

	t.Logf("value returned from JS %s", res)

}

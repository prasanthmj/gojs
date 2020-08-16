# GoJS
A wrapper around the goja library (Javascript interpreter in native Go)
This is a simple wrapper to make it easier to use the goja library.


## Usage
Suppose this is the Javascript Code you want to run:

```js
function someObj()
{
    this.greet=function(vv)
    {
        return "hello "+vv;
    }
    return this;
}
var someobj = new someObj()
```
Remember that the code should be es5 compatible  

In order to run the greet() method of the someObj object, these are the steps  

```go
//Load the code
js, err := gojs.New(code)

// execute the code so that the global "someobject"  is created
_, err = js.Run()

//Get the object
obj, err := js.GetObject("someobj")

//Now run the greet() function, passing a parameter
res, err := obj.Call("greet", "me")

//Extract string value from the result
restr := res.String()

//restr will be "hello me"

```
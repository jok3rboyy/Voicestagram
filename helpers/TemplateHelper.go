package helpers

import (
	"encoding/json"
	"html/template"
	"reflect"
)

type Template struct {
	*template.Template
}

func (t *Template) AddFuncs(funcs template.FuncMap) {
	t.Funcs(funcs)
}

func NewTemplate() (*Template, error) {
	templ, err := template.New("").Funcs(template.FuncMap{
		"parseJSON": parseJSON,
		"sub":       sub,
		"add":       add,
		"addFloat":  addFloat,
		"div":       div,
		"mult":      mult,
		"min":       min,
		"seq":       seq,
		"reverse":   reverse,
		"append":    append,
	}).ParseGlob("./templates/*")
	return &Template{
		Template: templ,
	}, err
}

func parseJSON(data string) (interface{}, error) {
	var result interface{}
	err := json.Unmarshal([]byte(data), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func call(fn interface{}, args ...interface{}) []reflect.Value {
	v := reflect.ValueOf(fn)
	if v.Kind() != reflect.Func {
		panic("Call: first argument is not a function")
	}
	input := make([]reflect.Value, len(args))
	for i, arg := range args {
		input[i] = reflect.ValueOf(arg)
	}
	return v.Call(input)
}

func sub(a, b int) int {
	return a - b
}

func add(a, b int) int {
	return a + b
}

func addFloat(a, b float64) float64 {
	return a + b
}

func div(a, b int) int {
	return a / b
}

func mult(a, b int) int {
	return a * b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func reverse(slice []interface{}) []interface{} {
	for i := len(slice)/2 - 1; i >= 0; i-- {
		opp := len(slice) - 1 - i
		slice[i], slice[opp] = slice[opp], slice[i]
	}
	return slice
}

func append(slice, elems interface{}) interface{} {
	s := reflect.ValueOf(slice)
	e := reflect.ValueOf(elems)
	if s.Kind() != reflect.Slice || e.Kind() != reflect.Slice {
		panic("Append: arguments must be slices")
	}
	return reflect.AppendSlice(s, e).Interface()
}

func seq(start, end int) []int {
	s := make([]int, end-start+1)
	for i := range s {
		s[i] = start + i
	}
	return s
}

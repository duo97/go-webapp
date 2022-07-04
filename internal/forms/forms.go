package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

//Form is a custom form struct
type Form struct {
	url.Values
	Errors errors
}

//New initializes a form struct
func New(data url.Values) *Form {

	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

//Has checks it a field in a form is filled
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		f.Errors.Add(field, "This field cannot be blank")
		return false
	}
	return true
}

//Valid checks if there is no error in the form
func (f *Form) Valid() bool {
	fmt.Printf("errors:%v", f.Errors)
	return len(f.Errors) == 0
}

//Required raise errors if sny required fileds are not filled
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			// fmt.Printf("added an error")
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

//to do : add email validator

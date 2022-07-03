package forms

import (
	"net/http"
	"net/url"
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
		return false
	}
	return true
}

//Valid checks if there is no error in the form
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

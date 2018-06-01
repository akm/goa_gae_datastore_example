// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "appengine": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/akm/goa_gae_datastore_example/design
// --out=$(GOPATH)/src/github.com/akm/goa_gae_datastore_example
// --version=v1.3.1

package app

import (
	"github.com/goadesign/goa"
)

// user (default view)
//
// Identifier: application/vnd.user+json; view=default
type User struct {
	// id(int64)
	ID interface{} `form:"id" json:"id" yaml:"id" xml:"id"`
	// id(string)
	IDStr string `form:"id_str" json:"id_str" yaml:"id_str" xml:"id_str"`
	// name
	Name string `form:"name" json:"name" yaml:"name" xml:"name"`
}

// Validate validates the User media type instance.
func (mt *User) Validate() (err error) {

	if mt.IDStr == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id_str"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	return
}

// UserCollection is the media type for an array of User (default view)
//
// Identifier: application/vnd.user+json; type=collection; view=default
type UserCollection []*User

// Validate validates the UserCollection media type instance.
func (mt UserCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"github.com/apito-cms/gqlgen/_examples/scalars/external"
)

type Address struct {
	ID       external.ObjectID `json:"id"`
	Location *Point            `json:"location,omitempty"`
}

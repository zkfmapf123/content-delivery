package mysql

import (
	"testing"
)

func Test_ping(t *testing.T) {
	
	_, err := CreateDBConnection().
		WithHost("43.201.213.254").
		WithPort("3305").
		WithUser("dobby").
		WithPassword("1234").
		WithDatabase("users").
		Build()
	
	if err != nil {
		panic(err)
	}
}

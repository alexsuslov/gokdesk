package model

import "fmt"

type Error map[string]interface{}

func (err Error) String() string {
	return fmt.Sprintf("%#v", err)
}

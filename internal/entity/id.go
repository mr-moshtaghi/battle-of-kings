package entity

import "strings"

type ID string

func (i ID) Type() string {
	return strings.Split(string(i), ":")[0]
}

func (i ID) ID() string {
	return strings.Split(string(i), ":")[1]
}

func (i ID) String() string {
	return string(i)
}

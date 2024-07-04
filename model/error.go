package model

import "fmt"

type ErrNotFound struct {
	ID int
}

func (e ErrNotFound) Error() string {
	return fmt.Sprintf("Error Not foud id : %d ", e.ID)
}

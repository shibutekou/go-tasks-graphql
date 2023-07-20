package entity

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"io"
	"strconv"
)

type Task struct {
	ID          int    `json:"id" db:"id"`
	Status      string `json:"status" db:"status"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
}

func MarshalID(id int) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Quote(fmt.Sprintf("%d", id)))
	})
}

func UnmarshalID(v interface{}) (int, error) {
	id, ok := v.(string)
	if !ok {
		return 0, fmt.Errorf("ids must be strings")
	}
	i, e := strconv.Atoi(id)
	return int(i), e
}

type TaskInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

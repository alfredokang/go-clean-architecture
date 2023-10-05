package students

import "github.com/google/uuid"

type InputBody struct {
	Name string `json:"name" validate:"required|string|min_len:3|max_len:255"`
	Age  int    `json:"age" validate:"required|int|min:18|max:60"`
}

type OutputStudent struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Age  int       `json:"age"`
}

type OutputStudents struct {
	Students []OutputStudent `json:"students"`
}

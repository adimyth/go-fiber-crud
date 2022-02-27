package schema

type ToDo struct {
	Title       string `validate:"required,min=3,max=32"`
	Description string `validate:"required,min=3,max=256"`
	Done        bool
}

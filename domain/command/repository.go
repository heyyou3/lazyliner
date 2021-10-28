package command

type Repository interface {
	Fetch() []Command
}

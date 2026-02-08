package person

var (
	Public = 1
	private = 2
)

type Person struct {
	ID int
	Name string
}

func (p Person) GetId() int {
	return p.ID
}


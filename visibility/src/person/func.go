package person


func NewPerson(id int, name string) *Person {
	return &Person{
		ID: 1,
		Name: "Max",
	}
}

func GetId(p *Person) int {
	return p.ID
}

func GetName(p *Person) string {
	return p.Name
}

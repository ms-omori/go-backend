package domain

type User struct {
	ID           int
	Name         string
	Organization Organization
}

type Organization struct {
	ID   int
	Name string
}

package domain

//go:generate mockery --case=snake --outpkg=mocks --output=internal/repository/mocks --name=CarRepository

type Car struct {
	ID    string `json:"id"`
	Mark  string `json:"marca"`
	Model string `json:"model"`
	Price uint   `json:"price"`
}

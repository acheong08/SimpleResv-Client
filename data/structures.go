package Data

type Item struct {
	Id        int
	Name      string
	Details   string
	Available bool
	Status    string
}
type Request struct {
	Email       string
	Password    string
	Action      string
	Id          int
	Name        string
	Details     string
	Available   bool
	Status      string
	AddEmail    string
	AddPassword string
}
type Account struct {
	Email string
}

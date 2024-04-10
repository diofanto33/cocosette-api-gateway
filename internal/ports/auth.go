package ports

type AuthPort interface {
	Register() error
	Login() error
	Validate() error
}

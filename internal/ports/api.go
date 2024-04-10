package ports

type APIPort interface {
	Register() error
	Login() error
	Validate() error

	CreatePatient() error
}

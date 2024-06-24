package create

const (
	UseCommand       = `create`
	ShortDescription = `Creates a connection`
	LongDescription  = `This command creates connections, 
that you can reuse when connecting via ssh`
)

const (
	Placeholder = " must not be empty"
)

var (
	NameAlias    = "Alias"
	NameAddress  = "Address"
	NameLogin    = "Login"
	NamePassword = "Password"
)

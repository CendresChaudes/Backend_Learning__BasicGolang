package user

func Login() string {
	return "Leo"
}

type Profile struct {
	name string
}

func New(name string) Profile {
	return Profile{name: name}
}

func (p Profile) Name() string {
	return p.name
}

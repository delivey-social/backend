package shared

type Usuario struct {
	email    Email
	telefone Telefone
}

func NewUsuario(email Email, telefone Telefone) Usuario {
	return Usuario{
		email:    email,
		telefone: telefone,
	}
}

func (usuario *Usuario) Email() Email {
	return usuario.email
}
func (usuario *Usuario) Telefone() Telefone {
	return usuario.telefone
}

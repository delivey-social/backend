package shared

import "net/mail"

type Usuario struct {
	email    mail.Address
	telefone Telefone
}

func NewUsuario(email mail.Address, telefone Telefone) Usuario {
	return Usuario{
		email:    email,
		telefone: telefone,
	}
}

func (usuario *Usuario) Email() mail.Address {
	return usuario.email
}
func (usuario *Usuario) Telefone() Telefone {
	return usuario.telefone
}

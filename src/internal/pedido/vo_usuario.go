package pedido

type Usuario struct {
	email    Email
	telefone Telefone
	nome     string
}

func NewUsuario(email Email, telefone Telefone, nome string) Usuario {
	return Usuario{
		email:    email,
		telefone: telefone,
		nome:     nome,
	}
}

func (usuario *Usuario) Email() Email {
	return usuario.email
}
func (usuario *Usuario) Telefone() Telefone {
	return usuario.telefone
}

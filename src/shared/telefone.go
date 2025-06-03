package shared

type Telefone struct {
	Value string
}

// TODO: Validate telefone
func NewTelefone(input string) Telefone {
	return Telefone{
		Value: input,
	}
}

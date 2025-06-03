package shared

type Telefone struct {
	value string
}

// TODO: Validate telefone
func NewTelefone(input string) Telefone {
	return Telefone{
		value: input,
	}
}

func (telefone *Telefone) String() string {
	return telefone.value
}

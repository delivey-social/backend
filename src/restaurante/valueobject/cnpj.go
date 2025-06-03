package valueobject

type CNPJ struct {
	value string
}

// TODO: Add CNPJ validation
func NewCNPJ(input string) (CNPJ, error) {
	return CNPJ{
		value: input,
	}, nil
}

func (cnpj *CNPJ) String() string {
	return cnpj.value
}

package valueobject

type CNPJ struct {
	Value string
}

// TODO: Add CNPJ validation
func NewCNPJ(input string) (CNPJ, error) {
	return CNPJ{
		Value: input,
	}, nil
}

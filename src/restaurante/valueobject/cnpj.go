package valueobject

type CNPJ struct {
	Value string
}

func NewCNPJ(input string) (CNPJ, error) {
	return CNPJ{
		Value: input,
	}, nil
}

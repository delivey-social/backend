package shared

type CEP struct {
	Value string
}

// TODO: Add CEP validation
func NewCEP(input string) (CEP, error) {
	return CEP{
		Value: input,
	}, nil
}

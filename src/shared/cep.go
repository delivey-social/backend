package shared

type CEP struct {
	value string
}

// TODO: Add CEP validation
func NewCEP(input string) (CEP, error) {
	return CEP{
		value: input,
	}, nil
}

func (cep *CEP) String() string {
	return cep.value
}

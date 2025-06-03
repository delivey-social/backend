package shared

type Email struct {
	Value string
}

// TODO: Check if is a valid email
func NewEmail(input string) Email {
	return Email{
		Value: input,
	}
}

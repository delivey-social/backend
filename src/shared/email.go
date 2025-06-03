package shared

type Email struct {
	value string
}

// TODO: Check if is a valid email
func NewEmail(input string) Email {
	return Email{
		value: input,
	}
}

func (email *Email) String() string {
	return email.value
}

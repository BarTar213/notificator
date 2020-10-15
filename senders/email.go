package senders

type Email struct {
	TemplateName string
	Recipients   []int
	Data         map[string]string
}

func (e *Email) Send() {
}

func (e *Email) GetTemplateName() string {
	return e.TemplateName
}

func (e *Email) GetData() map[string]string {
	return e.Data
}

package notifications

type Internal struct {
	TemplateID int
	Recipients []int
	Data       map[string]string
}

func (i *Internal) Send() {

}

func (i *Internal) ParseTemplate() {

}

package forms

type FormData struct {
	Values map[string]string
	Errors map[string]string
}

func NewFormData() FormData {
	return FormData{
		Values: make(map[string]string),
		Errors: make(map[string]string),
	}
}

type PageData struct {
	Data interface{}
	Form FormData
}

func NewPageData(data interface{}, form FormData) PageData {
	return PageData{
		Data: data,
		Form: form,
	}
}

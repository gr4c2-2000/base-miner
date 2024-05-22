package easyview

type ViewInterface interface {
	GetTitle() string
	GetEndpointName() string
	GetTemplate() string
}

type View struct {
	Title        string
	EndpointName string
	Template     string
}

func (v *View) GetTitle() string {
	return v.Title
}
func (v *View) GetEndpointName() string {
	return v.EndpointName
}
func (v *View) GetTemplate() string {
	return v.Template
}

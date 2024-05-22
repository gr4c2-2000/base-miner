package easyview

type EndpointInterface interface {
	GetDataSource() string
	GetName() string
	GetMethod() string
	GetContentResult() string
	GetQueryName() string
}
type Endpoint struct {
	DataSource    string
	Name          string
	Method        string
	ContentResult string
	QueryName     string
}

func (e *Endpoint) GetDataSource() string {
	return e.DataSource
}
func (e *Endpoint) GetName() string {
	return e.Name
}
func (e *Endpoint) GetMethod() string {
	return e.Method
}
func (e *Endpoint) GetContentResult() string {
	return e.ContentResult
}
func (e *Endpoint) GetQueryName() string {
	return e.QueryName
}

package easyview

type QueryInterface interface {
	GetName() string
	GetParams() []string
	GetQueryNames() []string
	GetResultTransformation() []TransformField
}

type Query struct {
	Name                 string
	Params               []string
	QueryNames           []string
	ResultTransformation []TransformField
}

func (q *Query) GetName() string {
	return q.Name
}
func (q *Query) GetParams() []string {
	return q.Params
}
func (q *Query) GetQueryNames() []string {
	return q.QueryNames
}
func (q *Query) GetResultTransformation() []TransformField {
	return q.ResultTransformation
}

type TransformFieldInterface interface {
	GetResultField() string
	GetFunction() string
	GetArguments() []string
}

func (t *TransformField) GetResultField() string {
	return t.ResultField
}
func (t *TransformField) GetFunction() string {
	return t.Function
}
func (t *TransformField) GetArguments() []string {
	return t.Arguments
}

type TransformField struct {
	ResultField string   `yaml:"resultField" validate:"required"`
	Function    string   `yaml:"function" validate:"required"`
	Arguments   []string `yaml:"arguments"`
}

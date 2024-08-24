package elasticsearch

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/gr4c2-2000/base-miner/pkg/config"
	"github.com/rotisserie/eris"
)

type ElasticSearchConnector struct {
	config      *config.DataSource
	esInterface ElasticSearchInterface
}

func (e *ElasticSearchConnector) Replace(ctx context.Context, Index string, Type string, Id string, query map[string]interface{}) error {

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return eris.Wrapf(err, "")
	}

	err := e.esInterface.Replace(ctx, Index, Type, Id, &buf)
	if err != nil {
		return eris.Wrapf(err, "")
	}
	return nil
}
func (e *ElasticSearchConnector) Create(ctx context.Context, index string, docType string, query map[string]interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return eris.Wrapf(err, "")
	}

	err := e.esInterface.Create(ctx, index, docType, &buf)
	if err != nil {
		return eris.Wrapf(err, "")
	}
	return nil
}

func (e *ElasticSearchConnector) BulkIndexDocuments(ctx context.Context, index string, docType string, documents []map[string]interface{}) error {
	return e.esInterface.BulkIndexDocuments(ctx, index, docType, documents)
}

func (e *ElasticSearchConnector) ExecuteQuery(ctx context.Context, Index string, Type string, query map[string]interface{}) (*GenericResponse, error) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, eris.Wrapf(err, "")
	}

	bufResponse, err := e.esInterface.Search(ctx, Index, Type, &buf)
	if err != nil {
		return nil, eris.Wrapf(err, "")
	}
	newStr := bufResponse.String()
	ResposeStruct := GenericResponse{}
	err = json.Unmarshal([]byte(newStr), &ResposeStruct)
	if err != nil {
		return nil, eris.Wrapf(err, "")
	}
	return &ResposeStruct, nil
}

func (e *ElasticSearchConnector) GetById(ctx context.Context, Index string, Type string, id string) (*GenericResponse, error) {
	bufResponse, err := e.esInterface.GetById(ctx, Index, Type, id)
	if err != nil {
		return nil, eris.Wrapf(err, "")
	}
	newStr := bufResponse.String()
	ResposeStruct := GenericResponse{}
	err = json.Unmarshal([]byte(newStr), &ResposeStruct)
	if err != nil {
		return nil, eris.Wrapf(err, "")
	}
	return &ResposeStruct, nil
}

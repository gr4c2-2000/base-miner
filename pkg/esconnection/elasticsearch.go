package esconnection

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/gr4c2-2000/base-miner/pkg/config"
	"github.com/rotisserie/eris"
)

type ElasticSearchGatway struct {
	config      *config.DataSource
	esInterface ElasticSearchInterface
}

func (e *ElasticSearchGatway) Replace(Index string, Type string, Id string, query map[string]interface{}) error {

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return eris.Wrapf(err, "")
	}

	err := e.esInterface.Replace(context.Background(), Index, Type, Id, &buf)
	if err != nil {
		return eris.Wrapf(err, "")
	}
	return nil
}

func (e *ElasticSearchGatway) ExecuteQuery(Index string, Type string, query map[string]interface{}) (*ResposeES, error) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, eris.Wrapf(err, "")
	}

	bufResponse, err := e.esInterface.Search(context.Background(), Index, Type, &buf)
	if err != nil {
		return nil, eris.Wrapf(err, "")
	}
	newStr := bufResponse.String()
	ResposeStruct := ResposeES{}
	err = json.Unmarshal([]byte(newStr), &ResposeStruct)
	if err != nil {
		return nil, eris.Wrapf(err, "")
	}
	return &ResposeStruct, nil
}

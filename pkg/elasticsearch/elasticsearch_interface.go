package elasticsearch

import (
	"bytes"
	"context"
	"io"
	"net"

	"github.com/gr4c2-2000/base-miner/pkg/config"
)

type ElasticSearchInterface interface {
	SetConfig(Config *config.DataSource)
	SetESConfig(EsConfig *config.ElasticSearchConfig)
	GetESConfig() *config.ElasticSearchConfig
	Search(ctx context.Context, index string, docType string, query io.Reader) (*bytes.Buffer, error)
	Replace(ctx context.Context, index string, docType string, id string, document io.Reader) error
	Create(ctx context.Context, index string, docType string, document io.Reader) error
	getDialerContext() func(ctx context.Context, network, address string) (net.Conn, error)
	BulkIndexDocuments(ctx context.Context, index string, docType string, documents []map[string]interface{}) error
	setConnection()
}

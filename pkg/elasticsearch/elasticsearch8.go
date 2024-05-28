package elasticsearch

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/gr4c2-2000/base-miner/pkg/common"
	"github.com/gr4c2-2000/base-miner/pkg/config"
	"github.com/gr4c2-2000/base-miner/pkg/logger"
	"github.com/rotisserie/eris"
	"golang.org/x/net/proxy"
)

type ElasticSearchGatway8 struct {
	config   *config.DataSource
	esConfig *config.ElasticSearchConfig
	client   *elasticsearch.Client
}

func (e *ElasticSearchGatway8) SetConfig(Config *config.DataSource) {
	e.config = Config
}
func (e *ElasticSearchGatway8) SetESConfig(EsConfig *config.ElasticSearchConfig) {
	e.esConfig = EsConfig
}
func (e *ElasticSearchGatway8) GetESConfig() *config.ElasticSearchConfig {
	return e.esConfig
}
func (e *ElasticSearchGatway8) getDialerContext() func(ctx context.Context, network, address string) (net.Conn, error) {

	if e.esConfig.SocksProxy != "" {
		dialer, err := proxy.SOCKS5(common.TCP, e.esConfig.SocksProxy, nil, proxy.Direct)
		logger.CriticalErrorLog(err)
		contextDialer, _ := dialer.(proxy.ContextDialer)

		return contextDialer.DialContext
	}
	baseDialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}
	return (baseDialer).DialContext
}

func (e *ElasticSearchGatway8) setConnection() {
	cfg := elasticsearch.Config{
		Addresses: []string{e.esConfig.Host},
		Transport: &http.Transport{
			DisableKeepAlives:     true,
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: 30 * time.Second,
			DialContext:           e.getDialerContext(),
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: e.esConfig.IgnoreSSL,
			},
		},
	}
	var err error
	e.client, err = elasticsearch.NewClient(cfg)
	if err != nil {
		logger.CriticalErrorLog(err)
	}
}
func (e *ElasticSearchGatway8) Replace(ctx context.Context, index string, docType string, id string, document io.Reader) error {
	req := esapi.IndexRequest{Index: index, DocumentID: id, Body: document}
	res, err := req.Do(ctx, e.client.Transport)
	if err != nil {
		return eris.Wrapf(err, "")
	}
	defer res.Body.Close()
	if res.IsError() {
		return eris.New(res.String())
	}
	return nil
}

func (e *ElasticSearchGatway8) Create(ctx context.Context, index string, docType string, document io.Reader) error {

	res, err := e.client.Index(index, document, e.client.Index.WithContext(ctx))
	if err != nil {
		return eris.Wrapf(err, "")
	}
	defer res.Body.Close()
	if res.IsError() {
		return eris.New(res.String())
	}
	return nil
}

func (e *ElasticSearchGatway8) BulkIndexDocuments(ctx context.Context, index string, docType string, documents []map[string]interface{}) error {
	var buf bytes.Buffer
	for _, doc := range documents {
		meta := []byte(fmt.Sprintf(`{ "index" : { "_id" : "%v" } }%s`, doc["id"], "\n"))
		data, err := json.Marshal(doc)
		if err != nil {
			return err
		}
		data = append(data, "\n"...)
		buf.Write(meta)
		buf.Write(data)
	}

	req := esapi.BulkRequest{
		Index:   index,
		Body:    &buf,
		Refresh: "true",
	}
	res, err := req.Do(ctx, e.client)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error response status: %s", res.String())
	}
	return nil
}

func (e *ElasticSearchGatway8) Search(ctx context.Context, index string, docType string, query io.Reader) (*bytes.Buffer, error) {
	res, err := e.client.Search(
		e.client.Search.WithContext(ctx),
		e.client.Search.WithIndex(index),
		e.client.Search.WithBody(query),
		e.client.Search.WithPretty(),
	)

	if err != nil {
		return nil, eris.Wrapf(err, "")
	}
	defer res.Body.Close()
	if res.IsError() {
		return nil, eris.Wrapf(err, "")
	}

	bufResponse := new(bytes.Buffer)
	bufResponse.ReadFrom(res.Body)
	return bufResponse, nil
}

package esconnection

import (
	"github.com/gr4c2-2000/base-miner/pkg/config"
	"github.com/gr4c2-2000/base-miner/pkg/error_handlers"
	"github.com/gr4c2-2000/base-miner/pkg/logger"
	"github.com/rotisserie/eris"
)

type EsMap map[string]*ElasticSearchGatway

func InitES(Config *config.DataSource) EsMap {

	EsConnectionsArray := make(map[string]*ElasticSearchGatway, 0)
	var es ElasticSearchInterface
	for _, esConfig := range Config.ElasticSearch {
		switch esConfig.Version {
		case 5:
			es = &ElasticSearchGatway5{}
		case 8:
			es = &ElasticSearchGatway8{}
		default:
			err := error_handlers.ES_VERSION_NOT_SUPPORTED
			logger.CriticalErrorLog(err)
		}
		es.SetConfig(Config)
		es.SetESConfig(&esConfig)
		es.setConnection()
		ElasticSearchService := ElasticSearchGatway{Config, es}
		EsConnectionsArray[esConfig.DatabaseName] = &ElasticSearchService
	}
	return EsConnectionsArray
}

func (em EsMap) GetConnectionByName(Name string) (*ElasticSearchGatway, error) {
	val, ok := em[Name]
	if !ok {
		return nil, eris.New("No ES Connection with specified Name")
	}
	return val, nil
}

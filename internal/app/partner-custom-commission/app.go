package partnercustomcommission

// import (
// 	"github.com/gr4c2-2000/base-miner/internal/pkg/kernal"
// 	"github.com/gr4c2-2000/base-miner/internal/pkg/repository"
// )

// type App struct {
// 	*kernal.Kernal
// 	Service       *Service
// 	PartnerConfig PartnerConfig
// }

// func InitApp(kernal *kernal.Kernal) *App {
// 	esRepo := repository.ElasticSearchRepository{EsMap: kernal.ElasticGateway}
// 	mysql := repository.MysqlRepository{MySqlGateway: *kernal.MysqlGatway}
// 	Service := InitService(&esRepo, &mysql)
// 	app := App{Kernal: kernal, Service: Service}
// 	return &app
// }

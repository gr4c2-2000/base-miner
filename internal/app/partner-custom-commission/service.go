package partnercustomcommission

// import (
// 	"fmt"
// 	"os"

// 	"github.com/gr4c2-2000/base-miner/internal/pkg/data"
// 	"github.com/gr4c2-2000/base-miner/internal/pkg/repository"
// )

// type Service struct {
// 	ElasticSearchRepository *repository.ElasticSearchRepository
// 	MySqlRepository         *repository.MysqlRepository
// 	BankMidMap              *data.BankMidMap
// }

// func InitService(ElasticSearchRepository *repository.ElasticSearchRepository, MySqlRepository *repository.MysqlRepository) *Service {
// 	serv := Service{ElasticSearchRepository: ElasticSearchRepository, MySqlRepository: MySqlRepository}
// 	return &serv
// }

// func (s *Service) GatherData() {
// 	BankMibQuery, err := data.NewBankMidQuery()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	err = s.MySqlRepository.QueryFromFile(BankMibQuery)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	map1 := data.InitBankMidMap(BankMibQuery.Result)
// 	fmt.Println(map1)
// 	os.Exit(1)
// 	comprest, err := data.NewTransactionsAffiliantProgram(89558, "202211")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	err = s.MySqlRepository.QueryFromFile(comprest)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }

package global

import (
	"sync"

	"github.com/knightoverwatch/C2N-Launchpad-go/config"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	LP_DB     *gorm.DB
	LP_DBList map[string]*gorm.DB
	LP_CONFIG config.Server
	LP_VP     *viper.Viper
	lock      sync.RWMutex
)

// MustGetGlobalDBByDBName 通过名称获取db 如果不存在则panic
func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := LP_DBList[dbname]
	if !ok || db == nil {
		panic("db no init")
	}
	return db
}

package agollo

import (
	"github.com/zouyx/agollo/v2/agcache"
	. "github.com/zouyx/agollo/v2/component/log"
	"github.com/zouyx/agollo/v2/component/notify"
	"github.com/zouyx/agollo/v2/env"
	_ "github.com/zouyx/agollo/v2/env"
)

//InitCustomConfig init config by custom
func InitCustomConfig(loadAppConfig func() (*env.AppConfig, error)) {

	env.InitConfig(loadAppConfig)

	initDefaultConfig()
}

//start apollo
func Start() error {
	return startAgollo()
}

//SetLogger 设置自定义logger组件
func SetLogger(loggerInterface LoggerInterface) {
	if loggerInterface != nil {
		InitLogger(loggerInterface)
	}
}

//SetCache 设置自定义cache组件
func SetCache(cacheFactory *agcache.DefaultCacheFactory) {
	if cacheFactory != nil {
		initConfigCache(cacheFactory)
	}
}

//StartWithLogger 通过自定义logger启动agollo
func StartWithLogger(loggerInterface LoggerInterface) error {
	SetLogger(loggerInterface)
	return startAgollo()
}

//StartWithCache 通过自定义cache启动agollo
func StartWithCache(cacheFactory *agcache.DefaultCacheFactory) error {
	SetCache(cacheFactory)
	return startAgollo()
}

func startAgollo() error {
	//init server ip list
	go initServerIpList()
	//first sync
	if err := notifySyncConfigServices(); err != nil {
		return err
	}
	Logger.Debug("init notifySyncConfigServices finished")

	//start long poll sync config
	go StartRefreshConfig(&notify.NotifyConfigComponent{})

	Logger.Info("agollo start finished ! ")

	return nil
}

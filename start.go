package agollo

import "context"

//start apollo
func Start() error {
	return StartWith(context.TODO(), nil)
}

func StartWithLogger(ctx context.Context, loggerInterface LoggerInterface) error {
	return StartWith(ctx, loggerInterface)
}

func StartWith(ctx context.Context, loggerInterface LoggerInterface) error {
	if loggerInterface != nil {
		initLogger(loggerInterface)
	}

  //init server ip list
  go initServerIpList(ctx)

	//first sync
	err := notifySyncConfigServices()

	//first sync fail then load config file
	if err !=nil{
		config, _ := loadConfigFile(appConfig.BackupConfigPath)
		if config!=nil{
			updateApolloConfig(config,false)
		}
	}

	//start long poll sync config
	go StartRefreshConfig(ctx, &NotifyConfigComponent{})

	logger.Info("agollo start finished , error:",err)
	
	return err
}
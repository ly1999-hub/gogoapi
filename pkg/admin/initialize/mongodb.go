package initialize

import (
	"myapp/internal/module/mongodb"
	"myapp/pkg/admin/config"
)

//InitMongoDB ...
func InitMongoDB() {
	cfg := config.GetENV().MongoDB

	err := mongodb.Connect(
		cfg.Host,
		cfg.User,
		cfg.Password,
		cfg.DBName,
		cfg.Mechanism,
		cfg.Source,
	)

	if err != nil {
		panic(err)
	}

}

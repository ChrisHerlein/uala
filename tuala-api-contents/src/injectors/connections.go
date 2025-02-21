package injectors

import (
	"fmt"

	beanstalk "github.com/beanstalkd/go-beanstalk"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/ChrisHerlein/uala/tuala-api-contents/src/config"
)

type Connections struct {
	PostgreSQL *gorm.DB
	Redis      *redis.Client
	Beanstalk  *beanstalk.Conn
}

func GetConnections(cfg *config.Config) *Connections {
	return &Connections{
		PostgreSQL: getPostgreSql(cfg),
		Redis:      getRedis(cfg),
		Beanstalk:  getBeanstalk(cfg),
	}
}

func getPostgreSql(cfg *config.Config) *gorm.DB {
	connStr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.PgHost, cfg.PgUser, cfg.PgPassword, cfg.PgDb, cfg.PgPort,
	)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

func getRedis(cfg *config.Config) *redis.Client {
	url := fmt.Sprintf(
		"redis://%s/0?protocol=3",
		cfg.RedisHost,
	)
	opts, err := redis.ParseURL(url)
	if err != nil {
		panic(err)
	}

	return redis.NewClient(opts)
}

func getBeanstalk(cfg *config.Config) *beanstalk.Conn {
	conn, err := beanstalk.Dial("tcp", cfg.BeanstalkHost)
	if err != nil {
		panic(err)
	}

	return conn
}

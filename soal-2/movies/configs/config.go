package configs

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/schema"
	mysqlhook "github.com/muhfaris/logrus-mysql-hook"
)

// App is represent all data for application
type App struct {
	Name         string
	Port         int
	Environment  string
	Services     ServicesApp
	API          APIApp
	Decoder      *schema.Decoder
	Logger       *logrus.Logger
	MysqlHook    *mysqlhook.Hook
	WriteTimeout int
	ReadTimeout  int
	IdleTimeout  int
}

// ServicesApp is available all service for application
type ServicesApp struct {
	DB *sql.DB
}

// APIApp is some apis
type APIApp struct {
	OmbdbAPI       string
	OmbdbAPISecret string
}

// CreateApp is create new object od app
func CreateApp() *App {
	// initialize Services
	// initialize database SQL
	db, err := initDatabase()
	if err != nil {
		log.Fatalf("error:%v", err)
		return nil
	}

	mysqlHook := mysqlhook.Default(db, "log")
	//defer mysqlHook.Flush()

	// initialize logger
	logger := logrus.New()
	logger.AddHook(mysqlHook)

	return &App{
		Name:        viper.GetString("app.name"),
		Port:        viper.GetInt("app.port"),
		Environment: viper.GetString("app.environment"),
		Services: ServicesApp{
			DB: db,
		},
		API: APIApp{
			OmbdbAPI:       viper.GetString("api.ombdb.api"),
			OmbdbAPISecret: viper.GetString("api.ombdb.secret_key"),
		},
		Logger:       logger,
		MysqlHook:    mysqlHook,
		Decoder:      schema.NewDecoder(),
		WriteTimeout: viper.GetInt("apps.http.write_timeout"),
		ReadTimeout:  viper.GetInt("apps.http.read_timeout"),
		IdleTimeout:  viper.GetInt("apps.http.idle_timeout"),
	}
}

func initDatabase() (*sql.DB, error) {
	// dsn format  "root@(127.0.0.1:3306)/db?charset=utf8&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(
		"%s:%s@(%s:%d)/%s?charset=utf8&parseTime=true",
		viper.GetString("storage.database.username"),
		viper.GetString("storage.database.password"),
		viper.GetString("storage.database.host"),
		viper.GetInt("storage.database.port"),
		viper.GetString("storage.database.name"),
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println("connection database use default value")
		defaultDSN := "root@(127.0.0.1:3306)/db?charset=utf8&parseTime=True&loc=Local"
		db, err = sql.Open("mysql", defaultDSN)
		if err != nil {
			return nil, err
		}
	}

	log.Println("Database connected ...")
	return db, nil
}

package gorm

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"github.com/tosashimanto/iot-camera-web/bootstrap"
)

var (
	db *gorm.DB
)

// initialize database
func Init() {
	var adapter string
	adapter = bootstrap.App.DBConfig.String("adapter")
	switch adapter {
	case "mysql":
		mysqlConn()
		break
	case "postgres":
		postgresConn()
		break
	}
}

// setupPostgresConn: setup postgres database connection using the configuration from database.yaml
func postgresConn() {
	var (
		connectionString string
		err              error
	)

	connectionString = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s", bootstrap.App.DBConfig.String("username"), bootstrap.App.DBConfig.String("password"), bootstrap.App.DBConfig.String("host"), bootstrap.App.DBConfig.String("port"), bootstrap.App.DBConfig.String("database"), bootstrap.App.DBConfig.String("sslmode"))
	fmt.Println("connectionString=", connectionString)

	// if model_db, err = gorm.Open("postgres", "host=localhost user=flect001 dbname=sthing_gsurvey sslmode=disable password=123456"); err != nil {
	if db, err = gorm.Open("postgres", connectionString); err != nil {
		panic(err)
	}

	if err = db.DB().Ping(); err != nil {
		panic(err)
	}

	// テーブル名が複数系でない場合の指定
	db.SingularTable(true)
	db.LogMode(true)
	// model_db.Exec("CREATE EXTENSION postgis")

	db.DB().SetMaxIdleConns(bootstrap.App.DBConfig.Int("idle_conns"))
	db.DB().SetMaxOpenConns(bootstrap.App.DBConfig.Int("open_conns"))
}

// mysqlConn: setup mysql database connection using the configuration from database.yaml
func mysqlConn() {
	var (
		connectionString string
		err              error
	)

	connectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", bootstrap.App.DBConfig.String("username"), bootstrap.App.DBConfig.String("password"), bootstrap.App.DBConfig.String("host"), bootstrap.App.DBConfig.String("port"), bootstrap.App.DBConfig.String("database"))

	if db, err = gorm.Open("mysql", connectionString); err != nil {
		panic(err)
	}

	if err = db.DB().Ping(); err != nil {
		panic(err)
	}
	// テーブル名が複数系でない場合の指定
	db.SingularTable(true)
	db.LogMode(true)

	db.DB().SetMaxIdleConns(bootstrap.App.DBConfig.Int("idle_conns"))
	db.DB().SetMaxOpenConns(bootstrap.App.DBConfig.Int("open_conns"))
}

/*
 * Gorm: return GORM's postgres database connection instance.
 */
func DBManager() *gorm.DB {
	return db
}

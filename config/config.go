package config

import "fmt"

type ConfigStruct struct {
	DriverName 		string
	DataSourceName 	dataSourceName 
}

type dataSourceName struct {
	user 		string
	dbname 		string
	dbpassword 	string
	sslmode 	string
}

func (d *dataSourceName) Set() {
	d.user = "postgres"
	d.dbname = "shopapp"
	d.dbpassword = ""
	d.sslmode = "disable"
}

func (d *dataSourceName) Get() string {
	return fmt.Sprintf("user=%v dbname=%v password=%v sslmode=%v",
		d.user,d.dbname,d.dbpassword,d.sslmode,
	)
}

func CreateConfig() (ConfigStruct) {
	data := dataSourceName{}
	data.Set()

	return ConfigStruct{
		DriverName: "postgres",
		DataSourceName: data,
	}
}



module github.com/smartwalle/utube/examples

go 1.16

require (
	github.com/go-sql-driver/mysql v1.4.1 // indirect
	github.com/smartwalle/dbs v1.0.0 // indirect
	github.com/smartwalle/utube v0.0.0
)

replace (
	github.com/smartwalle/utube => ../
)

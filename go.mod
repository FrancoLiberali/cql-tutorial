module github.com/ditrit/badaas-orm-tutorial

go 1.18

require (
	github.com/ditrit/badaas v0.0.0-20230911080244-4b834e9e4926
	gorm.io/driver/sqlite v1.5.3
	gorm.io/gorm v1.25.4
)

require (
	github.com/elliotchance/pie/v2 v2.7.0 // indirect
	github.com/google/uuid v1.3.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-sqlite3 v1.14.17 // indirect
	golang.org/x/exp v0.0.0-20230817173708-d852ddb80c63 // indirect
)

// TODO este deberia ser temporal
replace gorm.io/driver/sqlite v1.5.3 => github.com/ditrit/sqlite v0.0.0-20230906140046-2f37a3f972de

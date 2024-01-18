module mpesa

go 1.19

require gorm.io/driver/mysql v1.3.6

require github.com/google/uuid v1.0.0 // indirect

require (
	github.com/bitactro/UUIDv4 v0.0.0-20220313143507-9a5732793937
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/joho/godotenv v1.4.0
	github.com/pborman/uuid v1.2.1
	gorm.io/gorm v1.23.10 // indirect
)

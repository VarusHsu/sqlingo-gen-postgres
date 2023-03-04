# Sqlingo-gen-postgres
#### A postgres orm model code generator for sqlingo
#### Wanna more information about sqlingo, see:
#### https://github.com/lqs/sqlingo
Feature:
*  Supported postgres array type

Tips:
* Array will be parsed to string.

Usage:
```shell
# install
go install github.com/VarusHsu/sqlingo-gen-postgres@latest
# append to your env (optional)
export PATH="$PATH:$(go env GOPATH)/bin"
# usage 
mkdir -p generated/sqlingo
sqlingo-gen-postgres \
-host localhost \
-user postgres \
-password password \
-port 5432 \
-db your_db \
> generated/sqlingo/database_name.dsl.go  
```

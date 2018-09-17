## Prerequisites

Always keep Go as fresh as possible.

```
brew install go # version 1.11
export PATH=$PATH:$GOPATH/bin
export GO111MODULE=on
```

Install dependency libraries.

```
GO111MODULE=off go get -u github.com/sarawoot/erp-api
go get
```

### Setup Local Databases

Install PostgresQL locally and create new databases for dev and test with following commands

```bash
psql -U postgres -c "CREATE USER erp_dev WITH SUPERUSER PASSWORD 'erp_dev';"
createdb erp_dev -O erp_dev
psql -U postgres -c "ALTER DATABASE erp_dev SET timezone TO 'UTC';"

psql -U postgres -c "CREATE USER erp_test WITH SUPERUSER PASSWORD 'erp_test';"
createdb erp_test -O erp_dev
psql -U postgres -c "ALTER DATABASE erp_test SET timezone TO 'UTC';"
```
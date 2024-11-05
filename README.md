# UMetrika_server

### dependencies
- `brew install goose`
- `goose -dir migrations create {{table_name}} sql`
- `export GOOSE_DRIVER=postgres`
- `export GOOSE_DBSTRING=postgresql://admin:321@127.0.0.1:9876/inside_study?sslmode=disable`
- теперь можно поднимать миграции
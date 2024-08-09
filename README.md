# neversitup-test

### My structure project for Go

```
├── Dockerfile
├── Makefile
├── README.md
├── api
│   ├── account.go
│   ├── account_test.go
│   ├── main_test.go
│   ├── middleware.go
│   ├── middleware_test.go
│   ├── server.go
│   ├── token.go
│   ├── transfer.go
│   ├── transfer_test.go
│   ├── user.go
│   ├── user_test.go
│   └── validator.go
├── app.env
├── db
│   ├── migration
│   │   ├── 000001_init_schema.down.sql
│   │   ├── 000001_init_schema.up.sql
│   │   ├── 000002_add_users.down.sql
│   │   ├── 000002_add_users.up.sql
│   │   ├── 000003_add_sessions.down.sql
│   │   └── 000003_add_sessions.up.sql
│   ├── mock
│   │   └── store.go
│   ├── query
│   │   ├── account.sql
│   │   ├── entry.sql
│   │   ├── session.sql
│   │   ├── transfer.sql
│   │   └── user.sql
│   └── sqlc
│   ├── account.sql.go
│   ├── account_test.go
│   ├── db.go
│   ├── entry.sql.go
│   ├── main_test.go
│   ├── models.go
│   ├── querier.go
│   ├── session.sql.go
│   ├── store.go
│   ├── store_test.go
│   ├── transfer.sql.go
│   ├── user.sql.go
│   └── user_test.go
├── doc
│   ├── db.dbml
│   ├── schema.sql
│   ├── statik
│   │   └── statik.go
│   └── swagger
│   ├── favicon-16x16.png
│   ├── favicon-32x32.png
│   ├── index.css
│   ├── index.html
│   ├── oauth2-redirect.html
│   ├── simple_bank.swagger.json
│   ├── swagger-initializer.js
│   ├── swagger-ui-bundle.js
│   ├── swagger-ui-bundle.js.map
│   ├── swagger-ui-es-bundle-core.js
│   ├── swagger-ui-es-bundle-core.js.map
│   ├── swagger-ui-es-bundle.js
│   ├── swagger-ui-es-bundle.js.map
│   ├── swagger-ui-standalone-preset.js
│   ├── swagger-ui-standalone-preset.js.map
│   ├── swagger-ui.css
│   ├── swagger-ui.css.map
│   ├── swagger-ui.js
│   └── swagger-ui.js.map
├── docker-compose.yml
├── eks
│   ├── aws-auth.yaml
│   ├── deployment.yaml
│   └── service.yaml
├── gapi
│   ├── converter.go
│   ├── metadata.go
│   ├── rpc_create_user.go
│   ├── rpc_login_user.go
│   └── server.go
├── go.mod
├── go.sum
├── main.go
├── pb
│   ├── rpc_create_user.pb.go
│   ├── rpc_login_user.pb.go
│   ├── service_simple_bank.pb.go
│   ├── service_simple_bank.pb.gw.go
│   ├── service_simple_bank_grpc.pb.go
│   └── user.pb.go
├── proto
│   ├── google
│   │   └── api
│   │   ├── annotations.proto
│   │   ├── field_behavior.proto
│   │   ├── http.proto
│   │   └── httpbody.proto
│   ├── protoc-gen-openapiv2
│   │   └── options
│   │   ├── annotations.proto
│   │   └── openapiv2.proto
│   ├── rpc_create_user.proto
│   ├── rpc_login_user.proto
│   ├── service_simple_bank.proto
│   └── user.proto
├── sqlc.yaml
├── start.sh
├── token
│   ├── jwt_maker.go
│   ├── jwt_maker_test.go
│   ├── maker.go
│   ├── paseto_maker.go
│   ├── paseto_maker_test.go
│   └── payload.go
├── tools
│   └── tools.go
└── util
├── config.go
├── currency.go
├── password.go
├── password_test.go
└── random.go
```

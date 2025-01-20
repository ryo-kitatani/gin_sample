
ディレクトリ構造
```
.
├── README.md
├── cmd
│   └── main.go
├── docker-compose.yml
├── go.mod
├── go.sum
└── internal
    ├── domain
    │   ├── entity
    │   │   └── todo.go
    │   └── repository
    │       └── todo.go
    ├── infrastructure
    │   ├── persistence
    │   │   └── mysql
    │   │       └── todo.go
    │   └── router
    │       └── router.go
    ├── interface
    │   └── handler
    │       └── todo.go
    └── usecase
        └── todo.go
```

起動
```
go run cmd/main.go
```

一覧
```
GET localhost:8080/api/v1/todos
```

新規作成
```
POST localhost:8080/api/v1/todos

{
    "title": "勉強"
}
```

単体取得
```
GET localhost:8080/api/v1/todos/:id
```

ステータス更新
```
PUT localhost:8080/api/v1/todos/:id

{
    "status": "doing"
}
```


削除
```
DELETE localhost:8080/api/v1/todos/3
```


# wantedly2019summer_webapp
## 課題1
以下を実行する。
```
$ cd kadai1
$ docker-compose build
$ docker-compose up -d
```
リクエスト
```
$ curl -XGET -H 'Content-Type:application/json' http://localhost:8080/
```
結果
```
{"message":"Hello World!!"}
```

## 課題2
リクエスト
```
$ curl -XGET -H 'Content-Type:application/json' https://sleepy-cliffs-75300.herokuapp.com/
```
結果
```
{"message":"Hello World!!"}
```

## 課題3
以下を実行する。
```
$ cd kadai3
$ docker-compose build
$ docker-compose up
```
### 課題1

リクエスト
```
$ curl -XGET -H 'Content-Type:application/json' http://localhost:8080/
```
結果
```
{"message":"Hello World!!"}
```

### user を作成
リクエスト
```
curl -XPOST -H 'Content-Type:application/json' http://localhost:8080/users -d '{"name": "test", "email": "hoge@example.com" }'
```
結果
```
{"id":1,"name":"test","email":"hoge@example.com","created_at":"2019-05-23T23:45:16.245937Z","updated_at":"2019-05-23T23:45:16.245937Z"}
```

### user を一覧
```
curl -XGET -H 'Content-Type:application/json' http://localhost:8080/users
```
結果
```
[{"id":1,"name":"test","email":"hoge@example.com","created_at":"2019-05-23T23:42:25.660355Z","updated_at":"2019-05-23T23:42:25.660355Z"},{"id":2,"name":"test","email":"hoge@example.com","created_at":"2019-05-23T23:43:04.572914Z","updated_at":"2019-05-23T23:43:04.572914Z"},{"id":3,"name":"test","email":"hoge@example.com","created_at":"2019-05-23T23:45:16.245937Z","updated_at":"2019-05-23T23:45:16.245937Z"}]
```

### user を削除
```
$ curl -XDELETE -H 'Content-Type:application/json' http://localhost:8080/users/1
```
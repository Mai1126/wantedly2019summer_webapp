# wantedly2019summer_webapp
## 課題1
以下を実行してコンテナを立ち上げる
```bash
$ cd kadai1
$ docker-compose build #初回のみ
$ docker-compose up -d
```
コンテナを落とすときは以下を実行する
```bash
$ docker-compose down
```


- リクエスト
```bash
$ curl -XGET -H 'Content-Type:application/json' http://localhost:8080/
```
- 結果
```
{"message":"Hello World!!"}
```



## 課題2
公開URL

https://sleepy-cliffs-75300.herokuapp.com/

- リクエスト
```bash
$ curl -XGET -H 'Content-Type:application/json' https://sleepy-cliffs-75300.herokuapp.com/
```
- 結果
```
{"message":"Hello World!!"}
```

## 課題3
以下を実行してコンテナを立ち上げる
```bash
$ cd kadai3
$ docker-compose build #初回のみ
$ docker-compose up
```
コンテナを落とすときは以下を実行する
```bash
$ docker-compose down
```

#### 課題1

- リクエスト
```bash
$ curl -XGET -H 'Content-Type:application/json' http://localhost:8080/
```
- 結果
```
{"message":"Hello World!!"}
```

#### user を作成
- リクエスト
```bash
$ curl -XPOST -H 'Content-Type:application/json' http://localhost:8080/users -d '{"name": "test", "email": "hoge@example.com" }'
```
- 結果
```
{"id":1,"name":"test","email":"hoge@example.com","created_at":"2019-05-24T11:38:09.104572Z","updated_at":"2019-05-24T11:38:09.104572Z"}
```

#### user を更新
- リクエスト
```bash
curl -XPUT -H 'Content-Type:application/json' http://localhost:8080/users/1 -d '{"name": "koudaiii", "email": "hoge@example.com" }'
```
- 結果
```
{"id":1,"name":"koudaiii","email":"hoge@example.com","created_at":"2019-05-24T11:38:09.104572Z","updated_at":"2019-05-24T11:38:56.455301Z"}
```

#### user を確認
- リクエスト
```bash
$ curl -XGET -H 'Content-Type:application/json' http://localhost:8080/users/1
```
- 結果
```
{"id":1,"name":"koudaiii","email":"hoge@example.com","created_at":"2019-05-24T11:38:09.104572Z","updated_at":"2019-05-24T11:38:56.455301Z"}
```

#### user を一覧
- リクエスト
```bash
curl -XGET -H 'Content-Type:application/json' http://localhost:8080/users
```
- 結果
```
[{"id":1,"name":"koudaiii","email":"hoge@example.com","created_at":"2019-05-24T11:38:09.104572Z","updated_at":"2019-05-24T11:38:56.455301Z"}]
```

### user を削除
- リクエスト
```bash
$ curl -XDELETE -H 'Content-Type:application/json' http://localhost:8080/users/1
```
- 結果
```
```
確認すると
- リクエスト
```bash
$ curl -XGET -H 'Content-Type:application/json' http://localhost:8080/users
```
- 結果
```
null
```
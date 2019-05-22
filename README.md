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
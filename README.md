## MEMO

現在はブラウザから
http://localhost:3306/
へアクセス


## ファイル構造

### api
ここはapi,github.comのファイルに分かれている
#### api
kakomonListとswaggerのファイルに分かれている<br>
kakomonList → main.goなどのGOファイルが置かれている
swagger → GUI。RESTful APIのドキュメントや、サーバ、クライアントコード、エディタ、またそれらを扱うための仕様などを提供するフレームワーク
※RESTful API(REST API)は、Webシステムを外部から利用するためのプログラムの呼び出し規約(API)の種類の一つのこと

### db
dbが作られていないため、ここのファイルは空で、まだ反映されていない（2020/2/14）


## 環境構築まで手順
docker-composeを使って、goとMySQL環境を構築<br>
Dockerfileとdocker-compose.ymlを参照<br>

### 1 docker-compose up -d
docker 立ち上げ
### 2 docker-compose build {コンテナ名} bash
コンテナ実行
### 3　bee api {プロジェクト名}
プロジェクト作成<br>
2020/2/14現在、プロジェクト名kakomonListが存在




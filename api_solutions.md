# api設計
## api概要
* [ユーザー登録](#api-detail-user-signup)
* [ユーザーログイン](#api-detail-user-login)
* [ユーザー情報取得](#api-detail-user-info)
* [未読お知らせ数](#api-detail-unread-notification)
* [お知らせリスト](#api-detail-notification-list)
* [目標体重設定](#api-detail-set-target-weight)
* [達成率、体重体脂肪率記録取得](#api-detail-homepage-banner-info)
* [食事記録](#api-detail-ate-foods-store)
* [食事履歴](#api-detail-ate-foods-history)
* [自分の体の記録](#api-detail-body-info-store)
* [運動記録](#api-detail-sport-store)
* [自分の日記](#api-detail-my-diary-store)
* [自分の記録ページ情報取得](#api-detail-my-log-info)
* [体重体脂肪率記録取得](#api-detail-weight-fat-rate)
* [カラムページ情報取得](#api-column-page-info)

## api共通応答ステータスコード定義
正常 200

未ログインまたはtoken無効　401

権限なし 403

他の業務エラー　400

サーバー内部エラー 500


## api詳細定義
<h3 id="api-detail-user-signup">ユーザー登録</h3>

* uri:
```
  /user/signup
```
* method
  
    post

* curl例
```bash
curl --location --request POST '127.0.0.1:8080/signup' \
--header 'X-Requested-With: XMLHttpRequest' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "assad11",
    "password": "123321",
    "password_confirmation":"123321"
}'
```

* request body
```json
{
    "username": "assad11",
    "password": "123321",
    "password_confirmation":"123321"
}
```
* 正常 response body
```json
{
    "token": "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJodHRwOi8vMTI3LjAuMC4xOjgwODAvbG9naW4iLCJpYXQiOjE2ODA0MDc5OTEsImV4cCI6MTY4MDQxMTU5MSwibmJmIjoxNjgwNDA3OTkxLCJqdGkiOiJQSkVCWFRNNnVtOE9mejhJIiwic3ViIjoiMSIsInBydiI6IjIzYmQ1Yzg5NDlmNjAwYWRiMzllNzAxYzQwMDg3MmRiN2E1OTc2ZjcifQ.YJos0tGXrib2z3B6q8E95t9WSYkiODhEUTnOH05Zx_w"
}
```

* 異常 response body
```json
{
    "messege": "ユーザー登録失敗、原因:xxxxxxx"
}
```


---

<h3 id="api-detail-user-login">ユーザーログイン</h3>

* uri:
```
  /user/login
```

* method
  
    post

* curl例
```bash
curl --location --request POST '127.0.0.1:8080/login' \
--header 'X-Requested-With: XMLHttpRequest' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username":"testname",
    "password":"123321"
}'
```

* request body
```json
{
    "username": "testname",
    "password": "123321",
}
```


* 正常 response body
```json
{
    "token": "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJodHRwOi8vMTI3LjAuMC4xOjgwODAvbG9naW4iLCJpYXQiOjE2ODA0MDc5OTEsImV4cCI6MTY4MDQxMTU5MSwibmJmIjoxNjgwNDA3OTkxLCJqdGkiOiJQSkVCWFRNNnVtOE9mejhJIiwic3ViIjoiMSIsInBydiI6IjIzYmQ1Yzg5NDlmNjAwYWRiMzllNzAxYzQwMDg3MmRiN2E1OTc2ZjcifQ.YJos0tGXrib2z3B6q8E95t9WSYkiODhEUTnOH05Zx_w",
    "username":"usernamexxx"
}
```
* 異常 response body
```json
{
    "messege": "ユーザーログイン失敗、原因:xxxxxxx"
}
```
---

<h3 id="api-detail-user-info">ユーザー情報取得</h3>
ユーザーの基本情報取得用api

* uri:
```
  /user/info
```
* method
  
    get

* curl例
```bash
curl --location --request GET '127.0.0.1:8080/info' \
--header 'X-Requested-With: XMLHttpRequest' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer 登録またはログインapi取得したtoken'
```

* request header
Authorization: Bearer 登録またはログインapi取得したtoken

* 正常 response body
```json
{
    "username":"usernamexxx"
}
```
* 異常 response body
```json
{
    "messege": "ユーザーログイン失敗、原因:xxxxxxx"
}
```
---
<h3 id="api-detail-user-info">未読お知らせ数</h3>
未読お知らせ数取得用api

* uri:
```
  /unread_notification_num
```

* method

    get

* curl例
```bash
curl --location --request GET '127.0.0.1:8080/unread_notification_num' \
--header 'X-Requested-With: XMLHttpRequest' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer 登録またはログインapi取得したtoken'
```
* request header
Authorization: Bearer 登録またはログインapi取得したtoken

* 正常 response body
```json
{
    "unread_notification_num": 3
}
```
* 異常 response body
```json
{
    "messege": "取得失敗、原因:xxxxxxx"
}
```
---


<h3 id="api-detail-notification-list">お知らせリスト</h3>
未読お知らせ数取得用api

* uri:
```
  /notification_list
```

* method

    get

* request header
Authorization: Bearer 登録またはログインapi取得したtoken

* curl例
```bash
curl --location --request GET '127.0.0.1:8080/notification_list' \
--header 'X-Requested-With: XMLHttpRequest' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer 登録またはログインapi取得したtoken'
```

* 正常 response body
```json
{
    "notification_list": [
        {
            "title":"notification title",
            "content": "notification content"
        },
        {
            "title":"notification title2",
            "content": "notification content2"
        }
    ]
}
```
* 異常 response body
```json
{
    "messege": "取得失敗、原因:xxxxxxx"
}
```
---

<h3 id="api-detail-set-target-weight">目標体重設定</h3>

* uri:
```
  /target_weight
```

* method

    post

* request header
Authorization: Bearer 登録またはログインapi取得したtoken

* curl例
```bash
curl --location --request POST '127.0.0.1:8080/target_weight' \
--header 'X-Requested-With: XMLHttpRequest' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer 登録またはログインapi取得したtoken' \
--data-raw '{
    "target_weight": 49.5,
    "now_weight": 55.5
}'
```

* 正常 response body
```json
{
    "status": "success"
}
```
* 異常 response body
```json
{
    "messege": "xxxxxxx"
}
```
---


<h3 id="api-detail-homepage-banner-info">達成率、体重体脂肪率記録取得</h3>
homeページのbannerに日付、達成率、体重体脂肪率履歴書を取得api

* uri:
```
  /banner_info
```

* method

    get

* request header
Authorization: Bearer 登録またはログインapi取得したtoken

* curl例
```bash
curl --location --request GET '127.0.0.1:8080/banner_info' \
--header 'X-Requested-With: XMLHttpRequest' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer 登録またはログインapi取得したtoken' 
```

* 正常 response body
```json
{
    "now_date": "2023-01-01",
    "complete_rate": 85.5,
    "body_history": [
        {
            "date": "2022-01-01",
            "weight": 55.5,
            "fat_rate": 35.5
        },
        {
            "date": "2022-02-01",
            "weight": 55,
            "fat_rate": 35
        },
        {
            "date": "2022-03-01",
            "weight": 55,
            "fat_rate": 35
        },
    ]
}
```
* 異常 response body
```json
{
    "messege": "xxxxxxx"
}
```
---

<h3 id="api-detail-ate-foods-store">食事記録</h3>

* uri:
```
  /ate_foods
```

* method

    post


* request header
Authorization: Bearer 登録またはログインapi取得したtoken

* request body
```json
{
    "type": 1,
    "name": "おこめ",
    "pic": "cdn uri",
    "ate_at": "1970-01-01 15:00:00"
}
```
request body説明：

type：食事の類別区分番号、1:朝食　2:お昼　3:夕食 4:おやつ

pic: 写真のcdnまたはossのuri

ate_at：食事の時点

* curl例
```bash
curl --location --request POST '127.0.0.1:8080/ate_foods' \
--header 'X-Requested-With: XMLHttpRequest' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer 登録またはログインapi取得したtoken' \
--data-raw '{
    "type": 1,
    "name": "おこめ",
    "pic": "cdn uri",
    "ate_at": "1970-01-01 15:00:00"
}'
```

* 正常 response body
```json
{
    "status": "success"
}
```
* 異常 response body
```json
{
    "messege": "xxxxxxx"
}
```
---

---
<h3 id="api-detail-ate-foods-history">食事履歴</h3>

* uri:
```
  /ate_foods
```

* method

    get

* request header
Authorization: Bearer 登録またはログインapi取得したtoken

* get params
  
  skip: 0

  pagesize: 10

  type: 1

  * get params説明：
  
    skip: 指定された数の食事履歴をskip、ディフォルト　0

    pagesize: 一回取得する食事履歴数、ディフォルト　10

    type: 食事の類別区分番号、0: 制限なし　1:朝食　2:お昼　3:夕食 4:おやつ、ディフォルト　0

* curl例
```bash
curl --location --request GET '127.0.0.1:8080/ate_foods?skip=0&pagesize=10&type=0' \
--header 'X-Requested-With: XMLHttpRequest' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer 登録またはログインapi取得したtoken'
```

* 正常 response body
```json
{
    "ate_foods_list": [
        {
            "id": 1,
            "type": 1,
            "name": "おこめ",
            "pic": "cdn uri",
            "ate_at": "1970-01-01 15:00:00"
        },
        {
            "id": 2,
            "type": 2,
            "name": "おこめ",
            "pic": "cdn uri",
            "ate_at": "1970-01-01 15:00:00"
        },
        {
            "id": 3,
            "type": 3,
            "name": "おこめ",
            "pic": "cdn uri",
            "ate_at": "1970-01-01 15:00:00"
        },
    ],
    "total": 115   
}
```

* 異常 response body
```json
{
    "messege": "xxxxxxx"
}
```

---

<h3 id="api-detail-body-info-store">自分の体の記録</h3>

* uri:
```
  /body_info
```

* method

    post

* request body
```json
{
    "weight": 55.5,
    "fat_rate": 18.8
}
```
request body説明：
weight: 体重
fat_rate: 体脂肪率

* curl例
```bash
curl --location --request POST '127.0.0.1:8080/body_info' \
--header 'X-Requested-With: XMLHttpRequest' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer 登録またはログインapi取得したtoken' \
--data-raw '{
    "weight": 55.5,
    "fat_rate": 18.8
}'
```

* 正常 response body
```json
{
    "status": "success"
}
```
* 異常 response body
```json
{
    "messege": "xxxxxxx"
}
```

---

<h3 id="api-detail-sport-store">運動記録</h3>

* uri:
```
  /sport
```

* method

    post

* request header
Authorization: Bearer 登録またはログインapi取得したtoken

* request body
```json
{
    "name": "ジョギング",
    "minute": 30
}
```

request body説明：
name: 運動名
minute: 続ける時間、単位:分

* curl例
```bash
curl --location --request POST '127.0.0.1:8080/sport' \
--header 'X-Requested-With: XMLHttpRequest' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer 登録またはログインapi取得したtoken' \
--data-raw '{
    "name": "ジョギング",
    "minute": 30
}'
```

* 正常 response body
```json
{
    "status": "success"
}
```
* 異常 response body
```json
{
    "messege": "xxxxxxx"
}
```

---

<h3 id="api-detail-my-diary-store">自分の日記</h3>

* uri:
```
  /diary
```

* method

    post

* request header
Authorization: Bearer 登録またはログインapi取得したtoken

* request body
```json
{
    "title": "diary title",
    "content": "日記内容"
}
```

* curl例
```bash
curl --location --request POST '127.0.0.1:8080/diary' \
--header 'X-Requested-With: XMLHttpRequest' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer 登録またはログインapi取得したtoken' \
--data-raw '{
    "title": "diary title",
    "content": "日記内容"
}'
```

* 正常 response body
```json
{
    "status": "success"
}
```
* 異常 response body
```json
{
    "messege": "xxxxxxx"
}
```
---
<h3 id="api-detail-my-log-info">自分の記録ページ情報取得</h3>


* uri:
```
  /my_log_info
```

* method

    get

* request header
Authorization: Bearer 登録またはログインapi取得したtoken

* curl例
```bash
curl --location --request GET '127.0.0.1:8080/my_log_info' \
--header 'X-Requested-With: XMLHttpRequest' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer 登録またはログインapi取得したtoken'
```

* 正常 response body
```json
{
    "body_info":[
        {
            "date": "1970-01-01 00:00:00",
            "weight": 55.4,
            "fat_rate": 20
        },
        {
            "date": "1970-01-02 00:00:00",
            "weight": 55.4,
            "fat_rate": 19
        },
        {
            "date": "1970-01-03 00:00:00",
            "weight": 55.4,
            "fat_rate": 19
        },
    ],
    "sport_log": [
        {
            "date": "1970-01-02",
            "name": "ジョギング",
            "content": "日記内容"
        },
        {
            "date": "1970-01-03",
            "name": "ジョギング",
            "content": "日記内容"
        },
        {
            "date": "1970-01-04",
            "name": "ジョギング",
            "content": "日記内容"
        },
    ],
    "diary": [
        {
            "date": "1970-01-01",
            "title": "title",
            "content": "日記内容"
        },
        {
            "date": "1970-01-02",
            "title": "title2",
            "content": "日記内容222"
        }
    ]
}
```
* 異常 response body
```json
{
    "messege": "xxxxxxx"
}
```

---

* [体重体脂肪率記録取得](#api-detail-weight-fat-rate)
<h3 id="api-detail-my-log-info">自分の記録ページ情報取得</h3>
自分の記録ページに体重体脂肪率グラフの表示範囲変更の時このAPIを叩く

* uri:
```
  /weight_fat_rate
```

* method

    get

* request header
Authorization: Bearer 登録またはログインapi取得したtoken


* get params
  
  * show_by: year

  * get params説明：
  
    show_by: 指定された範囲の体重体脂肪率記録を取得、year : 年　、　month : 月、week : 週、day : 日  ディフォルト　year


* curl例
```bash
curl --location --request GET '127.0.0.1:8080/weight_fat_rate' \
--header 'X-Requested-With: XMLHttpRequest' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer 登録またはログインapi取得したtoken'
```
---
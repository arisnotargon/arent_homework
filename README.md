# arent_homework

## 使い方
docker composeでデータベースのコンテナを起動：
```bash
docker-compose up -d
```
go で起動：
```bash
go run main.go 
```

api定義をapi_solutions.mdに確認してください

完成した部分：
ユーザー登録、ユーザーログイン、ユーザー情報取得、未読お知らせ数、お知らせリスト、目標体重設定、達成率、体重体脂肪率記録取得、食事記録、食事履歴
request例はpostmanでpostman-request.jsonを導入し、参考してください


完成していない部分：
運動記録、自分の日記、自分の記録ページ情報取得、体重体脂肪率記録取得、カラムページ情報取得
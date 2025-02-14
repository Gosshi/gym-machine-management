# API 設計

## 📌 エンドポイント一覧

| HTTP メソッド | エンドポイント | 説明 |
|--------------|--------------|------|
| **GET** | `/health` | サーバーのヘルスチェック |
| **GET** | `/manufacturers` | メーカー一覧を取得 |
| **POST** | `/manufacturers` | 新しいメーカーを登録 |
| **GET** | `/manufacturers/{manufacturer_id}` | 特定のメーカーの詳細取得 |
| **POST** | `/gyms` | 新しいジムを登録 |
| **GET** | `/gyms` | ジム一覧を取得 |
| **GET** | `/gyms/{gym_id}` | 特定のジムの詳細取得 |
| **POST** | `/machines` | 新しいマシンを登録 |
| **GET** | `/machines` | マシン一覧を取得 |


## 📌 リクエスト・レスポンス例
### `POST /manufacturers`
[/manufacturers](/api/manufacturers.md)


### `POST /machines`
[/machinesについてはこちら](/api/machines.md)




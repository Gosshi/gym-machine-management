# 🏋️‍♂️ メーカー管理 API 設計

マシンの登録・管理を行うための API 設計

---

## **1️⃣ API 一覧**

| HTTP メソッド | エンドポイント                     | 説明                         |
| ------------- | ---------------------------------- | ---------------------------- |
| **POST**      | `/manufacturers`                   | 新しいメーカーを登録         |
| **GET**       | `/manufacturers`                   | すべてのメーカーの一覧を取得 |
| **GET**       | `/manufacturers/{manufacturer_id}` | 特定のメーカーの詳細を取得   |

---

## **2️⃣ API 詳細**

### **📌 `POST /manufacturers`（メーカーの登録）**

管理者が新しいメーカーを登録するエンドポイント。この API は、認証された管理者のみが利用可能 です

#### **リクエスト**

```json
{
  "name": "Technogym",
  "official_website": "https://www.technogym.com",
  "description": "A global leader in gym equipment manufacturing.",
  "logo_url": "https://example.com/logos/technogym.png",
  "contact_email": "info@technogym.com"
}
```

#### **各フィールドの説明**

| フィールド         | 必須 | 型     | 説明                     |
| ------------------ | ---- | ------ | ------------------------ |
| `name`             | 必須 | 文字列 | メーカー名               |
| `official_website` | 任意 | 文字列 | 公式サイト URL           |
| `description`      | 任意 | 文字列 | メーカーの説明           |
| `logo_url`         | 任意 | 文字列 | メーカーのロゴ画像 URL   |
| `contact_email`    | 任意 | 文字列 | 問い合わせメールアドレス |

#### **バリデーション**

| フィールド         | バリデーションルール                       |
| ------------------ | ------------------------------------------ |
| `name`             | 255 文字以内である必要があります。         |
| `official_website` | 有効な URL である必要があります。          |
| `description`      | 512 文字以内である必要があります。         |
| `logo_url`         | 有効な URL である必要があります。          |
| `contact_email`    | 有効なメールアドレスである必要があります。 |

#### **レスポンス**

メーカーごとのユニークな ID を返却します

##### **成功時**

```json
201 Created
```

```json
{
  "manufacturer_id": "V1StGXR8_Z5jdHi6B-myd"
}
```

##### **エラー時**

###### **バリデーションエラー**

```json
400 Bad Request
```

```json
{
  "error": "Invalid input",
  "details": {
    "name": "Name is required",
    "official_website": "Invalid URL format"
  }
}
```

###### **メーカーが既に存在する場合**

```json
409 Conflict
```

```json
{
  "error": "Manufacturer already exists",
  "existing_id": "x7s9k8a4h1d2f3"
}
```

###### **サーバ内部エラー**

```json
500 Internal Server Error
```

```json
{
  "error": "Internal server error",
  "message": "An unexpected error occurred"
}
```

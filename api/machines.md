# 🏋️‍♂️ マシン管理 API 設計

マシンの登録・管理を行うための API 設計。

---

## **1️⃣ API 一覧**
| HTTP メソッド | エンドポイント | 説明 |
|--------------|--------------|------|
| **POST** | `/machines` | 新しいマシン（マスターデータ）を登録 |
| **GET** | `/machines` | すべてのマシンの一覧を取得 |
| **GET** | `/machines/{machine_id}` | 特定のマシンの詳細を取得 |

---

## **2️⃣ API 詳細**
### **📌 `POST /machines`（マシンの登録）**
管理者が新しいマシンを登録するエンドポイント。この API は、認証された管理者のみが利用可能 です
#### **リクエスト**
```json
{
  "name": "レッグプレス",
  "manufacturer": "Technogym",
  "model_number": "LP-500",
  "category": "脚"
}
```
#### **レスポンス**
```json
{
  "id": id,
  "name": "レッグプレス",
  "manufacturer": "Technogym",
  "model_number": "LP-500",
  "category": "脚"
}
```         

### **📌 `GET /machines`（マシンの一覧取得）**
#### **リクエスト**
``` 


#### **レスポンス**
```json
[
  {
    "id": id,
    "name": "レッグプレス",
    "manufacturer": "Technogym",
    "model_number": "LP-500",
    "category": "脚"
  }
]
``` 

### **📌 `GET /machines/{machine_id}`（特定のマシンの詳細取得）**
#### **リクエスト**
```
GET /machines/10
```


#### **レスポンス**
```json 
{
  "id": id,
  "name": "レッグプレス",
  "manufacturer": "Technogym",
  "model_number": "LP-500",
  "category": "脚"
}
```
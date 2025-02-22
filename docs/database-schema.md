# データベース設計

## 📌 ER 図



## 📌 テーブル一覧

### gyms（ジム情報）
| カラム名 | 型 | 説明 |
|---------|------|------|
| id | SERIAL | ジムのID |
| name | VARCHAR(255) | ジムの名前 |
| location | TEXT | 住所 |
| created_at | TIMESTAMP | 作成日時 |

### machines（マシン情報）
| カラム名 | 型 | 説明 |
|---------|------|------|
| id | SERIAL | マシンのID |
| gym_id | INT | ジム ID（外部キー） |
| name | VARCHAR(255) | マシンの名前 |
| manufacturer | VARCHAR(255) | メーカー名 |
| model_number | VARCHAR(50) | 型番 |
| last_maintenance_date | DATE | 最終メンテナンス日 |
| next_maintenance_date | DATE | 次回メンテナンス予定日 |


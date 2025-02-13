# Gym Machine Management

ジムのマシン管理を効率化する Web アプリケーション。

## 📌 環境構築

### 1️⃣ 必要なツール
- [Node.js](https://nodejs.org/)（推奨: v18 以上）
- [Docker](https://www.docker.com/)（PostgreSQL の利用に必要）
- [Go](https://go.dev/)（バックエンド開発用）
- WSL2（Windows の場合推奨）

---

### 2️⃣ リポジトリをクローン
```sh
git clone https://github.com/yourusername/gym-machine-management.git
cd gym-machine-management
```

---

### 3️⃣ フロントエンドセットアップ (Next.js)
```sh
cd frontend
npm install
npm run dev
```

✅ 成功すると以下の URL でアクセス可能
```
http://localhost:3000
```

---

### 4️⃣ バックエンドセットアップ (Go + Fiber)
```sh
cd backend
go mod tidy
go run main.go  # 開発用サーバー起動

```

✅ API が動作する場合
```
http://localhost:8080
```

---

### 5️⃣ データベースセットアップ (PostgreSQL)
```sh
docker run --name gym-db -e POSTGRES_USER=gym_admin -e POSTGRES_PASSWORD=password -p 5432:5432 -d postgres
```
接続情報

- ホスト: localhost
- ポート: 5432
- ユーザー: gym_admin
- パスワード: password
- データベース名: postgres

---



# システムアーキテクチャ

## 🔹 技術スタック

- **フロントエンド**：Next.js（React, TypeScript）
- **バックエンド**：Go（Fiber）
- **データベース**：PostgreSQL
- **認証**：Amazon Cognito（JWT）
- **デプロイ**：AWS（Lambda, RDS）

## 📌 アーキテクチャ図

## 📌 コンポーネント構成

- **API サーバー**（Go + Fiber）
  - ルーティング、認証、データ管理
- **データベース**（PostgreSQL）
  - ジム・マシン・ユーザーの情報を管理
- **フロントエンド**（Next.js）
  - ジム運営者向けの管理画面

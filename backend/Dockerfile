# Stage 1: ビルド環境
FROM golang:1.22 AS build

# /go/bin にインストールされたバイナリを使うために PATH を更新
ENV PATH=$PATH:/go/bin

WORKDIR /app

# go.mod と go.sum を先にコピーして依存関係のキャッシュを利用
COPY go.mod go.sum ./
RUN go mod tidy

# swag ツールを最新でインストール（バージョン指定は不要）
RUN go install github.com/swaggo/swag/cmd/swag@latest

# アプリケーションのソースコードをコピー
COPY . .

RUN go mod tidy

# Swagger ドキュメントを生成（必要に応じて --parseDependency, --parseInternal オプションなどを付与）
RUN /go/bin/swag init --parseDependency --parseInternal

# テストを実行
RUN go test -v ./...

# アプリケーションをビルド（エントリーポイントは main.go と仮定）
RUN go build -o app main.go

# コンテナ起動時にアプリケーションを実行
CMD ["./app"]

# Go の公式イメージを使用
FROM golang:1.20

# 作業ディレクトリを設定
WORKDIR /app

# 必要なツールをインストール
RUN go install github.com/swaggo/swag/cmd/swag@latest

# go.mod と go.sum をコピー
COPY go.mod go.sum ./
RUN go mod tidy

# ソースコードをコピー
COPY . .

# Swagger ドキュメントを生成
RUN swag init


# ビルド
RUN go build -o app main.go

# 実行
CMD ["/app/app"]

FROM golang:1.20

ENV GOPRIVATE=github.com/gosshi/gym-machine-management

# 依存パッケージを取得
RUN git config --global url."git@github.com:".insteadOf "https://github.com/"
RUN go install github.com/gosshi/gym-machine-management@latest

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

CMD ["/app/app"]

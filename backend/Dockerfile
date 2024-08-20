# ベースイメージを指定
FROM golang:latest

# コンテナ内の作業ディレクトリを設定
WORKDIR /app

# ホットリロードを実現するAirをインストール
RUN go install github.com/air-verse/air@latest

# ローカルのソースコードをコンテナにコピー
COPY . .

# 必要なパッケージをインストール
RUN go mod download

# 公開予定のコンテナのポートを明示
EXPOSE 8080

# Airを使用して自動的にビルドと実行を管理する(ホットリロード)
CMD ["air"]

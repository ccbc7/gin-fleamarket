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

# Air(ホットリロード)を実行, 標準出力と標準エラー出力をログファイルに出力
CMD ["sh", "-c", "air 2>&1 | tee tmp/build-output.log"]

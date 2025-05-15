# goのイメージをDockerHubから流用する(Alpine Linux)
FROM golang:1.24.2-alpine as server-build
# Linuxパッケージ情報の最新化+gitがないのでgitを入れる
RUN apk update && apk add git
# コンテナ内の作業ディレクトリを指定
WORKDIR /app
# コンテナ内にコピー
COPY . .

# 起動時のコマンド
CMD ["go","run","main.go"]
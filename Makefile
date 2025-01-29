.PHONY: run build clean

# デフォルトのターゲット
run:
	go run cmd/main.go

# バイナリのビルド
build:
	go build -o bin/app cmd/main.go

# ビルドファイルの削除
clean:
	rm -rf bin/

# GETリクエストの実行
get:
	curl -s http://localhost:8080/users | jq '.'
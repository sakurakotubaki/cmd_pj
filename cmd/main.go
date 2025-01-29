package main

import (
	"log"
	"net/http"

	"github.com/user/cmd_pj/internal/handler"
	"github.com/user/cmd_pj/internal/repository"
	"github.com/user/cmd_pj/internal/usecase"
	"github.com/user/cmd_pj/pkg/database"
)

func main() {
	// データベース初期化
	db := database.InitDB()

	// リポジトリ、ユースケース、ハンドラーの初期化
	userRepo := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepo)
	userHandler := handler.NewUserHandler(userUseCase)

	// ルーティング
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			userHandler.CreateUser(w, r)
		case http.MethodGet:
			userHandler.GetUsers(w, r)
		case http.MethodPut:
			userHandler.UpdateUser(w, r)
		case http.MethodDelete:
			userHandler.DeleteUser(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// サーバー起動
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

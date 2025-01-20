package main

import (
	"gin-sample/internal/domain/entity"
	"gin-sample/internal/infrastructure/persistence/mysql"
	"gin-sample/internal/infrastructure/router"
	"gin-sample/internal/interface/handler"
	"gin-sample/internal/usecase"
	"log"

	gorm_mysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// db接続
	dsn := "root:test@tcp(localhost:53310)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(gorm_mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// マイグレーション
	err = db.AutoMigrate(&entity.Todo{})
	if err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}

	// 依存関係の注入
	todoRepo := mysql.NewTodoRepository(db)
	// usecase
	todoUseCase := usecase.NewTodoUsecase(todoRepo)
	// handler
	todoHander := handler.NewTodoHandler(todoUseCase)
	// router
	r := router.NewRouter(todoHander)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}

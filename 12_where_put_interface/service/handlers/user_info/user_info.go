package userinfo
import (
	"service/storage/users"
	"context"
)
func New(userRepo users.Storage) handlers.Handler {
    return func(ctx context.Context) {
        // Получение UID из запроса
        uid := 1
        user, err := userRepo.User(uid)
        if err != nil {
            // Обработка ошибки
        }
		_ = user
    }
}
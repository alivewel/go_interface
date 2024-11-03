package userinfo

import (
	"context"
	"fmt"
	"service/handlers"
	"service/storage/users"
)

type UserProvider interface {
    User(int) (users.User, error)
}

func New(userProvider UserProvider) handlers.Handler {
	return func(ctx context.Context) error {
		// Получение UID из запроса
		uid := 1
		user, err := userProvider.User(uid)
		if err != nil {
			// Обработка ошибки
			return fmt.Errorf("failed to get user: %w", err)
		}
		fmt.Printf("User: %+v\n", user)
		return nil
	}
}

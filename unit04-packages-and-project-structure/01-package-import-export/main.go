package main

import (
	"errors"
	"fmt"

	// repository
	"learn-go/unit04-packages-and-project-structure/01-package-import-export/repository"
	"learn-go/unit04-packages-and-project-structure/01-package-import-export/service"
)

func main() {
	// 1. 创建 MemoryUserRepository
	repo := repository.NewMemoryUserRepository()

	// 2. 把 repo 注入 UserService
	svc := service.NewUserService(repo)

	// 3. service.GetUser(1)
	user, err := svc.GetUser(1)

	// 4. 如果 err != nil：
	//    打印错误并 return
	if err != nil {
		fmt.Println("get user failed:", err)
		return
	}

	// 5. 打印找到的用户
	fmt.Println("found:", user.Name)

	// 6. 再调用 service.GetUser(999)
	_, err = svc.GetUser(999)

	// 7. 使用：
	// errors.Is(err, repository.ErrUserNotFound)
	// 判断 NotFound
	switch {
	case errors.Is(err, repository.ErrUserNotFound):
		fmt.Println("user 999 was not found")
	case err != nil:
		fmt.Println("unexpected error:", err)
	}
}

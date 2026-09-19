package main

import (
	"errors"
	"fmt"
	"learn-go/unit04-packages-and-project-structure/02-package-design-and-internal/internal/user"
)

func main() {
	repo := user.NewMemoryRepository()
	svc := user.NewService(repo)
	usr, err := svc.GetUser(1)
	if err != nil {
		fmt.Println("get user failed:", err)
		return
	}
	fmt.Println("found:", usr.Name)
	_, err = svc.GetUser(999)
	switch {
	case errors.Is(err, user.ErrNotFound):
		fmt.Println("user 999 was not found")
	case err != nil:
		fmt.Println("unexpected error:", err)
	}

}

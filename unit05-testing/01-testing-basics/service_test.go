package testingbasics

import (
	"errors"
	"testing"
)

func TestAdd(t *testing.T) {
	got := 1 + 1
	want := 2

	if got != want {
		t.Errorf("got %d, want %d", got, want) /*把当前测试标记为失败，但测试函数会继续往下执行。*/
	}
}

type fakeRepository struct {
	user *User
	err  error
}

func (r *fakeRepository) FindByID(id int64) (*User, error) {
	// 应该返回什么？
	return r.user, r.err
}

func TestServiceGetUserNotFound(t *testing.T) {
	// Arrange
	repo := &fakeRepository{
		user: nil,
		err:  ErrUserNotFound,
	}

	svc := NewService(repo)

	// Act
	got, err := svc.GetUser(999)

	// Assert
	if err == nil {
		t.Fatal("GetUser(999) err = nil, want an error")
	}

	if got != nil {
		t.Errorf("GetUser(999) user = %v, want nil", got)
	}

	if !errors.Is(err, ErrUserNotFound) {
		t.Errorf("GetUser(999) err = %v, want ErrUserNotFound", err)
	}
}

func TestServiceGetUserSuccess(t *testing.T) {
	// Arrange
	repo := &fakeRepository{
		user: &User{ID: 1, Name: "Alice"},
	}

	svc := NewService(repo)

	// Act
	got, err := svc.GetUser(1)

	// Assert
	if err != nil {
		// 这里应该用 Errorf 还是 Fatalf？
		t.Fatalf("GetUser(1) returned error: %v", err)
		/*因为 GetUser 出错时返回的 got 必然是 nil，继续往下走到 got.ID 就是 nil pointer panic，测试会崩而不是报清晰的错。
		Fatalf = 记失败 + 立即终止本测试函数。*/
	}

	if got == nil {
		t.Fatal("GetUser(1) returned nil user") // 无参数可格式化时用 Fatalf 对应的 Fatal
		//用 t.Fatalf，防止 got.ID panic
	}

	if got.ID != 1 {
		t.Errorf("got ID = %d, want %d", got.ID, 1)
	}

	if got.Name != "Alice" {
		t.Errorf("got Name = %q, want %q", got.Name, "Alice")
	}
}

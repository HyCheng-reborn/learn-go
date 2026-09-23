package interactiontesting

import "testing"

type fakeRepository struct {
	user *User
	err  error

	calls  int
	lastID int64
}

func (r *fakeRepository) FindByID(id int64) (*User, error) {
	// 记录一次调用
	r.calls++
	// 记录传入的 id
	r.lastID = id

	return r.user, r.err
}

func TestServiceGetUserCallsRepository(t *testing.T) {
	repo := &fakeRepository{
		user: &User{
			ID:   42,
			Name: "Alice",
		},
	}

	svc := NewService(repo)

	got, err := svc.GetUser(42)

	// 先检查 err / got
	if err != nil {
		t.Fatalf("GetUser(42) err = %v, want nil", err)
	}
	if got == nil || got.ID != 42 || got.Name != "Alice" {
		t.Fatalf("GetUser(42) got = %v, want &{42 Alice}", got)
	}

	// 然后检查交互：Service 到底有没有、以什么参数调用了 Repository
	if repo.calls != 1 {
		t.Errorf("repo.FindByID calls = %d, want 1", repo.calls)
	}
	if repo.lastID != 42 {
		t.Errorf("repo.FindByID lastID = %d, want 42", repo.lastID)
	}
}

/*
然后故意制造两个 bug

正常测试通过后，我们做两个很小的 mutation experiment。

Bug A：错误参数

临时把 Service：

s.repo.FindByID(id)

改成：

s.repo.FindByID(999)

应该出现类似：

got lastID = 999, want 42

说明：

以前只看返回值可能发现不了的问题，现在能发现。

恢复代码。

Bug B：重复调用

临时改成：

s.repo.FindByID(id)
usr, err := s.repo.FindByID(id)

应该看到：

got calls = 2, want 1

恢复代码。

这个实验非常值得做。


-------------------

验证返回结果”和“验证交互”分别在测试什么？
什么时候不应该过度 Mock?
*/

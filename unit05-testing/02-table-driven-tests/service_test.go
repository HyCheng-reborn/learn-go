package tabledriventests

import (
	"errors"
	"reflect"
	"testing"
)

/*3 个 case：
success
not found
database error

使用：
[]struct
for range
t.Run
errors.Is*/

type fakeRepository struct {
	user *User
	err  error
}

func (r *fakeRepository) FindByID(id int64) (*User, error) {
	return r.user, r.err
}

func TestServiceGetUser(t *testing.T) {
	tests := []struct {
		name     string
		repoUser *User
		repoErr  error
		wantUser *User
		wantErr  error
	}{
		{
			name:     "success",
			repoUser: &User{ID: 1, Name: "Alice"},
			repoErr:  nil,
			wantUser: &User{ID: 1, Name: "Alice"},
			wantErr:  nil,
		},
		{
			name:     "not found",
			repoUser: nil,
			repoErr:  ErrUserNotFound,
			wantUser: nil,
			wantErr:  ErrUserNotFound,
		},
		{
			name:     "database error",
			repoUser: nil,
			repoErr:  ErrDatabase,
			wantUser: nil,
			wantErr:  ErrDatabase,
		},
	}

	for _, tt := range tests { //index + 每次迭代把元素拷贝给 tt（值类型，不是引用）。
		t.Run(tt.name, func(t *testing.T) { //t.Run 在父测试里注册一个命名的执行单元。
			//每个子测试拿到自己的 *testing.T（注意闭包参数里的 t 遮蔽了外层的 t）
			svc := NewService(&fakeRepository{user: tt.repoUser, err: tt.repoErr})
			//编译期检查 fake 实现了 Repository 接口（FindByID 方法集匹配），这就是 03/04 unit 里 DI 的回报——service 不知道也不关心对面是真是假。
			got, err := svc.GetUser(1)

			if (err != nil) != (tt.wantErr != nil) {
				t.Fatalf("GetUser(1) err = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr != nil && !errors.Is(err, tt.wantErr) {
				t.Errorf("GetUser(1) err = %v, want %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.wantUser) {
				// got == tt.wantUser 比的是指针地址：两个新建的 &User{1,"Alice"} 内容相同也判不等。
				// DeepEqual 先查类型再解引用逐字段递归，所以能比内容。
				// 不用 *got != *wantUser 作通用写法：一旦 User 加入 slice/map 字段，== 直接编译失败。
				// 取舍：真实业务里对象不复杂时可改为逐字段断言，失败信息（got Name="Bob"）比整包 %v 更好定位。
				t.Errorf("GetUser(1) user = %v, want %v", got, tt.wantUser)
			}
		})
	}
}

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
			repoErr:  ErrDataBase,
			wantUser: nil,
			wantErr:  ErrDataBase,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := NewService(&fakeRepository{user: tt.repoUser, err: tt.repoErr})

			got, err := svc.GetUser(1)

			if (err != nil) != (tt.wantErr != nil) {
				t.Fatalf("GetUser(1) err = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr != nil && !errors.Is(err, tt.wantErr) {
				t.Errorf("GetUser(1) err = %v, want %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.wantUser) {
				t.Errorf("GetUser(1) user = %v, want %v", got, tt.wantUser)
			}
		})
	}
}

package store

import (
	"context"
	"go-db-labs/mysql/db"
	"go-db-labs/mysql/models"
	"go-db-labs/util"
	"testing"
)

func TestUsersStore_Create(t *testing.T) {
	ctx := context.Background()
	
	conn := db.NewConnection(ctx,db.NewDefaultConfig())
	defer db.Close(conn)
	
	usersStore := NewUsersStore(conn)

	user := models.NewDummyUser()

	err := usersStore.Create(ctx, user)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUsersStore_GetAll(t *testing.T) {
	ctx := context.Background()

	conn := db.NewConnection(ctx,db.NewDefaultConfig())
	defer db.Close(conn)

	usersStore := NewUsersStore(conn)

	user := models.NewDummyUser()

	err := usersStore.Create(ctx, user)
	if err != nil {
		t.Fatal(err)
	}

	users, err := usersStore.GetAll(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(users) == 0 {
		t.Fatal("invalid length result")
	}

}

func TestUsersStore_Get(t *testing.T) {
	ctx := context.Background()

	conn := db.NewConnection(ctx,db.NewDefaultConfig())
	defer db.Close(conn)

	usersStore := NewUsersStore(conn)

	user := models.NewDummyUser()

	err := usersStore.Create(ctx, user)
	if err != nil {
		t.Fatal(err)
	}

	found, err := usersStore.Get(ctx, user.ID)
	if err != nil {
		t.Fatal(err)
	}
	if user.Name != found.Name {
		t.Fatalf("expected=%s, actual=%s", user.Name, found.Name)
	}
	if user.Username != found.Username {
		t.Fatalf("expected=%s, actual=%s", user.Username, found.Username)
	}
	if user.Email != found.Email {
		t.Fatalf("expected=%s, actual=%s", user.Email, found.Email)
	}
	if user.Password != found.Password {
		t.Fatalf("expected=%s, actual=%s", user.Password, found.Password)
	}
	if user.Status != found.Status {
		t.Fatalf("expected=%s, actual=%s", user.Status, found.Status)
	}
	if !util.IsDateEqual(user.CreatedAt, found.CreatedAt) {
		t.Fatalf("expected=%s, actual=%s", user.CreatedAt.String(), found.CreatedAt.String())
	}
	if !util.IsDateEqual(user.UpdatedAt, found.UpdatedAt) {
		t.Fatalf("expected=%s, actual=%s", user.UpdatedAt.String(), found.UpdatedAt.String())
	}
}
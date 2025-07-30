package tests

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go-crud-api/models"
	"go-crud-api/repositories"
)

func setupTestDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=user port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database!:%v", err)
	}
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Failed to migrate:%v", err)
	}
	return db
}

func cleanDB(db *gorm.DB) {
	// ล้างข้อมูลก่อนทดสอบ
	db.Exec("TRUNCATE TABLE users RESTART IDENTITY CASCADE")
}

func TestCreateUser(t *testing.T) {
	db := setupTestDB()
	cleanDB(db)

	repo := repositories.NewUserRepository(db)
	user := &models.User{Username: "Test", Email: "test@example.com"}
	err := repo.Create(user)

	assert.Nil(t, err)
	assert.NotZero(t, user.ID)
}

func TestFindAllUser(t *testing.T) {
	db := setupTestDB()
	cleanDB(db)

	repo := repositories.NewUserRepository(db)
	user := &models.User{Username: "Test", Email: "test@example.com"}
	repo.Create(user)

	found, err := repo.FindAll()

	assert.Nil(t, err)
	assert.Len(t, found, 1)
	assert.Equal(t, user.Email, found[0].Email)
	assert.Equal(t, user.Username, found[0].Username)
}

func TestFindUserByID(t *testing.T) {
	db := setupTestDB()
	cleanDB(db)

	repo := repositories.NewUserRepository(db)
	user := &models.User{Username: "Test", Email: "test@example.com"}
	repo.Create(user)

	found, err := repo.FindByID(user.ID)

	assert.Nil(t, err)
	assert.Equal(t, user.Email, found.Email)
}

func TestUpdateUser(t *testing.T) {
	db := setupTestDB()
	cleanDB(db)

	repo := repositories.NewUserRepository(db)
	user := &models.User{Username: "Test", Email: "test@example.com"}
	repo.Create(user)

	user.Email = "newemail@example.com"
	err := repo.Update(user)

	assert.Nil(t, err)

	updated, _ := repo.FindByID(user.ID)
	assert.Equal(t, "newemail@example.com", updated.Email)
}

func TestDeleteUser(t *testing.T) {
	db := setupTestDB()
	cleanDB(db)

	repo := repositories.NewUserRepository(db)
	user := &models.User{Username: "Test", Email: "test@example.com"}
	repo.Create(user)

	err := repo.Delete(user)
	assert.Nil(t, err)

	_, err = repo.FindByID(user.ID)
	assert.NotNil(t, err) // ควร error เพราะ user ถูกลบแล้ว
}

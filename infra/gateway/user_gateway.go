package gateway

import (
	"database/sql"
	"gaudy_code/domain/model"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type UserRepository struct{}

func (repo *UserRepository) FindUser(userID int64) (*model.User, error) {
	db, err := sql.Open("mysql", "root:my-secret-pw@tcp(localhost:3306)/game_db")
	if err != nil {
		return nil, err
	}
	defer db.Close()
	row, err := db.Query("SELECT id, login_date, coins, is_login_rewarded FROM user where id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var user model.User
	if row.Next() {
		var tmpLoginDate []uint8
		err := row.Scan(&user.ID, &tmpLoginDate, &user.Coin, &user.IsLoginRewaredNeed)
		if err != nil {
			return nil, err
		}
		loginDate, err := time.Parse("2006-01-02 15:04:05", string(tmpLoginDate))
		if err != nil {
			log.Fatalf("Failed to parse login_date: %v", err)
			return nil, err
		}
		user.LastLoginDate = loginDate
	}
	if err = row.Err(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) UpdateUser(user model.User) error {
	db, err := sql.Open("mysql", "root:my-secret-pw@tcp(localhost:3306)/game_db")
	if err != nil {
		return err
	}
	query := `
		UPDATE user
		SET
			login_date = ?,
			coins = ?,
			is_login_rewarded = ?
		WHERE id = ?`

	if _, err = db.Exec(query, user.LastLoginDate, user.Coin, user.IsLoginRewaredNeed, user.ID); err != nil {
		log.Printf("Failed to update user: %v", err)
		return err
	}
	return nil
}

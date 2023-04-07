package gateway

import (
	"database/sql"
	"gaudy_code/domain/model"
)

type MonsterRepository struct {
	// 何らかのDB
}

func (repo *MonsterRepository) FindMyMonsters(userID int64) ([]model.MyMonster, error) {
	db, err := sql.Open("mysql", "root:my-secret-pw@tcp(localhost:3306)/game_db")
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query("SELECT id, monster_id, user_id, name, level FROM monster where user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var myMonsters []model.MyMonster
	for rows.Next() {
		var myMonster model.MyMonster
		if err := rows.Scan(&myMonster.ID, &myMonster.MonsterID, &myMonster.UserID, &myMonster.Name, &myMonster.Level); err != nil {
			return nil, err
		}
		myMonsters = append(myMonsters, myMonster)
	}
	return myMonsters, nil
}

func (repo *MonsterRepository) FindDefeatedMonsters(userID int64) ([]model.EnemyMonster, error) {
	db, err := sql.Open("mysql", "root:my-secret-pw@tcp(localhost:3306)/game_db")
	if err != nil {
		return nil, err
	}
	defer db.Close()
	query := "SELECT id, user_id, monster_id, count, defeated_at, last_reset_at FROM user_monster_defeated where user_id = ?"
	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var enemyMonsters []model.EnemyMonster
	for rows.Next() {
		var enemyMonster model.EnemyMonster
		if err := rows.Scan(&enemyMonster.ID, &enemyMonster.UserID, &enemyMonster.MonsterID, &enemyMonster.Count,
			&enemyMonster.DefeatedAt, &enemyMonster.LastResetAts); err != nil {
			return nil, err
		}
		enemyMonsters = append(enemyMonsters, enemyMonster)
	}
	return enemyMonsters, nil
}

func (repo *MonsterRepository) UpdateMyMonster(monster model.MyMonster) error {
	db, err := sql.Open("mysql", "root:my-secret-pw@tcp(localhost:3306)/game_db")
	if err != nil {
		return err
	}
	defer db.Close()
	query := "INSERT INTO monster (monster_id, user_id, name, level) VALUES (?, ?, ?, ?)"
	result, err := db.Exec(query, monster.MonsterID, monster.UserID, monster.Name, monster.Level)
	if err != nil {
		return err
	}
	if _, err := result.LastInsertId(); err != nil {
		return err
	}

	return nil

}

func (repo *MonsterRepository) UpdateEnemyMonster(monster model.EnemyMonster) error {
	db, err := sql.Open("mysql", "root:my-secret-pw@tcp(localhost:3306)/game_db")
	if err != nil {
		return err
	}
	defer db.Close()
	query := "INSERT INTO user_monster_defeated (user_id, monster_id, count, defeated_at, last_reset_at) VALUES (?, ?, ?, ?, ?)"
	result, err := db.Exec(query, monster.UserID, monster.MonsterID, monster.Count, monster.DefeatedAt, monster.LastResetAts)
	if err != nil {
		return err
	}
	if _, err := result.LastInsertId(); err != nil {
		return err
	}
	return nil
}

package players

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

type PostgresRepository struct {
	conn *sqlx.DB
}

func NewPostgresRepository(dbClient *sqlx.DB) Repository {
	r := PostgresRepository{
		conn: dbClient,
	}
	var ifm Repository = &r
	return ifm
}

func (r *PostgresRepository) Get(id int64) (Player, error) {
	query := `SELECT id, name, wins, losses, time_played, achievements FROM player WHERE id = :id`
	params := map[string]interface{}{
		"id": id,
	}
	rows, err := r.conn.NamedQuery(query, params)
	if err != nil {
		return Player{}, err
	}
	defer rows.Close()

	if rows.Next() {
		player := Player{}
		err = rows.Scan(&player.ID, &player.Name, &player.Wins, &player.Losses, &player.Time_played, &player.Achievements)
		if err != nil {
			return Player{}, err
		}
		return player, nil
	}
	return Player{}, nil
}

// Create creates a new player in the repository
func (r *PostgresRepository) Create(player Player) (int64, error) {
	query := `INSERT INTO player (id, name, wins, losses, time_played, achievements)
		VALUES (DEFAULT, :name, :wins, :losses, :time_played, :achievements) RETURNING id`
	params := map[string]interface{}{
		"name":         player.Name,
		"wins":         player.Wins,
		"losses":       player.Losses,
		"time_played":  player.Time_played,
		"achievements": player.Achievements,
	}
	rows, err := r.conn.NamedQuery(query, params)
	if err != nil {
		return -1, errors.New("Couldn't create the Player " + err.Error())
	}
	defer rows.Close()

	var id int64
	if rows.Next() {
		rows.Scan(&id)
	} else {
		return -1, errors.New("Error creating Player " + err.Error())
	}
	return id, nil
}

// Update updates a player in the repository
func (r *PostgresRepository) Update(player Player) error {
	query := `UPDATE player 
		SET  name = :name, wins = :wins, losses = :losses, time_played = :time_played, achievements = :achievements 
		WHERE id = :id`
	params := map[string]interface{}{
		"id":           player.ID,
		"name":         player.Name,
		"wins":         player.Wins,
		"losses":       player.Losses,
		"time_played":  player.Time_played,
		"achievements": player.Achievements,
	}
	res, err := r.conn.NamedExec(query, params)
	if err != nil {
		return errors.New("Couldn't update the Player:" + err.Error())
	}
	i, err := res.RowsAffected()
	if err != nil {
		return errors.New("Error with the affected rows:" + err.Error())
	}
	if i != 1 {
		return errors.New("No row inserted (or multiple row inserted) instead of 1 row")
	}
	return nil
}

// Delete deletes a player in the repository
func (r *PostgresRepository) Delete(id int64) error {
	query := `DELETE FROM player WHERE id = :id`
	params := map[string]interface{}{
		"id": id,
	}
	res, err := r.conn.NamedExec(query, params)
	if err != nil {
		return errors.New("Couldn't delete the Player:" + err.Error())
	}
	i, err := res.RowsAffected()
	if err != nil {
		return errors.New("Error with the affected rows:" + err.Error())
	}
	if i != 1 {
		return errors.New("No row inserted (or multiple row inserted) instead of 1 row")
	}
	return nil
}

// GetAll returns all players in the repository
func (r *PostgresRepository) GetAll() (map[int64]Player, error) {
	query := `SELECT id, name, wins, losses, time_played, achievements FROM player`
	rows, err := r.conn.Query(query)
	if err != nil {
		return nil, errors.New("Couldn't retrieve players " + err.Error())
	}
	defer rows.Close()

	players := make(map[int64]Player, 0)
	for rows.Next() {
		player := Player{}
		err := rows.Scan(&player.ID, &player.Name, &player.Wins, &player.Losses, &player.Time_played, &player.Achievements)
		if err != nil {
			return nil, errors.New("Couldn't scan the retrieved data: " + err.Error())
		}
		players[player.ID] = player
	}
	return players, nil
}

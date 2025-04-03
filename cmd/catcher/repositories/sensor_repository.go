package repositories

import (
	"database/sql"
	"log"
)

type CanReceiveSensorData interface {
	Exists(sensorId int) bool
}

type SensorRepository struct {
	db *sql.DB
}

func NewSensorRepository(db *sql.DB) SensorRepository {
	return SensorRepository{
		db: db,
	}
}

func (repos SensorRepository) Exists(sensorId int) bool {
	query := `SELECT count(*) FROM sensors WHERE id=$1`
	rows, err := repos.db.Query(query, sensorId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	countResult := CountDbResult{}

	rows.Next()
	err = rows.Scan(&countResult.count)
	if err != nil {
		log.Fatal(err)
	}

	return countResult.count > 0
}

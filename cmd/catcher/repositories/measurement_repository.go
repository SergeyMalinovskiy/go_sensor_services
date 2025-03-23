package repositories

import (
	"database/sql"
	"log"
	"time"
)

type MeasurementData struct {
	value      float32
	valueExtra string
	sensorId   int
	date       time.Time
}

type MeasurementSaver interface {
	Save(data MeasurementData) int
}

type CountDbResult struct {
	count int
}

type MeasurementRepository struct {
	db *sql.DB
}

func (MeasurementRepository) Save(data MeasurementData) int {

	return 0
}

func NewMeasurementRepository(db *sql.DB) MeasurementRepository {
	return MeasurementRepository{
		db: db,
	}
}

func (repos MeasurementRepository) insertData(data MeasurementData) {
	query := `INSERT INTO measurements (value, value_extra, sensor_id, date, created_at) VALUES ($1, $2, $3, $4, $5)`
	_, err := repos.db.Exec(
		query,
		data.value,
		data.valueExtra,
		data.sensorId,
		data.date.String(),
		time.Now(),
	)
	if err != nil {
		log.Fatal(err)
	}
}

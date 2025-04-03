package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	contract "github.com/SergeyMalinovskiy/growther/libs/contract/golang"
	"time"
)

type MeasurementSaver interface {
	Save(data contract.MeasurementData) (int, error)
}

type CountDbResult struct {
	count int
}

type MeasurementRepository struct {
	db *sql.DB
}

func (repos MeasurementRepository) Save(
	data contract.MeasurementData,
) (int, error) {
	err := repos.insertData(data)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("Measurement repository: save error: %v", err))
	}

	return 1, nil
}

func NewMeasurementRepository(db *sql.DB) MeasurementRepository {
	return MeasurementRepository{
		db: db,
	}
}

func (repos MeasurementRepository) insertData(data contract.MeasurementData) error {
	query := `INSERT INTO measurements (value, value_extra, sensor_id, date, created_at) VALUES ($1, $2, $3, $4, $5)`
	_, err := repos.db.Exec(
		query,
		data.Value,
		sql.NullString{String: data.ValueExtra, Valid: data.ValueExtra != ""},
		data.SensorId,
		data.Datetime,
		time.Now(),
	)

	return err
}

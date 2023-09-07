package repositories

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/ffelipelimao/booking/internal/entities/place"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func setupSuite(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	assert.Nil(t, err)

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS place (
			id VARCHAR(255) PRIMARY KEY,
    		capacity INT,
    		assigned_seating BOOLEAN,
    		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	assert.Nil(t, err)

	return db
}

func tearDownSuite(db *sql.DB) {
	db.Close()
	db.Exec("drop table place")
}

func TestPlaceRepository_Save(t *testing.T) {
	db := setupSuite(t)

	defer tearDownSuite(db)

	tests := []struct {
		name          string
		place         *place.Place
		expectedError error
	}{
		{
			name: "Should insert a new place",
			place: &place.Place{
				ID:              "test-id",
				Capacity:        100,
				AssignedSeating: true,
				CreatedAt:       time.Now(),
			},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewPlaceRepository(db)
			err := repo.Save(context.Background(), tt.place)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}

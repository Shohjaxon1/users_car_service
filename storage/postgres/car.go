package postgres

import (
	"database/sql"
	pb "student-service/genproto/user_service"
	"time"

	"github.com/spf13/cast"
)

// Create inserts a new user into the database
func (r *userRepo) CreateCar(car *pb.Car) (*pb.Car, error) {
	query := `
		INSERT INTO car (id, model, color, year, price, created_at) 
		VALUES ($1, $2, $3, $4, $5, $6) 
		RETURNING id, model, color, year, price, created_at
	`
	err := r.db.QueryRow(query, car.Id, car.Model, car.Color, car.Year, car.Price, time.Now()).Scan(
		&car.Id,
		&car.Model,
		&car.Color,
		&car.Year,
		&car.Price,
		&car.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return car, nil
}

// GetById retrieves a user by their ID
func (r *userRepo) GetCarById(req *pb.GetByIdRequest) (*pb.Car, error) {
	query := "SELECT id, model, color, year, price, created_at FROM car WHERE id = $1"
	row := r.db.QueryRow(query, req.Id)

	car := &pb.Car{}
	err := row.Scan(
		&car.Id,
		&car.Model,
		&car.Color,
		&car.Year,
		&car.Price,
		&car.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return car, nil
}

// Update updates a user's details in the database
func (r *userRepo) UpdateCar(req *pb.Car) (*pb.Car, error) {
	query := `
        UPDATE car
        SET model = $2, color = $3, year = $4, price = $5, updated_at = $6
        WHERE id = $1
        RETURNING id, model, color, year, price, updated_at
    `
	row := r.db.QueryRow(query, req.Id, req.Model, req.Color, req.Year, req.Price, time.Now())
	updatedCar := &pb.Car{}
	err := row.Scan(
		&updatedCar.Id,
		&updatedCar.Model,
		&updatedCar.Color,
		&updatedCar.Year,
		&updatedCar.Price,
		&updatedCar.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return updatedCar, nil
}

// Delete removes a user from the database
func (r *userRepo) DeleteCar(req *pb.GetByIdRequest) (*pb.Car, error) {
	query := "DELETE FROM car WHERE id = $1 RETURNING id, model, color, year, price, created_at, updated_at"
	row := r.db.QueryRow(query, req.Id)
	var updatedAt sql.NullTime
	deletedCar := &pb.Car{}
	err := row.Scan(
		&deletedCar.Id,
		&deletedCar.Model,
		&deletedCar.Color,
		&deletedCar.Year,
		&deletedCar.Price,
		&deletedCar.CreatedAt,
		&updatedAt,
	)
	if err != nil {
		return nil, err
	}
	if updatedAt.Valid {
		deletedCar.UpdatedAt = updatedAt.Time.String()
	} else {
		deletedCar.UpdatedAt = ""
	}
	return deletedCar, nil
}

func (r *userRepo) GetAllCars(req *pb.GetAllRequest) (*pb.AllCars, error) {
	intLimit := cast.ToInt(req.Limit)
	intPage := cast.ToInt(req.Page)
	offset := (intPage - 1) * intLimit

	query := `
        SELECT id, model, color, year, price, created_at, updated_at
        FROM car
        LIMIT $1 OFFSET $2
    `
	rows, err := r.db.Query(query, intLimit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cars := []*pb.Car{}

	for rows.Next() {
		var updatedAt sql.NullTime
		car := &pb.Car{}
		err := rows.Scan(
			&car.Id,
			&car.Model,
			&car.Color,
			&car.Year,
			&car.Price,
			&car.CreatedAt,
			&updatedAt,
		)
		if err != nil {
			return nil, err
		}
		if updatedAt.Valid {
			car.UpdatedAt = updatedAt.Time.String()
		} else {
			car.UpdatedAt = ""
		}
		cars = append(cars, car)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &pb.AllCars{Cars: cars}, nil
}

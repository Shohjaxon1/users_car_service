package postgres

import (
	"database/sql"
	pb "student-service/genproto/user_service"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/cast"
)

type userRepo struct {
	db *sqlx.DB
}

// NewUserRepo creates a new userRepocmd/main.go
func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{db: db}
}

// Create inserts a new user into the database
func (r *userRepo) Create(user *pb.User) (*pb.User, error) {
	query := `
		INSERT INTO users (id, full_name, username, phone_number, created_at) 
		VALUES ($1, $2, $3, $4, $5) 
		RETURNING id, full_name, username, phone_number, created_at
	`
	err := r.db.QueryRow(query, user.Id, user.FullName, user.Username, user.PhoneNumber, time.Now()).Scan(
		&user.Id,
		&user.FullName,
		&user.Username,
		&user.PhoneNumber,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepo) GiveCarToUser(user *pb.User) (*pb.User, error) {
	query := `
		INSERT INTO users (id, full_name, username, phone_number, created_at) 
		VALUES ($1, $2, $3, $4, $5) 
		RETURNING id, full_name, username, phone_number, created_at
	`
	err := r.db.QueryRow(query, user.Id, user.FullName, user.Username, user.PhoneNumber, time.Now()).Scan(
		&user.Id,
		&user.FullName,
		&user.Username,
		&user.PhoneNumber,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetById retrieves a user by their ID
func (r *userRepo) GetById(req *pb.GetByIdRequest) (*pb.User, error) {
	query := "SELECT id, full_name, username, phone_number, created_at FROM users WHERE id = $1"
	row := r.db.QueryRow(query, req.Id)

	user := &pb.User{}
	err := row.Scan(
		&user.Id,
		&user.FullName,
		&user.Username,
		&user.PhoneNumber,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepo) GetByPhoneNumber(req *pb.GetByPhoneNumberRequest) (*pb.User, error) {
	query := "SELECT id, full_name, username, phone_number, created_at FROM users WHERE phone_number = $1"
	row := r.db.QueryRow(query, req.PhoneNumber)

	user := &pb.User{}
	err := row.Scan(
		&user.Id,
		&user.FullName,
		&user.Username,
		&user.PhoneNumber,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (r *userRepo) GetUserWithCars(req *pb.GetByIdRequest) (*pb.AllUsersWithCar, error) {
	query := `
        SELECT 
            u.id,
            u.full_name, 
            u.username, 
            u.phone_number,
            u.created_at, 
            c.id, 
            c.model, 
            c.color, 
            c.year,
            c.price
        FROM 
            car_and_user cu 
        JOIN 
            "users" u ON cu.user_id = u.id 
        JOIN 
            car c ON cu.car_id = c.id 
        WHERE 
            u.id = $1 
        ORDER BY 
            u.full_name;`

	rows, err := r.db.Query(query, req.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := &pb.AllUsersWithCar{
		Users: []*pb.User{},
		Cars:  []*pb.Car{},
	}

	for rows.Next() {
		user := &pb.User{}
		car := &pb.Car{}
		err := rows.Scan(
			&user.Id,
			&user.FullName,
			&user.Username,
			&user.PhoneNumber,
			&user.CreatedAt,

			&car.Id,
			&car.Model,
			&car.Color,
			&car.Year,
			&car.Price,
		)
		if err != nil {
			return nil, err
		}
		users.Users = append(users.Users, user)
		users.Cars = append(users.Cars, car)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// Update updates a user's details in the database
func (r *userRepo) Update(req *pb.User) (*pb.User, error) {
	query := `
        UPDATE users
        SET full_name = $2, username = $3, phone_number = $4, updated_at = $5
        WHERE id = $1
        RETURNING id, full_name, username, phone_number, updated_at
    `
	row := r.db.QueryRow(query, req.Id, req.FullName, req.Username, req.PhoneNumber, time.Now())
	updatedUser := &pb.User{}
	err := row.Scan(
		&updatedUser.Id,
		&updatedUser.FullName,
		&updatedUser.Username,
		&updatedUser.PhoneNumber,
		&updatedUser.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}

// Delete removes a user from the database
func (r *userRepo) Delete(req *pb.GetByIdRequest) (*pb.User, error) {
	query := "DELETE FROM users WHERE id = $1 RETURNING id, full_name, username, phone_number, created_at, updated_at"
	row := r.db.QueryRow(query, req.Id)
	var updatedAt sql.NullTime
	deletedUser := &pb.User{}
	err := row.Scan(
		&deletedUser.Id,
		&deletedUser.FullName,
		&deletedUser.Username,
		&deletedUser.PhoneNumber,
		&deletedUser.CreatedAt,
		&updatedAt,
	)
	if err != nil {
		return nil, err
	}
	if updatedAt.Valid {
		deletedUser.UpdatedAt = updatedAt.Time.String()
	} else {
		deletedUser.UpdatedAt = ""
	}
	return deletedUser, nil
}

func (r *userRepo) GetAllUserWithCars(req *pb.GetAllRequest) (*pb.AllUsersWithCar, error) {
	intLimit := cast.ToInt(req.Limit)
	intPage := cast.ToInt(req.Page)
	offset := (intPage - 1) * intLimit

	query := `
        SELECT 
            u.id,
            u.full_name, 
            u.username, 
            u.phone_number,
            u.created_at, 
			u.updated_at,
            c.id, 
            c.model, 
            c.color, 
            c.year,
            c.price
		FROM 
            car_and_user cu 
        JOIN 
            "users" u ON cu.user_id = u.id 
        JOIN 
            car c ON cu.car_id = c.id 
        ORDER BY 
            u.full_name
        LIMIT $1 OFFSET $2;
    `

	rows, err := r.db.Query(query, intLimit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := &pb.AllUsersWithCar{
		Users: []*pb.User{},
		Cars:  []*pb.Car{},
	}

	for rows.Next() {
		var updatedAt sql.NullTime
		user := &pb.User{}
		car := &pb.Car{}
		err := rows.Scan(
			&user.Id,
			&user.FullName,
			&user.Username,
			&user.PhoneNumber,
			&user.CreatedAt,
			&updatedAt,
			&car.Id,
			&car.Model,
			&car.Color,
			&car.Year,
			&car.Price,
		)
		if err != nil {
			return nil, err
		}
		if updatedAt.Valid {
			user.UpdatedAt = updatedAt.Time.String()
		} else {
			user.UpdatedAt = ""
		}
		users.Users = append(users.Users, user)
		users.Cars = append(users.Cars, car)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepo) GetAll(req *pb.GetAllRequest) (*pb.AllUsers, error) {
	intLimit := cast.ToInt(req.Limit)
	intPage := cast.ToInt(req.Page)
	offset := (intPage - 1) * intLimit

	query := `
        SELECT id, full_name, username, phone_number, created_at, updated_at
        FROM users
        LIMIT $1 OFFSET $2
    `
	rows, err := r.db.Query(query, intLimit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []*pb.User{}

	for rows.Next() {
		var updatedAt sql.NullTime
		user := &pb.User{}
		err := rows.Scan(
			&user.Id,
			&user.FullName,
			&user.Username,
			&user.PhoneNumber,
			&user.CreatedAt,
			&updatedAt,
		)
		if err != nil {
			return nil, err
		}
		if updatedAt.Valid {
			user.UpdatedAt = updatedAt.Time.String()
		} else {
			user.UpdatedAt = ""
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &pb.AllUsers{Users: users}, nil
}

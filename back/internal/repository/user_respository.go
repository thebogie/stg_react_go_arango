// internal/repository/user_repository.go
package repository

import (
	genmodel "back/graph/generated/model"
	connection "back/internal/db"
	"context"
)

type UserRepository interface {
	FindByEmail(email string) (*genmodel.User, error)
	Save(user *genmodel.User) error
	FindUserByID(id string) (*genmodel.User, error)
}

type userRepository struct {
	dbconn *connection.DatabaseConnection
}

func NewUserRepository(dbconn *connection.DatabaseConnection) UserRepository {
	return &userRepository{
		dbconn: dbconn}
}

// Implement the UserRepository methods
func (r *userRepository) FindByEmail(email string) (*genmodel.User, error) {
	query := "FOR d IN player FILTER d.email == @email RETURN d"
	bindVars := map[string]interface{}{
		"email": email,
	}
	//TODO: pass context from top level?
	ctx := context.Background()
	cursor, err := r.dbconn.Db.Query(ctx, query, bindVars)
	if err != nil {

	}
	defer cursor.Close()
	return &genmodel.User{}, nil

	// Implement the logic to find a user by email
	return nil, nil
}

func (r *userRepository) FindUserByID(userid string) (*genmodel.User, error) {
	query := "FOR d IN player FILTER d.email == @email RETURN d"
	bindVars := map[string]interface{}{
		"email": userid,
	}
	//TODO: pass context from top level?
	ctx := context.Background()
	cursor, err := r.dbconn.Db.Query(ctx, query, bindVars)
	if err != nil {

	}
	defer cursor.Close()
	return &genmodel.User{}, nil

	// Implement the logic to find a user by email
	return nil, nil
}

func (r *userRepository) Save(user *genmodel.User) error {
	// Implement the logic to save a user
	return nil
}

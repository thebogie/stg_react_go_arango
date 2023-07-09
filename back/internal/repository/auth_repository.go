package repository

import (
	genmodel "back/graph/generated/model"
	connection "back/internal/db"
	"context"
)

type AuthRepository interface {
	FindByEmail(email string) (*genmodel.User, error)
}

type authRepository struct {
	dbconn *connection.DatabaseConnection
}

func NewAuthRepository(dbconn *connection.DatabaseConnection) AuthRepository {
	return &authRepository{
		dbconn: dbconn}
}

func (r *authRepository) FindByEmail(email string) (*genmodel.User, error) {
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

	var retuser = &genmodel.User{}
	_, err = cursor.ReadDocument(ctx, &retuser)
	if err != nil {

	}
	return retuser, err

}

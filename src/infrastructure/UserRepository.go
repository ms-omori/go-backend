package infrastructure

import (
	"database/sql"
	"fmt"

	"ms-omori/domain"
)

type UserRepository interface {
	GetUserFromName(name string) (*domain.User, error)
	ListUsers() (*[]domain.User, error)
	GetOrganization(id int) (*domain.Organization, error)
}

type UserRepositoryImpl struct{}

func NewUserRepository() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (ur *UserRepositoryImpl) GetUser(name string) (*domain.User, error) {
	row, err := db.Query(fmt.Sprintf("SELECT id, name, organization_id FROM user WHERE name = %v", name))
	if err != nil {
		return nil, err
	}
	defer row.Close()
	row.Next()

	return ur.setUserInfo(row, err)
}

func (ur *UserRepositoryImpl) ListUsers() (*[]domain.User, error) {
	rows, err := db.Query("select * from user")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []domain.User
	for rows.Next() {
		user, err := ur.setUserInfo(rows, err)
		if err != nil {
			return nil, err
		}
		users = append(users, *user)
	}
	return &users, nil
}

func (ur *UserRepositoryImpl) GetOrganization(id int) (*domain.Organization, error) {
	row, err := db.Query("SELECT * FROM organization WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	row.Next()
	organization := domain.Organization{}
	if err := row.Scan(&organization); err != nil {
		return nil, err
	}
	return &organization, nil
}

func (ur *UserRepositoryImpl) setUserInfo(row *sql.Rows, err error) (*domain.User, error) {
	var id int
	var userName string
	var organizationID int
	if err := row.Scan(&id, &userName, &organizationID); err != nil {
		return nil, err
	}
	organization, err := ur.GetOrganization(organizationID)
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:           id,
		Name:         userName,
		Organization: *organization,
	}, nil
}

package repository

import (
	"fmt"
	"ubiquitous-payment/webshop/model"
)

func (repo *Repository) GetPrivilegesByUserID(id uint) (*[]string, error) {
	var privileges []string
	if err := repo.RelationalDatabase.Raw("select p.name from privileges p, role_privileges rp "+
		"where rp.role_id in (select r.id from roles r, user_roles ur where ur.user_id = ? and ur.role_id = r.id) "+
		"and p.id = rp.privilege_id", id).Scan(&privileges).Error; err != nil {
		return nil, err
	}
	return &privileges, nil
}

func (repo *Repository) GetUserByEmail(email string) (*model.User, error) {
	user := &model.User{}
	if err := repo.RelationalDatabase.Preload("Roles").Table("users").First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *Repository) GetUserById(id uint) (*model.User, error) {
	user := &model.User{}
	if err := repo.RelationalDatabase.Table("users").First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *Repository) CreateProfile(profile *model.Profile) error {
	result := repo.RelationalDatabase.Create(profile)
	if result.RowsAffected == 0 {
		return fmt.Errorf("ProfileVertex not created")
	}
	fmt.Println("Profile created")
	return nil
}

func (repo *Repository) CreateUser(user *model.User) error {
	result := repo.RelationalDatabase.Create(user)
	if result.RowsAffected == 0 {
		return fmt.Errorf("user not created")
	}
	fmt.Println("User created")
	return nil
}

func (repo *Repository) GetRoleByName(name string) (*model.Role, error) {
	role := &model.Role{}
	if err := repo.RelationalDatabase.Table("roles").First(&role, "name = ?", name).Error; err != nil {
		return nil, err
	}
	return role, nil
}

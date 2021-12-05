package repository

import (
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

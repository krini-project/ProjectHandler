package persistence

import "log"

// ListUserPorjects List all projects that are associated with a user
func (handler *DatabaseHandler) ListUserProjects(userID string) ([]*Project, error) {
	user := User{
		UserID: userID,
	}

	if err := handler.Database.Preload("Projects").First(&user).Error; err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return user.Projects, nil
}

// CreateNewUser Creates a new User
func (handler *DatabaseHandler) CreateNewUser(userID string, userMail string) error {
	user := User{
		UserID:   userID,
		UserMail: userMail,
	}

	if err := handler.Database.Create(&user).Error; err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

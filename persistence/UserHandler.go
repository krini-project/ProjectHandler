package persistence

import (
	"context"
	"log"

	"github.com/krini-project/ProjectHandler/models"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// ListUserProjects List all projects that are associated with a user
func (handler *DatabaseHandler) ListUserProjects(userID string) ([]*models.Project, error) {
	user, err := models.FindUser(context.TODO(), handler.DB, userID)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	usersProjectUsers, err := user.ProjectUsers(qm.Load(qm.Rels(models.ProjectUserRels.ProjectsProject))).All(context.TODO(), handler.DB)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var projects []*models.Project

	for _, usersProjectUser := range usersProjectUsers {
		project := usersProjectUser.R.ProjectsProject
		projects = append(projects, project)
	}

	return projects, nil
}

// CreateNewUser Creates a new User
func (handler *DatabaseHandler) CreateNewUser(userID string, userMail string) error {
	exists, err := models.Users(qm.Where("UserID=?", userID)).Exists(context.TODO(), handler.DB)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	if !exists {
		user := models.User{
			UserID: userID,
		}

		err := user.Insert(context.TODO(), handler.DB, boil.Infer())
		if err != nil {
			log.Println(err.Error())
			return err
		}
	}

	return nil
}

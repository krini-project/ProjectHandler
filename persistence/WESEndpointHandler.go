package persistence

import (
	"context"
	"log"

	"github.com/krini-project/ProjectHandler/models"
	"github.com/volatiletech/sqlboiler/boil"
)

func (handler *DatabaseHandler) AddWESEndpoint(name string, link string, ownerProject string) error {
	wesEndpoint := models.WESEndpoint{
		Name:    name,
		Link:    link,
		OwnerID: ownerProject,
	}

	if err := wesEndpoint.Insert(context.TODO(), handler.DB, boil.Infer()); err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (handler *DatabaseHandler) ListWESEndpoints(userid string, projectid string) ([]*models.WESEndpoint, error) {
	project, err := models.FindProject(context.TODO(), handler.DB, projectid)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	wesEndpointsSlice, err := project.OwnerWESEndpoints().All(context.TODO(), handler.DB)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var wesEndpoints []*models.WESEndpoint
	for _, endpoint := range wesEndpointsSlice {
		wesEndpoints = append(wesEndpoints, endpoint)
	}

	return wesEndpoints, nil
}

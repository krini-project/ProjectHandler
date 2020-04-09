package persistence

import (
	"context"
	"log"

	"github.com/krini-project/ProjectHandler/models"
	"github.com/volatiletech/sqlboiler/boil"
)

const AddUser = "add_users"
const RawDataAccess = "raw_data_access"
const AddDatasets = "add_dataset"
const AddWorkflows = "add_workflow"
const SetWESEndpoints = "set_WES_endpoints"
const ExecuteWorkflow = "execute_workflow"

//const computeResources = "compute_resources"

const AdminRole = "admin"
const MaintainerRole = "maintainer"
const InternRole = "intern"

var defaultRights []string
var defaultRoles = make(map[string][]string)

func (handler *DatabaseHandler) CreateDefaultRoles() error {
	defaultRoles[AdminRole] = []string{AddUser, RawDataAccess, AddDatasets, AddWorkflows, SetWESEndpoints, ExecuteWorkflow}
	defaultRoles[MaintainerRole] = []string{RawDataAccess, AddDatasets, AddDatasets, SetWESEndpoints, ExecuteWorkflow}
	defaultRoles[InternRole] = []string{ExecuteWorkflow}

	defaultRights := []string{AddUser, RawDataAccess, AddDatasets, AddWorkflows, SetWESEndpoints, ExecuteWorkflow}

	for _, right := range defaultRights {
		righStruct := models.Right{
			RightName: right,
		}

		err := righStruct.Insert(context.TODO(), handler.DB, boil.Infer())
		if err != nil {
			log.Println(err.Error())
		}
	}

	rightSlice, err := models.Rights().All(context.TODO(), handler.DB)
	if err != nil {
		log.Println(err.Error())
	}

	rightMap := make(map[string]*models.Right)

	for _, right := range rightSlice {
		rightMap[right.RightName] = right
	}

	for role := range defaultRoles {
		roleStruct := models.Role{
			RoleName: role,
		}

		err := roleStruct.Insert(context.TODO(), handler.DB, boil.Infer())
		if err != nil {
			log.Println(err.Error())
		}
	}

	roleSlice, err := models.Roles().All(context.TODO(), handler.DB)
	if err != nil {
		log.Println(err.Error())
	}

	for _, role := range roleSlice {
		roleRights := defaultRoles[role.RoleName]
		var roleRightsStruct []*models.Right
		for _, roleRight := range roleRights {
			roleRightsStruct = append(roleRightsStruct, rightMap[roleRight])
		}
		role.AddRightsRightIDRights(context.TODO(), handler.DB, false, roleRightsStruct...)
	}

	return nil
}

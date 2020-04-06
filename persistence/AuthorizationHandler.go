package persistence

import "log"

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

var defaultRights = []*Right{
	&Right{
		Name: AddUser,
	},
	&Right{
		Name: RawDataAccess,
	},
	&Right{
		Name: AddDatasets,
	},
	&Right{
		Name: AddWorkflows,
	},
	&Right{
		Name: SetWESEndpoints,
	},
	&Right{
		Name: ExecuteWorkflow,
	},
}

var defaultRoles = []*Role{
	&Role{
		Name:   AdminRole,
		Rights: defaultRights,
	},
	&Role{
		Name: MaintainerRole,
		Rights: []*Right{
			&Right{
				Name: RawDataAccess,
			},
			&Right{
				Name: AddDatasets,
			},
			&Right{
				Name: AddWorkflows,
			},
			&Right{
				Name: SetWESEndpoints,
			},
			&Right{
				Name: ExecuteWorkflow,
			},
		},
	},
	&Role{
		Name: InternRole,
		Rights: []*Right{
			&Right{
				Name: ExecuteWorkflow,
			},
		},
	},
}

func (handler *DatabaseHandler) CreateDefaultRights() error {
	for _, right := range defaultRights {
		if err := handler.Database.Create(right).Error; err != nil {
			log.Println(err.Error())
			return err
		}
	}

	for _, role := range defaultRoles {
		for _, roleRight := range role.Rights {
			if err := handler.Database.First(roleRight).Error; err != nil {
				log.Println(err.Error())
				return err
			}
		}
		if err := handler.Database.Create(role).Error; err != nil {
			log.Println(err.Error())
			return err
		}
	}

	return nil
}

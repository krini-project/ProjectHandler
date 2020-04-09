// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Datasets", testDatasets)
	t.Run("ProjectUsers", testProjectUsers)
	t.Run("Projects", testProjects)
	t.Run("Rights", testRights)
	t.Run("Roles", testRoles)
	t.Run("Users", testUsers)
	t.Run("WESEndpoints", testWESEndpoints)
	t.Run("Workflowruns", testWorkflowruns)
	t.Run("Workflows", testWorkflows)
}

func TestDelete(t *testing.T) {
	t.Run("Datasets", testDatasetsDelete)
	t.Run("ProjectUsers", testProjectUsersDelete)
	t.Run("Projects", testProjectsDelete)
	t.Run("Rights", testRightsDelete)
	t.Run("Roles", testRolesDelete)
	t.Run("Users", testUsersDelete)
	t.Run("WESEndpoints", testWESEndpointsDelete)
	t.Run("Workflowruns", testWorkflowrunsDelete)
	t.Run("Workflows", testWorkflowsDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Datasets", testDatasetsQueryDeleteAll)
	t.Run("ProjectUsers", testProjectUsersQueryDeleteAll)
	t.Run("Projects", testProjectsQueryDeleteAll)
	t.Run("Rights", testRightsQueryDeleteAll)
	t.Run("Roles", testRolesQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
	t.Run("WESEndpoints", testWESEndpointsQueryDeleteAll)
	t.Run("Workflowruns", testWorkflowrunsQueryDeleteAll)
	t.Run("Workflows", testWorkflowsQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Datasets", testDatasetsSliceDeleteAll)
	t.Run("ProjectUsers", testProjectUsersSliceDeleteAll)
	t.Run("Projects", testProjectsSliceDeleteAll)
	t.Run("Rights", testRightsSliceDeleteAll)
	t.Run("Roles", testRolesSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
	t.Run("WESEndpoints", testWESEndpointsSliceDeleteAll)
	t.Run("Workflowruns", testWorkflowrunsSliceDeleteAll)
	t.Run("Workflows", testWorkflowsSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Datasets", testDatasetsExists)
	t.Run("ProjectUsers", testProjectUsersExists)
	t.Run("Projects", testProjectsExists)
	t.Run("Rights", testRightsExists)
	t.Run("Roles", testRolesExists)
	t.Run("Users", testUsersExists)
	t.Run("WESEndpoints", testWESEndpointsExists)
	t.Run("Workflowruns", testWorkflowrunsExists)
	t.Run("Workflows", testWorkflowsExists)
}

func TestFind(t *testing.T) {
	t.Run("Datasets", testDatasetsFind)
	t.Run("ProjectUsers", testProjectUsersFind)
	t.Run("Projects", testProjectsFind)
	t.Run("Rights", testRightsFind)
	t.Run("Roles", testRolesFind)
	t.Run("Users", testUsersFind)
	t.Run("WESEndpoints", testWESEndpointsFind)
	t.Run("Workflowruns", testWorkflowrunsFind)
	t.Run("Workflows", testWorkflowsFind)
}

func TestBind(t *testing.T) {
	t.Run("Datasets", testDatasetsBind)
	t.Run("ProjectUsers", testProjectUsersBind)
	t.Run("Projects", testProjectsBind)
	t.Run("Rights", testRightsBind)
	t.Run("Roles", testRolesBind)
	t.Run("Users", testUsersBind)
	t.Run("WESEndpoints", testWESEndpointsBind)
	t.Run("Workflowruns", testWorkflowrunsBind)
	t.Run("Workflows", testWorkflowsBind)
}

func TestOne(t *testing.T) {
	t.Run("Datasets", testDatasetsOne)
	t.Run("ProjectUsers", testProjectUsersOne)
	t.Run("Projects", testProjectsOne)
	t.Run("Rights", testRightsOne)
	t.Run("Roles", testRolesOne)
	t.Run("Users", testUsersOne)
	t.Run("WESEndpoints", testWESEndpointsOne)
	t.Run("Workflowruns", testWorkflowrunsOne)
	t.Run("Workflows", testWorkflowsOne)
}

func TestAll(t *testing.T) {
	t.Run("Datasets", testDatasetsAll)
	t.Run("ProjectUsers", testProjectUsersAll)
	t.Run("Projects", testProjectsAll)
	t.Run("Rights", testRightsAll)
	t.Run("Roles", testRolesAll)
	t.Run("Users", testUsersAll)
	t.Run("WESEndpoints", testWESEndpointsAll)
	t.Run("Workflowruns", testWorkflowrunsAll)
	t.Run("Workflows", testWorkflowsAll)
}

func TestCount(t *testing.T) {
	t.Run("Datasets", testDatasetsCount)
	t.Run("ProjectUsers", testProjectUsersCount)
	t.Run("Projects", testProjectsCount)
	t.Run("Rights", testRightsCount)
	t.Run("Roles", testRolesCount)
	t.Run("Users", testUsersCount)
	t.Run("WESEndpoints", testWESEndpointsCount)
	t.Run("Workflowruns", testWorkflowrunsCount)
	t.Run("Workflows", testWorkflowsCount)
}

func TestHooks(t *testing.T) {
	t.Run("Datasets", testDatasetsHooks)
	t.Run("ProjectUsers", testProjectUsersHooks)
	t.Run("Projects", testProjectsHooks)
	t.Run("Rights", testRightsHooks)
	t.Run("Roles", testRolesHooks)
	t.Run("Users", testUsersHooks)
	t.Run("WESEndpoints", testWESEndpointsHooks)
	t.Run("Workflowruns", testWorkflowrunsHooks)
	t.Run("Workflows", testWorkflowsHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Datasets", testDatasetsInsert)
	t.Run("Datasets", testDatasetsInsertWhitelist)
	t.Run("ProjectUsers", testProjectUsersInsert)
	t.Run("ProjectUsers", testProjectUsersInsertWhitelist)
	t.Run("Projects", testProjectsInsert)
	t.Run("Projects", testProjectsInsertWhitelist)
	t.Run("Rights", testRightsInsert)
	t.Run("Rights", testRightsInsertWhitelist)
	t.Run("Roles", testRolesInsert)
	t.Run("Roles", testRolesInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
	t.Run("WESEndpoints", testWESEndpointsInsert)
	t.Run("WESEndpoints", testWESEndpointsInsertWhitelist)
	t.Run("Workflowruns", testWorkflowrunsInsert)
	t.Run("Workflowruns", testWorkflowrunsInsertWhitelist)
	t.Run("Workflows", testWorkflowsInsert)
	t.Run("Workflows", testWorkflowsInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("DatasetToProjectUsingProject", testDatasetToOneProjectUsingProject)
	t.Run("ProjectUserToProjectUsingProjectsProject", testProjectUserToOneProjectUsingProjectsProject)
	t.Run("ProjectUserToRoleUsingRole", testProjectUserToOneRoleUsingRole)
	t.Run("ProjectUserToUserUsingUser", testProjectUserToOneUserUsingUser)
	t.Run("WESEndpointToProjectUsingOwner", testWESEndpointToOneProjectUsingOwner)
	t.Run("WorkflowrunToProjectUsingProject", testWorkflowrunToOneProjectUsingProject)
	t.Run("WorkflowrunToWESEndpointUsingWESEndpoint", testWorkflowrunToOneWESEndpointUsingWESEndpoint)
	t.Run("WorkflowrunToWorkflowUsingWorkflow", testWorkflowrunToOneWorkflowUsingWorkflow)
	t.Run("WorkflowToProjectUsingProject", testWorkflowToOneProjectUsingProject)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("ProjectToDatasets", testProjectToManyDatasets)
	t.Run("ProjectToProjectsProjectProjectUsers", testProjectToManyProjectsProjectProjectUsers)
	t.Run("ProjectToOwnerWESEndpoints", testProjectToManyOwnerWESEndpoints)
	t.Run("ProjectToWorkflowruns", testProjectToManyWorkflowruns)
	t.Run("ProjectToWorkflows", testProjectToManyWorkflows)
	t.Run("RightToRolesRoleIDRoles", testRightToManyRolesRoleIDRoles)
	t.Run("RoleToProjectUsers", testRoleToManyProjectUsers)
	t.Run("RoleToRightsRightIDRights", testRoleToManyRightsRightIDRights)
	t.Run("UserToProjectUsers", testUserToManyProjectUsers)
	t.Run("WESEndpointToWESEndpointWorkflowruns", testWESEndpointToManyWESEndpointWorkflowruns)
	t.Run("WorkflowToWorkflowruns", testWorkflowToManyWorkflowruns)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("DatasetToProjectUsingDatasets", testDatasetToOneSetOpProjectUsingProject)
	t.Run("ProjectUserToProjectUsingProjectsProjectProjectUsers", testProjectUserToOneSetOpProjectUsingProjectsProject)
	t.Run("ProjectUserToRoleUsingProjectUsers", testProjectUserToOneSetOpRoleUsingRole)
	t.Run("ProjectUserToUserUsingProjectUsers", testProjectUserToOneSetOpUserUsingUser)
	t.Run("WESEndpointToProjectUsingOwnerWESEndpoints", testWESEndpointToOneSetOpProjectUsingOwner)
	t.Run("WorkflowrunToProjectUsingWorkflowruns", testWorkflowrunToOneSetOpProjectUsingProject)
	t.Run("WorkflowrunToWESEndpointUsingWESEndpointWorkflowruns", testWorkflowrunToOneSetOpWESEndpointUsingWESEndpoint)
	t.Run("WorkflowrunToWorkflowUsingWorkflowruns", testWorkflowrunToOneSetOpWorkflowUsingWorkflow)
	t.Run("WorkflowToProjectUsingWorkflows", testWorkflowToOneSetOpProjectUsingProject)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {
	t.Run("DatasetToProjectUsingDatasets", testDatasetToOneRemoveOpProjectUsingProject)
	t.Run("WorkflowrunToProjectUsingWorkflowruns", testWorkflowrunToOneRemoveOpProjectUsingProject)
	t.Run("WorkflowrunToWESEndpointUsingWESEndpointWorkflowruns", testWorkflowrunToOneRemoveOpWESEndpointUsingWESEndpoint)
	t.Run("WorkflowrunToWorkflowUsingWorkflowruns", testWorkflowrunToOneRemoveOpWorkflowUsingWorkflow)
}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("ProjectToDatasets", testProjectToManyAddOpDatasets)
	t.Run("ProjectToProjectsProjectProjectUsers", testProjectToManyAddOpProjectsProjectProjectUsers)
	t.Run("ProjectToOwnerWESEndpoints", testProjectToManyAddOpOwnerWESEndpoints)
	t.Run("ProjectToWorkflowruns", testProjectToManyAddOpWorkflowruns)
	t.Run("ProjectToWorkflows", testProjectToManyAddOpWorkflows)
	t.Run("RightToRolesRoleIDRoles", testRightToManyAddOpRolesRoleIDRoles)
	t.Run("RoleToProjectUsers", testRoleToManyAddOpProjectUsers)
	t.Run("RoleToRightsRightIDRights", testRoleToManyAddOpRightsRightIDRights)
	t.Run("UserToProjectUsers", testUserToManyAddOpProjectUsers)
	t.Run("WESEndpointToWESEndpointWorkflowruns", testWESEndpointToManyAddOpWESEndpointWorkflowruns)
	t.Run("WorkflowToWorkflowruns", testWorkflowToManyAddOpWorkflowruns)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("ProjectToDatasets", testProjectToManySetOpDatasets)
	t.Run("ProjectToWorkflowruns", testProjectToManySetOpWorkflowruns)
	t.Run("RightToRolesRoleIDRoles", testRightToManySetOpRolesRoleIDRoles)
	t.Run("RoleToRightsRightIDRights", testRoleToManySetOpRightsRightIDRights)
	t.Run("WESEndpointToWESEndpointWorkflowruns", testWESEndpointToManySetOpWESEndpointWorkflowruns)
	t.Run("WorkflowToWorkflowruns", testWorkflowToManySetOpWorkflowruns)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("ProjectToDatasets", testProjectToManyRemoveOpDatasets)
	t.Run("ProjectToWorkflowruns", testProjectToManyRemoveOpWorkflowruns)
	t.Run("RightToRolesRoleIDRoles", testRightToManyRemoveOpRolesRoleIDRoles)
	t.Run("RoleToRightsRightIDRights", testRoleToManyRemoveOpRightsRightIDRights)
	t.Run("WESEndpointToWESEndpointWorkflowruns", testWESEndpointToManyRemoveOpWESEndpointWorkflowruns)
	t.Run("WorkflowToWorkflowruns", testWorkflowToManyRemoveOpWorkflowruns)
}

func TestReload(t *testing.T) {
	t.Run("Datasets", testDatasetsReload)
	t.Run("ProjectUsers", testProjectUsersReload)
	t.Run("Projects", testProjectsReload)
	t.Run("Rights", testRightsReload)
	t.Run("Roles", testRolesReload)
	t.Run("Users", testUsersReload)
	t.Run("WESEndpoints", testWESEndpointsReload)
	t.Run("Workflowruns", testWorkflowrunsReload)
	t.Run("Workflows", testWorkflowsReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Datasets", testDatasetsReloadAll)
	t.Run("ProjectUsers", testProjectUsersReloadAll)
	t.Run("Projects", testProjectsReloadAll)
	t.Run("Rights", testRightsReloadAll)
	t.Run("Roles", testRolesReloadAll)
	t.Run("Users", testUsersReloadAll)
	t.Run("WESEndpoints", testWESEndpointsReloadAll)
	t.Run("Workflowruns", testWorkflowrunsReloadAll)
	t.Run("Workflows", testWorkflowsReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Datasets", testDatasetsSelect)
	t.Run("ProjectUsers", testProjectUsersSelect)
	t.Run("Projects", testProjectsSelect)
	t.Run("Rights", testRightsSelect)
	t.Run("Roles", testRolesSelect)
	t.Run("Users", testUsersSelect)
	t.Run("WESEndpoints", testWESEndpointsSelect)
	t.Run("Workflowruns", testWorkflowrunsSelect)
	t.Run("Workflows", testWorkflowsSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Datasets", testDatasetsUpdate)
	t.Run("ProjectUsers", testProjectUsersUpdate)
	t.Run("Projects", testProjectsUpdate)
	t.Run("Rights", testRightsUpdate)
	t.Run("Roles", testRolesUpdate)
	t.Run("Users", testUsersUpdate)
	t.Run("WESEndpoints", testWESEndpointsUpdate)
	t.Run("Workflowruns", testWorkflowrunsUpdate)
	t.Run("Workflows", testWorkflowsUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Datasets", testDatasetsSliceUpdateAll)
	t.Run("ProjectUsers", testProjectUsersSliceUpdateAll)
	t.Run("Projects", testProjectsSliceUpdateAll)
	t.Run("Rights", testRightsSliceUpdateAll)
	t.Run("Roles", testRolesSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
	t.Run("WESEndpoints", testWESEndpointsSliceUpdateAll)
	t.Run("Workflowruns", testWorkflowrunsSliceUpdateAll)
	t.Run("Workflows", testWorkflowsSliceUpdateAll)
}

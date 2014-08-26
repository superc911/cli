package resources

import (
	"fmt"

	"github.com/cloudfoundry/cli/cf/models"
)

type ServicePlanResource struct {
	Resource
	Entity ServicePlanEntity
}

type ServicePlanEntity struct {
	Name                string
	Free                bool
	Public              bool
	Active              bool
	Description         string                  `json:"description"`
	ServiceOfferingGuid string                  `json:"service_guid"`
	ServiceOffering     ServiceOfferingResource `json:"service"`
}

type ServicePlanDescription struct {
	ServiceLabel    string
	ServicePlanName string
	ServiceProvider string
}

func (resource ServicePlanResource) ToFields() models.ServicePlan {
	return models.ServicePlan{
		Guid:                resource.Metadata.Guid,
		Name:                resource.Entity.Name,
		Free:                resource.Entity.Free,
		Description:         resource.Entity.Description,
		Public:              resource.Entity.Public,
		Active:              resource.Entity.Active,
		ServiceOfferingGuid: resource.Entity.ServiceOfferingGuid,
	}
}

func (planDesc ServicePlanDescription) String() string {
	if planDesc.ServiceProvider == "" {
		return fmt.Sprintf("%s %s", planDesc.ServiceLabel, planDesc.ServicePlanName) // v2 plan
	} else {
		return fmt.Sprintf("%s %s %s", planDesc.ServiceLabel, planDesc.ServiceProvider, planDesc.ServicePlanName) // v1 plan
	}
}

type ServiceMigrateV1ToV2Response struct {
	ChangedCount int `json:"changed_count"`
}

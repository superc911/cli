package resources

import "github.com/cloudfoundry/cli/cf/models"

type PaginatedServiceOfferingResources struct {
	Resources []ServiceOfferingResource
}

type ServiceOfferingResource struct {
	Resource
	Entity ServiceOfferingEntity
}

type ServiceOfferingEntity struct {
	Label            string                `json:"label"`
	Version          string                `json:"version"`
	Description      string                `json:"description"`
	DocumentationUrl string                `json:"documentation_url"`
	Provider         string                `json:"provider"`
	BrokerGuid       string                `json:"service_broker_guid"`
	ServicePlans     []ServicePlanResource `json:"service_plans"`
}

func (resource ServiceOfferingResource) ToFields() models.ServiceOffering {
	return models.ServiceOffering{
		Label:            resource.Entity.Label,
		Version:          resource.Entity.Version,
		Provider:         resource.Entity.Provider,
		Description:      resource.Entity.Description,
		BrokerGuid:       resource.Entity.BrokerGuid,
		Guid:             resource.Metadata.Guid,
		DocumentationUrl: resource.Entity.DocumentationUrl,
	}
}

func (resource ServiceOfferingResource) ToModel() models.ServiceOffering {
	offering := resource.ToFields()
	for _, p := range resource.Entity.ServicePlans {
		servicePlan := models.ServicePlan{}
		servicePlan.Name = p.Entity.Name
		servicePlan.Guid = p.Metadata.Guid
		offering.Plans = append(offering.Plans, servicePlan)
	}
	return offering
}

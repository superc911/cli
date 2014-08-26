package resources

import "github.com/cloudfoundry/cli/cf/models"

type PaginatedServiceInstanceResources struct {
	TotalResults int `json:"total_results"`
	Resources    []ServiceInstanceResource
}

type ServiceInstanceResource struct {
	Resource
	Entity ServiceInstanceEntity
}

type ServiceInstanceEntity struct {
	Name            string
	ServiceBindings []ServiceBindingResource `json:"service_bindings"`
	ServicePlan     ServicePlanResource      `json:"service_plan"`
}

func (resource ServiceInstanceResource) ToFields() (serviceInstance models.ServiceInstance) {
	serviceInstance.Guid = resource.Metadata.Guid
	serviceInstance.Name = resource.Entity.Name
	return
}

func (resource ServiceInstanceResource) ToModel() models.ServiceInstance {
	serviceInstance := resource.ToFields()

	serviceInstance.ServicePlan = resource.Entity.ServicePlan.ToFields()

	serviceInstance.ServiceBindings = []models.ServiceBindingFields{}
	for _, bindingResource := range resource.Entity.ServiceBindings {
		serviceInstance.ServiceBindings = append(serviceInstance.ServiceBindings, bindingResource.ToFields())
	}
	return serviceInstance
}

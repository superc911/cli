package models

type ServiceInstance struct {
	Guid             string
	Name             string
	SysLogDrainUrl   string
	ApplicationNames []string
	Params           map[string]interface{}
	ServiceBindings  []ServiceBindingFields
	ServicePlan      ServicePlan
	ServiceOffering  ServiceOffering
}

func (inst ServiceInstance) IsUserProvided() bool {
	return inst.ServicePlan.Guid == ""
}

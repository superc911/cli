package models

type SpaceFields struct {
	Guid string
	Name string
}

type Space struct {
	SpaceFields
	Organization     OrganizationFields
	Applications     []ApplicationFields
	ServiceInstances []ServiceInstance
	Domains          []DomainFields
	SecurityGroups   []SecurityGroupFields
	SpaceQuotaGuid   string
}

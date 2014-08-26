package maker

import "github.com/cloudfoundry/cli/cf/models"

func NewServiceOffering(label string) models.ServiceOffering {
	return models.ServiceOffering{
		Label:       label,
		Guid:        serviceOfferingGuid(),
		Description: "some service description",
	}
}

var serviceOfferingGuid = guidGenerator("services")

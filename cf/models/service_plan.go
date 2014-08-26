package models

type ServicePlan struct {
	Guid                string
	Name                string
	Free                bool
	Public              bool
	Description         string
	Active              bool
	ServiceOfferingGuid string
	OrgNames            []string
	ServiceOffering     ServiceOffering
}

func (servicePlan ServicePlan) OrgHasVisibility(orgName string) bool {
	if servicePlan.Public {
		return true
	}
	for _, org := range servicePlan.OrgNames {
		if org == orgName {
			return true
		}
	}
	return false
}

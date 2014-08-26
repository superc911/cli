package models

type ServiceOfferings []ServiceOffering

type ServiceOffering struct {
	Guid             string
	BrokerGuid       string
	Label            string
	Provider         string
	Version          string
	Description      string
	DocumentationUrl string
	Plans            []ServicePlan
}

func (s ServiceOfferings) Len() int {
	return len(s)
}

func (s ServiceOfferings) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ServiceOfferings) Less(i, j int) bool {
	return s[i].Label < s[j].Label
}

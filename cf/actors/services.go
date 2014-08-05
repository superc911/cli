package actors

import (
	"github.com/cloudfoundry/cli/cf/actors/broker_builder"
	"github.com/cloudfoundry/cli/cf/actors/service_builder"
	"github.com/cloudfoundry/cli/cf/api"
	"github.com/cloudfoundry/cli/cf/models"
)

type ServiceActor interface {
	FilterBrokers(brokerFlag string, serviceFlag string, orgFlag string) ([]models.ServiceBroker, error)
}

type ServiceHandler struct {
	orgRepo        api.OrganizationRepository
	brokerBuilder  broker_builder.BrokerBuilder
	serviceBuilder service_builder.ServiceBuilder
}

func NewServiceHandler(org api.OrganizationRepository, brokerBuilder broker_builder.BrokerBuilder, serviceBuilder service_builder.ServiceBuilder) ServiceHandler {
	return ServiceHandler{
		orgRepo:        org,
		brokerBuilder:  brokerBuilder,
		serviceBuilder: serviceBuilder,
	}
}

func (actor ServiceHandler) FilterBrokers(brokerFlag string, serviceFlag string, orgFlag string) ([]models.ServiceBroker, error) {
	if orgFlag == "" {
		return actor.getServiceBrokers(brokerFlag, serviceFlag)
	} else {
		err := actor.checkForOrgExistence(orgFlag)
		if err != nil {
			return nil, err
		}
		return actor.buildBrokersVisibleFromOrg(brokerFlag, serviceFlag, orgFlag)
	}
}

func (actor ServiceHandler) checkForOrgExistence(orgName string) error {
	_, err := actor.orgRepo.FindByName(orgName)
	return err
}

func (actor ServiceHandler) getServiceBrokers(brokerName string, serviceName string) ([]models.ServiceBroker, error) {
	if serviceName != "" {
		brokers, err := actor.brokerBuilder.GetBrokerWithSpecifiedService(serviceName)
		if err != nil {
			return nil, err
		}

		if brokerName != "" {
			if brokers != nil && brokers[0].Name != brokerName {
				return nil, nil
			}
		}
		return brokers, nil
	}

	if brokerName != "" && serviceName == "" {
		return actor.brokerBuilder.GetBrokerWithAllServices(brokerName)
	}

	return actor.brokerBuilder.GetAllServiceBrokers()
}

func (actor ServiceHandler) buildBrokersVisibleFromOrg(brokerFlag string, serviceFlag string, orgFlag string) ([]models.ServiceBroker, error) {
	if serviceFlag != "" && brokerFlag != "" {
		services, err := actor.serviceBuilder.GetServiceVisibleToOrg(serviceFlag, orgFlag)
		if err != nil {
			return nil, err
		}
		return actor.brokerBuilder.GetSpecificBrokerForServices(brokerFlag, services)
	}

	if serviceFlag != "" && brokerFlag == "" {
		services, err := actor.serviceBuilder.GetServiceVisibleToOrg(serviceFlag, orgFlag)
		if err != nil {
			return nil, err
		}
		return actor.brokerBuilder.GetBrokersForServices(services)
	}

	if serviceFlag == "" && brokerFlag != "" {
		services, err := actor.serviceBuilder.GetServicesVisibleToOrg(orgFlag)
		if err != nil {
			return nil, err
		}
		return actor.brokerBuilder.GetSpecificBrokerForServices(brokerFlag, services)
	}

	if serviceFlag == "" && brokerFlag == "" {
		services, err := actor.serviceBuilder.GetServicesVisibleToOrg(orgFlag)
		if err != nil {
			return nil, err
		}
		return actor.brokerBuilder.GetBrokersForServices(services)
	}

	return nil, nil
}

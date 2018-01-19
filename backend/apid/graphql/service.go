package graphql

import (
	"github.com/sensu/sensu-go/backend/apid/graphql/schema"
	"github.com/sensu/sensu-go/backend/messaging"
	"github.com/sensu/sensu-go/backend/store"
	"github.com/sensu/sensu-go/graphql"
)

// ServiceConfig describes values required to instantiate service.
type ServiceConfig struct {
	Store store.Store
	Bus   messaging.MessageBus
}

// NewService instantiates new GraphQL service
func NewService(cfg ServiceConfig) (*graphql.Service, error) {
	svc := graphql.NewService()
	store := cfg.Store
	nodeResolver := newNodeResolver(store)

	// Register types
	schema.RegisterAsset(svc, &assetImpl{})
	schema.RegisterDeleteRecordInput(svc)
	schema.RegisterDeleteRecordPayload(svc, &deleteRecordPayload{})
	schema.RegisterEvent(svc, &eventImpl{})
	schema.RegisterHandler(svc, newHandlerImpl(store))
	schema.RegisterHandlerSocket(svc, &handlerSocketImpl{})
	schema.RegisterQuery(svc, newQueryImpl(store, nodeResolver))
	schema.RegisterMutation(svc, newMutationImpl(store))
	schema.RegisterMutator(svc, &mutatorImpl{})
	schema.RegisterNamespace(svc, &namespaceImpl{})
	schema.RegisterNode(svc, &nodeImpl{nodeResolver})
	schema.RegisterNamespaceInput(svc)
	schema.RegisterPageInfo(svc, &pageInfoImpl{})
	schema.RegisterViewer(svc, newViewerImpl(store, cfg.Bus))
	schema.RegisterSchema(svc)

	// Register check types
	schema.RegisterCheck(svc, &checkImpl{})
	schema.RegisterCheckConfig(svc, newCheckCfgImpl(store))
	schema.RegisterCheckConfigConnection(svc, &schema.CheckConfigConnectionAliases{})
	schema.RegisterCheckConfigEdge(svc, &schema.CheckConfigEdgeAliases{})
	schema.RegisterCheckHistory(svc, &checkHistoryImpl{})
	schema.RegisterCheckConfigInputs(svc)
	schema.RegisterCreateCheckInput(svc)
	schema.RegisterCreateCheckPayload(svc, &checkMutationPayload{})
	schema.RegisterUpdateCheckInput(svc)
	schema.RegisterUpdateCheckPayload(svc, &checkMutationPayload{})

	// Register entity types
	schema.RegisterEntity(svc, newEntityImpl(store))
	schema.RegisterEntityConnection(svc, &schema.EntityConnectionAliases{})
	schema.RegisterEntityEdge(svc, &schema.EntityEdgeAliases{})
	schema.RegisterDeregistration(svc, &deregistrationImpl{})
	schema.RegisterNetwork(svc, &networkImpl{})
	schema.RegisterNetworkInterface(svc, &networkInterfaceImpl{})
	schema.RegisterSystem(svc, &systemImpl{})

	// Register event types
	schema.RegisterEvent(svc, &eventImpl{})
	schema.RegisterEventConnection(svc, &schema.EventConnectionAliases{})
	schema.RegisterEventEdge(svc, &schema.EventEdgeAliases{})

	// Register hook types
	schema.RegisterHook(svc, &hookImpl{})
	schema.RegisterHookConfig(svc, &hookCfgImpl{})
	schema.RegisterHookList(svc, &hookListImpl{})

	// Register time window
	schema.RegisterTimeWindowDays(svc, &timeWindowDaysImpl{})
	schema.RegisterTimeWindowWhen(svc, &timeWindowWhenImpl{})
	schema.RegisterTimeWindowTimeRange(svc, &timeWindowTimeRangeImpl{})

	// Register user types
	schema.RegisterRole(svc, &roleImpl{})
	schema.RegisterRule(svc, &ruleImpl{})
	schema.RegisterRuleResource(svc)
	schema.RegisterRulePermission(svc)
	schema.RegisterUser(svc, &userImpl{})

	err := svc.Regenerate()
	return svc, err
}

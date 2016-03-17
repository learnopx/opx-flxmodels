package models

var ConfigObjectMap = map[string]ConfigObj{
	"Vlan":                                             &Vlan{},      // created before auto YANG
	"IPv4Route":                                        &IPv4Route{}, // created before auto YANG
	"IPv4RouteState":                                   &IPv4RouteState{},
	"IPv4EventState":                                   &IPv4EventState{},
	"IPv4Intf":                                         &IPv4Intf{}, // created before auto YANG
	"LogicalIntfConfig":                                &LogicalIntfConfig{},
	"LogicalIntfState":                                 &LogicalIntfState{},
	//"ArpConfig":                                        &ArpConfig{},         // created before auto YANG
	//"ArpEntry":                                         &ArpEntry{},          // created before auto YANG
	"AggregationLacpConfig":                            &AggregationLacpConfig{},
	"EthernetConfig":                                   &EthernetConfig{},
	"AggregationLacpMemberStateCounters":               &AggregationLacpMemberStateCounters{},
	"AggregationLacpState":                             &AggregationLacpState{},
	"BfdSessionConfig":                                 &BfdSessionConfig{},
	"BfdSessionState":                                  &BfdSessionState{},
	"BfdGlobalConfig":                                  &BfdGlobalConfig{},
	"BfdGlobalState":                                   &BfdGlobalState{},
	"BfdIntfConfig":                                    &BfdIntfConfig{},
	"BfdIntfState":                                     &BfdIntfState{},
	"DhcpRelayHostDhcpState":                           &DhcpRelayHostDhcpState{},
	"DhcpRelayIntfServerState":                         &DhcpRelayIntfServerState{},
	"DhcpRelayIntfConfig":                              &DhcpRelayIntfConfig{},
	"DhcpRelayIntfState":                               &DhcpRelayIntfState{},
	"DhcpRelayGlobalConfig":                            &DhcpRelayGlobalConfig{},
	"PolicyDefinitionConfig":                           &PolicyDefinitionConfig{},
	"PolicyDefinitionState":                            &PolicyDefinitionState{},
	"PolicyStmtConfig":                                 &PolicyStmtConfig{},
	"PolicyStmtState":                                  &PolicyStmtState{},
	"PolicyConditionConfig":                            &PolicyConditionConfig{},
	"PolicyConditionState":                             &PolicyConditionState{},
	"PolicyActionConfig":                               &PolicyActionConfig{},
	"PolicyActionState":                                &PolicyActionState{},
	"RouteDistanceState":                               &RouteDistanceState{},
	"PortConfig":                                       &PortConfig{},
	"PortState":                                        &PortState{},
	"UserConfig":                                       &UserConfig{},
	"IPV4AddressBlock":                                 &IPV4AddressBlock{},
	"UserState":                                        &UserState{},
	"Dot1dStpPortEntryConfig":                          &Dot1dStpPortEntryConfig{},
	"Dot1dStpBridgeState":                              &Dot1dStpBridgeState{},
	"Dot1dStpBridgeConfig":                             &Dot1dStpBridgeConfig{},
	"Dot1dStpPortEntryStateCountersFsmStatesPortTimer": &Dot1dStpPortEntryStateCountersFsmStatesPortTimer{},
}

package models

import (
	"encoding/json"
	"fmt"
)

type PolicyPrefix struct {
	IpPrefix string
	/*
            Defines a range for the masklength, or 'exact' if
            the prefix has an exact length.

            Example: 10.3.192.0/21 through 10.3.192.0/24 would be
            expressed as prefix: 10.3.192.0/21,
            masklength-range: 21..24.

            Example: 10.3.192.0/21 would be expressed as
            prefix: 10.3.192.0/21,
            masklength-range: exact	
	*/
	MaskLengthRange string
}

type PolicyPrefixSet struct {
	BaseObj
	PrefixSetName string   `SNAPROUTE: "KEY"`
    /*
     List of prefix expressions that are part of the set
	*/
	IpPrefixList [] PolicyPrefix
}

func (obj PolicyPrefixSet) UnmarshalObject(body []byte) (ConfigObj, error) {
	var policyPrefixSet PolicyPrefixSet
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &policyPrefixSet); err != nil {
			fmt.Println("### Trouble in unmarshaling PolicyPrefixSet from Json", body)
		}
	}
	return policyPrefixSet, err
}


type PolicyDstIpMatchPrefixSetCondition struct {
	//yang_name: prefix-set class: leaf
	PrefixSet string
	//yang_name: match-set-options class: leaf
    Prefix PolicyPrefix 
}

type PolicyNeighbor struct {
	Address [] string
}

type PolicyNeighborSet struct {
	BaseObj
	NeighborSetName string `SNAPROUTE: "KEY"`
	NeihgborSet []PolicyNeighbor
}

func (obj PolicyNeighborSet) UnmarshalObject(body []byte) (ConfigObj, error) {
	var policyNeighborSet PolicyNeighborSet
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &policyNeighborSet); err != nil {
			fmt.Println("### Trouble in unmarshaling PolicyNeighborSet from Json", body)
		}
	}
	return policyNeighborSet, err
}

type PolicyMatchNeighborSetCondition struct {
	//yang_name: neighbor-set class: leaf
	NeighborSet string
	Neihgbor PolicyNeighbor
	//yang_name: match-set-options class: leaf
	MatchSetOptions int32
}

type PolicyTag struct {
	value [] uint32
}

type PolicyTagSet struct {
	BaseObj
	TagSetName string               `SNAPROUTE: "KEY"`
	TagList []PolicyTag
}

func (obj PolicyTagSet) UnmarshalObject(body []byte) (ConfigObj, error) {
	var policyTagSet PolicyTagSet
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &policyTagSet); err != nil {
			fmt.Println("### Trouble in unmarshaling PolicyTagSet from Json", body)
		}
	}
	return policyTagSet, err
}

type PolicyMatchTagSetCondition struct {
	/*
    This class was auto-generated by the GOSTRUCT plugin for PYANG
    from YANG module openconfig-routing-policy_mod
    based on the path /policy-definition/statement/match-tag-set.
    Each member element of the container is represented as a struct
    variable - with a specific YANG type.

	YANG Description: Match a referenced tag set according to the logic defined
in the match-options-set leaf 
    */

	//yang_name: tag-set class: leaf
	TagSet string
	Tag   PolicyTag  
	//yang_name: match-set-options class: leaf
	MatchSetOptions int32
}

type PolicyConditionConfig struct {
	BaseObj
	Name                                 string  `SNAPROUTE: "KEY"`
	ConditionType                        string
	MatchProtocolConditionInfo           string
    MatchDstIpPrefixConditionInfo        PolicyDstIpMatchPrefixSetCondition
    MatchNeighborConditionInfo           PolicyMatchNeighborSetCondition
	MatchTagConditionInfo                PolicyMatchTagSetCondition
}

func (obj PolicyConditionConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var policyConditionConfig PolicyConditionConfig
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &policyConditionConfig); err != nil {
			fmt.Println("### Trouble in unmarshaling PolicyConditionConfig from Json", body)
		}
	}
	return policyConditionConfig, err
}

type PolicyConditionState struct {
     BaseObj
	//yang_name: name class: leaf
	Name string  `SNAPROUTE: "KEY"`
	ConditionInfo string
	PolicyStmtList    []string
}

func (obj PolicyConditionState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var policyConditionState PolicyConditionState
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &policyConditionState); err != nil {
			fmt.Println("### Trouble in unmarshaling PolicyConditionState from Json", body)
		}
	}
	return policyConditionState, err
}


type PolicyActionConfig  struct {
	BaseObj 
	Name                        string `SNAPROUTE: "KEY"`
	ActionType                  string
	SetAdminDistanceValue       int
	Accept                      bool
	Reject                      bool
	RedistributeAction      string
	RedistributeTargetProtocol string 
	NetworkStatementTargetProtocol   string    
}

func (obj PolicyActionConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var policyActionConfig PolicyActionConfig
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &policyActionConfig); err != nil {
			fmt.Println("### Trouble in unmarshaling PolicyActionConfig from Json", body)
		}
	}
	return policyActionConfig, err
}

type PolicyActionState struct {
     BaseObj
	//yang_name: name class: leaf
	Name string  `SNAPROUTE: "KEY"`
	ActionInfo string
	PolicyStmtList    []string
}

func (obj PolicyActionState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var policyActionState PolicyActionState
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &policyActionState); err != nil {
			fmt.Println("### Trouble in unmarshaling PolicyActionState from Json", body)
		}
	}
	return policyActionState, err
}


type RouteDistanceState struct {
	BaseObj
	Protocol string `SNAPROUTE: "KEY"`
	Distance int
}
func (obj RouteDistanceState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var routeDistanceState RouteDistanceState
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &routeDistanceState); err != nil {
			fmt.Println("### Trouble in unmarshaling RouteDistanceState from Json", body)
		}
	}
	return routeDistanceState, err
}

type PolicyStmtConfig struct {
     BaseObj
	//yang_name: name class: leaf
	Name string  `SNAPROUTE: "KEY"`
//	AdminState string
	MatchConditions string
	Conditions []string
	Actions []string
}

func (obj PolicyStmtConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var policyStmt PolicyStmtConfig
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &policyStmt); err != nil {
			fmt.Println("### Trouble in unmarshaling PolicyStmtConfig from Json", body)
		}
	}
	return policyStmt, err
}

type PolicyStmtState struct {
     BaseObj
	//yang_name: name class: leaf
	Name string  `SNAPROUTE: "KEY"`
//	AdminState string
//	OperState  string
	MatchConditions string
	Conditions []string
	Actions []string
	PolicyList    []string
}

func (obj PolicyStmtState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var policyStmt PolicyStmtState
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &policyStmt); err != nil {
			fmt.Println("### Trouble in unmarshaling PolicyStmtState from Json", body)
		}
	}
	return policyStmt, err
}

type PolicyDefinitionStmtPrecedence struct {
	Precedence int
	Statement string
}

type PolicyDefinitionConfig struct {
	/*
    This class was auto-generated by the GOSTRUCT plugin for PYANG
    from YANG module openconfig-routing-policy_mod
    based on the path /policy-definition.
    Each member element of the container is represented as a struct
    variable - with a specific YANG type.

	YANG Description: List of top-level policy definitions, keyed by unique
name.  These policy definitions are expected to be
referenced (by name) in policy chains specified in import
or export configuration statements. 
    */
     BaseObj
	//yang_name: name class: leaf
	Name string  `SNAPROUTE: KEY`
	Precedence int
	MatchType string
	//yang_name: statement class: list
	StatementList []interface{} //PolicyDefinitionStmtPrecedence
}

func (obj PolicyDefinitionConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var policyDefinition PolicyDefinitionConfig
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &policyDefinition); err != nil {
			fmt.Println("### Trouble in unmarshaling PolicyDefinitionConfig from Json", body)
		}
	}
	return policyDefinition, err
}

type PolicyDefinitionState struct {
	/*
    This class was auto-generated by the GOSTRUCT plugin for PYANG
    from YANG module openconfig-routing-policy_mod
    based on the path /policy-definition/statement.
    Each member element of the container is represented as a struct
    variable - with a specific YANG type.

	YANG Description: Policy statements group conditions and actions
within a policy definition.  They are evaluated in
the order specified (see the description of policy
evaluation at the top of this module. 
    */
     BaseObj
	//yang_name: name class: leaf
	Name string  `SNAPROUTE: "KEY"`
	HitCounter int
	IpPrefixList []string
}

func (obj PolicyDefinitionState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var policyDefinitionStmt PolicyDefinitionState
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &policyDefinitionStmt); err != nil {
			fmt.Println("### Trouble in unmarshaling PolicyDefinitionState from Json", body)
		}
	}
	return policyDefinitionStmt, err
}

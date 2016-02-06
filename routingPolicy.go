package models

import (
	"encoding/json"
	"fmt"
)

type PolicyDefinitionSetsPrefix struct {
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

type PolicyDefinitionSetsPrefixSet struct {
	BaseObj
	PrefixSetName string   `SNAPROUTE: "KEY"`
    /*
     List of prefix expressions that are part of the set
	*/
	IpPrefixList [] PolicyDefinitionSetsPrefix
}

func (obj PolicyDefinitionSetsPrefixSet) UnmarshalObject(body []byte) (ConfigObj, error) {
	var policyDefinitionSetsPrefixSet PolicyDefinitionSetsPrefixSet
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &policyDefinitionSetsPrefixSet); err != nil {
			fmt.Println("### Trouble in unmarshaling PolicyDefinitionSetsPrefixSet from Json", body)
		}
	}
	return policyDefinitionSetsPrefixSet, err
}


type PolicyDefinitionStmtDstIpMatchPrefixSetCondition struct {
	BaseObj
	/*
    This class was auto-generated by the GOSTRUCT plugin for PYANG
    from YANG module openconfig-routing-policy_mod
    based on the path /policy-definition/statement/match-prefix-set.
    Each member element of the container is represented as a struct
    variable - with a specific YANG type.

	YANG Description: Match a referenced prefix-set according to the logic
defined in the match-set-options leaf 
    */
	Name string  `SNAPROUTE: "KEY"`
	//yang_name: prefix-set class: leaf
	PrefixSet string
	//yang_name: match-set-options class: leaf
    Prefix PolicyDefinitionSetsPrefix 
}

func (obj PolicyDefinitionStmtDstIpMatchPrefixSetCondition) UnmarshalObject(body []byte) (ConfigObj, error) {
	var policyDefinitionStmtMatchPrefixSetCondition PolicyDefinitionStmtDstIpMatchPrefixSetCondition
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &policyDefinitionStmtMatchPrefixSetCondition); err != nil {
			fmt.Println("### Trouble in unmarshaling PolicyDefinitionStmtDstIpMatchPrefixSetCondition from Json", body)
		}
	}
	return policyDefinitionStmtMatchPrefixSetCondition, err
}

type PolicyDefinitionSetsNeighbor struct {
	Address [] string
}

type PolicyDefinitionSetsNeighborSet struct {
	BaseObj
	NeighborSetName string `SNAPROUTE: "KEY"`
	NeihgborSet []PolicyDefinitionSetsNeighbor
}

func (obj PolicyDefinitionSetsNeighborSet) UnmarshalObject(body []byte) (ConfigObj, error) {
	var policyDefinitionSetsNeighborSet PolicyDefinitionSetsNeighborSet
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &policyDefinitionSetsNeighborSet); err != nil {
			fmt.Println("### Trouble in unmarshaling PolicyDefinitionSetsNeighborSet from Json", body)
		}
	}
	return policyDefinitionSetsNeighborSet, err
}

type PolicyDefinitionStmtMatchNeighborSetCondition struct {
	BaseObj
	/*
    This class was auto-generated by the GOSTRUCT plugin for PYANG
    from YANG module openconfig-routing-policy_mod
    based on the path /policy-definition/statement/match-neighbor-set.
    Each member element of the container is represented as a struct
    variable - with a specific YANG type.

	YANG Description: Match a referenced neighbor set according to the logic
defined in the match-set-options-leaf 
    */

	Name string   `SNAPROUTE: "KEY"`
	//yang_name: neighbor-set class: leaf
	NeighborSet string
	//yang_name: match-set-options class: leaf
	MatchSetOptions int32
}

func (obj PolicyDefinitionStmtMatchNeighborSetCondition) UnmarshalObject(body []byte) (ConfigObj, error) {
	var policyDefinitionStmtMatchNeighborSetCondition PolicyDefinitionStmtMatchNeighborSetCondition
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &policyDefinitionStmtMatchNeighborSetCondition); err != nil {
			fmt.Println("### Trouble in unmarshaling PolicyDefinitionStmtMatchNeighborSetCondition from Json", body)
		}
	}
	return policyDefinitionStmtMatchNeighborSetCondition, err
}

type PolicyDefinitionSetsTag struct {
	value [] uint32
}

type PolicyDefinitionSetsTagSet struct {
	BaseObj
	TagSetName string               `SNAPROUTE: "KEY"`
	TagList []PolicyDefinitionSetsTag
}

func (obj PolicyDefinitionSetsTagSet) UnmarshalObject(body []byte) (ConfigObj, error) {
	var policyDefinitionSetsTagSet PolicyDefinitionSetsTagSet
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &policyDefinitionSetsTagSet); err != nil {
			fmt.Println("### Trouble in unmarshaling PolicyDefinitionSetsTagSet from Json", body)
		}
	}
	return policyDefinitionSetsTagSet, err
}

type PolicyDefinitionStmtMatchTagSetCondition struct {
	BaseObj
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
	Name string   `SNAPROUTE: "KEY"`
	TagSet string
	//yang_name: match-set-options class: leaf
	MatchSetOptions int32
}

func (obj PolicyDefinitionStmtMatchTagSetCondition) UnmarshalObject(body []byte) (ConfigObj, error) {
	var policyDefinitionStmtMatchTagSetCondition PolicyDefinitionStmtMatchTagSetCondition
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &policyDefinitionStmtMatchTagSetCondition); err != nil {
			fmt.Println("### Trouble in unmarshaling PolicyDefinitionStmtMatchTagSetCondition from Json", body)
		}
	}
	return policyDefinitionStmtMatchTagSetCondition, err
}

type PolicyDefinitionStmtMatchProtocolCondition struct {
	BaseObj
	Name string  `SNAPROUTE: "KEY"`
	InstallProtocolEq string		
}

func (obj PolicyDefinitionStmtMatchProtocolCondition) UnmarshalObject(body []byte) (ConfigObj, error) {
	var policyDefinitionStmtMatchProtocolCondition PolicyDefinitionStmtMatchProtocolCondition
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &policyDefinitionStmtMatchProtocolCondition); err != nil {
			fmt.Println("### Trouble in unmarshaling PolicyDefinitionStmtMatchProtocolCondition from Json", body)
		}
	}
	return policyDefinitionStmtMatchProtocolCondition, err
}

type PolicyDefinitionStmtIgpActions struct {
	BaseObj
	/*
    This class was auto-generated by the GOSTRUCT plugin for PYANG
    from YANG module openconfig-routing-policy_mod
    based on the path /policy-definition/statement/igp-actions.
    Each member element of the container is represented as a struct
    variable - with a specific YANG type.

	YANG Description: Actions to set IGP route attributes; these actions
apply to multiple IGPs 
    */

	//yang_name: set-tag class: leaf
	Name string  `SNAPROUTE: "KEY"`
	SetTag []uint32
}

func (obj PolicyDefinitionStmtIgpActions) UnmarshalObject(body []byte) (ConfigObj, error) {
	var policyDefinitionStmtIgpActions PolicyDefinitionStmtIgpActions
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &policyDefinitionStmtIgpActions); err != nil {
			fmt.Println("### Trouble in unmarshaling PolicyDefinitionStmtIgpActions from Json", body)
		}
	}
	return policyDefinitionStmtIgpActions, err
}

type PolicyDefinitionStmtRouteDispositionAction struct {
	BaseObj
	Name string  `SNAPROUTE: "KEY"`
	RouteDisposition string    //Accept or Reject
}

func (obj PolicyDefinitionStmtRouteDispositionAction) UnmarshalObject(body []byte) (ConfigObj, error) {
	var policyDefinitionStmtRouteDispositionAction PolicyDefinitionStmtRouteDispositionAction
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &policyDefinitionStmtRouteDispositionAction); err != nil {
			fmt.Println("### Trouble in unmarshaling PolicyDefinitionStmtRouteDispositionAction from Json", body)
		}
	}
	return policyDefinitionStmtRouteDispositionAction, err
}

type PolicyDefinitionStmtRedistributionAction struct {
	BaseObj
	Name string  `SNAPROUTE: "KEY"`
	Redistribute bool
	RedistributeTargetProtocol string 
}

func (obj PolicyDefinitionStmtRedistributionAction) UnmarshalObject(body []byte) (ConfigObj, error) {
	var policyDefinitionStmtRedistributionAction PolicyDefinitionStmtRedistributionAction
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &policyDefinitionStmtRedistributionAction); err != nil {
			fmt.Println("### Trouble in unmarshaling PolicyDefinitionStmtRedistributionAction from Json", body)
		}
	}
	return policyDefinitionStmtRedistributionAction, err
}

type PolicyDefinitionStmtConfig struct {
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
//	AdminState string
	MatchConditions string
	Conditions []string
	Actions []string
//	Export  bool
//	Import  bool
}

func (obj PolicyDefinitionStmtConfig) UnmarshalObject(body []byte) (ConfigObj, error) {
	var policyDefinitionStmt PolicyDefinitionStmtConfig
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &policyDefinitionStmt); err != nil {
			fmt.Println("### Trouble in unmarshaling PolicyDefinitionStatementConfig from Json", body)
		}
	}
	return policyDefinitionStmt, err
}

type PolicyDefinitionStmtState struct {
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
//	AdminState string
//	OperState  string
	MatchConditions string
	Conditions []string
	Actions []string
	Export  bool
	Import  bool
	HitCounter int
	IpPrefixList []string
}

func (obj PolicyDefinitionStmtState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var policyDefinitionStmt PolicyDefinitionStmtState
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &policyDefinitionStmt); err != nil {
			fmt.Println("### Trouble in unmarshaling PolicyDefinitionStmtState from Json", body)
		}
	}
	return policyDefinitionStmt, err
}

type PolicyDefinitionConditionState struct {
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
	ConditionInfo string
	PolicyList    []string
}

func (obj PolicyDefinitionConditionState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var policyDefinitionConditionState PolicyDefinitionConditionState
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &policyDefinitionConditionState); err != nil {
			fmt.Println("### Trouble in unmarshaling PolicyDefinitionConditionState from Json", body)
		}
	}
	return policyDefinitionConditionState, err
}

type PolicyDefinitionActionState struct {
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
	ActionInfo string
	PolicyList    []string
}

func (obj PolicyDefinitionActionState) UnmarshalObject(body []byte) (ConfigObj, error) {
	var policyDefinitionActionState PolicyDefinitionActionState
	var err error
	if len(body) > 0 {
		if err = json.Unmarshal(body, &policyDefinitionActionState); err != nil {
			fmt.Println("### Trouble in unmarshaling PolicyDefinitionActionState from Json", body)
		}
	}
	return policyDefinitionActionState, err
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
	Export bool
	Import bool
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

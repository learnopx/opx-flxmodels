package models


type Dot1dStpPortEntryConfig struct {
	BaseObj
	Dot1dStpPort                  int32 `SNAPROUTE: KEY` //The port number of the port for which this entry contains Spanning Tree Protocol management information.
	Dot1dBrgIfIndex               int32 //The value of the instance of the ifIndex object,  for the bridge corresponding to this port.
	Dot1dStpPortPriority          int32 //The value of the priority field that is contained in the first (in network byte order) octet of the (2 octet long) Port ID.  The other octet of the Port ID is given by the value of dot1dStpPort. On bridges supporting IEEE 802.1t or IEEE 802.1w, permissible values are 0-240, in steps of 16.
	Dot1dStpPortEnable            int32 //The enabled/disabled status of the port.
	Dot1dStpPortPathCost          int32 //The contribution of this port to the path cost of paths towards the spanning tree root which include this port.  802.1D-1998 recommends that the default value of this parameter be in inverse proportion to    the speed of the attached LAN.  New implementations should support dot1dStpPortPathCost32. If the port path costs exceeds the maximum value of this object then this object should report the maximum value, namely 65535.  Applications should try to read the dot1dStpPortPathCost32 object if this object reports the maximum value.
	Dot1dStpPortPathCost32        int32 //The contribution of this port to the path cost of paths towards the spanning tree root which include this port.  802.1D-1998 recommends that the default value of this parameter be in inverse proportion to the speed of the attached LAN.  This object replaces dot1dStpPortPathCost to support IEEE 802.1t.
	Dot1dStpPortProtocolMigration int32 //When operating in RSTP (version 2) mode, writing true(1) to this object forces this port to transmit RSTP BPDUs. Any other operation on this object has no effect and it always returns false(2) when read.
	Dot1dStpPortAdminPointToPoint int32 //The administrative point-to-point status of the LAN segment attached to this port, using the enumeration values of the IEEE 802.1w clause.  A value of forceTrue(0) indicates that this port should always be treated as if it is connected to a point-to-point link.  A value of forceFalse(1) indicates that this port should be treated as having a shared media connection.  A value of auto(2) indicates that this port is considered to have a point-to-point link if it is an Aggregator and all of its    members are aggregatable, or if the MAC entity is configured for full duplex operation, either through auto-negotiation or by management means.  Manipulating this object changes the underlying adminPortToPortMAC.  The value of this object MUST be retained across reinitializations of the management system.
	Dot1dStpPortAdminEdgePort     int32 //The administrative value of the Edge Port parameter.  A value of true(1) indicates that this port should be assumed as an edge-port, and a value of false(2) indicates that this port should be assumed as a non-edge-port.    Setting this object will also cause the corresponding instance of dot1dStpPortOperEdgePort to change to the same value.  Note that even when this object's value is true, the value of the corresponding instance of dot1dStpPortOperEdgePort can be false if a BPDU has been received.  The value of this object MUST be retained across reinitializations of the management system.
	Dot1dStpPortAdminPathCost     int32 //The administratively assigned value for the contribution of this port to the path cost of paths toward the spanning tree root.  Writing a value of '0' assigns the automatically calculated default Path Cost value to the port.  If the default Path Cost is being used, this object returns '0' when read.  This complements the object dot1dStpPortPathCost or dot1dStpPortPathCost32, which returns the operational value of the path cost.    The value of this object MUST be retained across reinitializations of the management system.
}

type Dot1dStpPortEntryStateCountersFsmStates struct {
	BaseObj
	RstpOutPkts                    uint64 //Number of STP BPDUs transmitted
	RstpInPkts                     uint64 //Number of RSTP BPDUs received
	Dot1dStpPortDesignatedPort     string //The Port Identifier of the port on the Designated Bridge for this port's segment.
	Dot1dStpBridgePortForwardDelay int32  //The value that all bridges use for ForwardDelay as advertised by the root bridge.  Note that 802.1D-1998 specifies that the range for this parameter is related to the value of dot1dStpBridgeMaxAge.  The granularity of this timer is specified by 802.1D-1998 to be 1 second.  An agent may return a badValue error if a set is attempted to a value that is not a whole number of seconds.
	Dot1dStpPortPathCost           int32  //The contribution of this port to the path cost of paths towards the spanning tree root which include this port.  802.1D-1998 recommends that the default value of this parameter be in inverse proportion to    the speed of the attached LAN.  New implementations should support dot1dStpPortPathCost32. If the port path costs exceeds the maximum value of this object then this object should report the maximum value, namely 65535.  Applications should try to read the dot1dStpPortPathCost32 object if this object reports the maximum value.
	Dot1dStpPortDesignatedBridge   string //The Bridge Identifier of the bridge that this port considers to be the Designated Bridge for this port's segment.
	TcOutPkts                      uint64 //Number of TC BPDUs transmitted
	Dot1dStpPortAdminPointToPoint  int32  //The administrative point-to-point status of the LAN segment attached to this port, using the enumeration values of the IEEE 802.1w clause.  A value of forceTrue(0) indicates that this port should always be treated as if it is connected to a point-to-point link.  A value of forceFalse(1) indicates that this port should be treated as having a shared media connection.  A value of auto(2) indicates that this port is considered to have a point-to-point link if it is an Aggregator and all of its    members are aggregatable, or if the MAC entity is configured for full duplex operation, either through auto-negotiation or by management means.  Manipulating this object changes the underlying adminPortToPortMAC.  The value of this object MUST be retained across reinitializations of the management system.
	PvstInPkts                     uint64 //Number of PVST BPDUs received
	Dot1dStpPortProtocolMigration  int32  //When operating in RSTP (version 2) mode, writing true(1) to this object forces this port to transmit RSTP BPDUs. Any other operation on this object has no effect and it always returns false(2) when read.
	Dot1dStpPortState              int32  //The port's current state, as defined by application of the Spanning Tree Protocol.  This state controls what action a port takes on reception of a frame.  If the bridge has detected a port that is malfunctioning, it will place that port into the broken(6) state.  For ports that are disabled (see dot1dStpPortEnable), this object will have a value of disabled(1).
	Dot1dStpPortEnable             int32  //The enabled/disabled status of the port.
	Dot1dStpPortDesignatedRoot     string //The unique Bridge Identifier of the Bridge recorded as the Root in the Configuration BPDUs transmitted by the Designated Bridge for the segment to which the port is attached.
	BpduOutPkts                    uint64 //Number of BPDUs transmitted
	Dot1dStpBridgePortMaxAge       int32  //The value that all bridges use for MaxAge as advertised by the root bridge.  Note that 802.1D-1998 specifies that the range for this parameter is related to the value of dot1dStpBridgeHelloTime.  The granularity of this timer is specified by 802.1D-1998 to be 1 second.  An agent may return a badValue error if a set is attempted to a value that is not a whole number of seconds.
	Dot1dStpPortOperPointToPoint   int32  //The operational point-to-point status of the LAN segment attached to this port.  It indicates whether a port is considered to have a point-to-point connection. If adminPointToPointMAC is set to auto(2), then the value of operPointToPointMAC is determined in accordance with the specific procedures defined for the MAC entity concerned, as defined in IEEE 802.1w, clause 6.5.  The value is determined dynamically; that is, it is re-evaluated whenever the value of adminPointToPointMAC changes, and whenever the specific procedures defined for the MAC entity evaluate a change in its point-to-point status.
	Dot1dBrgIfIndex                int32  //The value of the instance of the ifIndex object,  for the bridge corresponding to this port.
	Dot1dStpPortOperEdgePort       int32  //The operational value of the Edge Port parameter.  The object is initialized to the value of the corresponding instance of dot1dStpPortAdminEdgePort.  When the corresponding instance of dot1dStpPortAdminEdgePort is set, this object will be changed as well.  This object will also be changed to false on reception of a BPDU.
	Dot1dStpPortAdminEdgePort      int32  //The administrative value of the Edge Port parameter.  A value of true(1) indicates that this port should be assumed as an edge-port, and a value of false(2) indicates that this port should be assumed as a non-edge-port.    Setting this object will also cause the corresponding instance of dot1dStpPortOperEdgePort to change to the same value.  Note that even when this object's value is true, the value of the corresponding instance of dot1dStpPortOperEdgePort can be false if a BPDU has been received.  The value of this object MUST be retained across reinitializations of the management system.
	Dot1dStpPortForwardTransitions uint32 //The number of times this port has transitioned from the Learning state to the Forwarding state.
	PvstOutPkts                    uint64 //Number of PVST BPDUs transmitted
	StpOutPkts                     uint64 //Number of STP BPDUs transmitted
	Dot1dStpPort                   int32  `SNAPROUTE: KEY` //The port number of the port for which this entry contains Spanning Tree Protocol management information.
	Dot1dStpBridgePortHelloTime    int32  //The value that all bridges use for HelloTime as advertised by the root bridge.  The granularity of this timer is specified by 802.1D-1998 to be 1 second.  An agent may return a badValue error if a set is attempted    to a value that is not a whole number of seconds.
	BpduInPkts                     uint64 //Number of BPDUs received
	Dot1dStpPortPriority           int32  //The value of the priority field that is contained in the first (in network byte order) octet of the (2 octet long) Port ID.  The other octet of the Port ID is given by the value of dot1dStpPort. On bridges supporting IEEE 802.1t or IEEE 802.1w, permissible values are 0-240, in steps of 16.
	StpInPkts                      uint64 //Number of STP PDUs received
	TcInPkts                       uint64 //Number of TC BPDUs received
	Dot1dStpPortAdminPathCost      int32  //The administratively assigned value for the contribution of this port to the path cost of paths toward the spanning tree root.  Writing a value of '0' assigns the automatically calculated default Path Cost value to the port.  If the default Path Cost is being used, this object returns '0' when read.  This complements the object dot1dStpPortPathCost or dot1dStpPortPathCost32, which returns the operational value of the path cost.    The value of this object MUST be retained across reinitializations of the management system.
	Dot1dStpPortPathCost32         int32  //The contribution of this port to the path cost of paths towards the spanning tree root which include this port.  802.1D-1998 recommends that the default value of this parameter be in inverse proportion to the speed of the attached LAN.  This object replaces dot1dStpPortPathCost to support IEEE 802.1t.
	Dot1dStpPortDesignatedCost     int32  //The path cost of the Designated Port of the segment connected to this port.  This value is compared to the Root Path Cost field in received bridge PDUs.
	PimPrevState                   string //PIM previous fsm state
	PimCurrState                   string //PIM current fsm state
	PrtmPrevState                  string //PRTM previous fsm state
	PrtmCurrState                  string //PRTM current fsm state
	PrxmPrevState                  string //PRXM previous fsm state
	PrxmCurrState                  string //PRXM current fsm state
	PstmPrevState                  string //PSTM previous fsm state
	PstmCurrState                  string //PSTM current fsm state
	TcmPrevState                   string //TCM previous fsm state
	TcmCurrState                   string //TCM current fsm state
	PpmPrevState                   string //PPM previous fsm state
	PpmCurrState                   string //PPM current fsm state
	PtxmPrevState                  string //PTXM previous fsm state
	PtxmCurrState                  string //PTXM current fsm state
	PtimPrevState                  string //PTIM previous fsm state
	PtimCurrState                  string //PTIM current fsm state
	BdmPrevState                   string //BDM previous fsm state
	BdmCurrState                   string //BDM current fsm state
}

type Dot1dStpBridgeConfig struct {
	BaseObj
	Dot1dBridgeAddress         string `SNAPROUTE: KEY` //The bridge identifier of the root of the spanning tree, as determined by the Spanning Tree Protocol, as executed by this node.  This value is used as the Root Identifier parameter in all Configuration Bridge PDUs originated by this node.
	Dot1dStpVlan               uint16 `SNAPROUTE: KEY` //Each bridge is associated with a domain.  Typically this domain is represented as the vlan; The default domain is typically 1
	Dot1dStpPriority           int32  `SNAPROUTE: KEY` //The value of the write-able portion of the Bridge ID (i.e., the first two octets of the (8 octet long) Bridge ID).  The other (last) 6 octets of the Bridge ID are given by the value of dot1dBridgeAddress. On bridges supporting IEEE 802.1t or IEEE 802.1w, permissible values are 0-61440, in steps of 4096.
	Dot1dStpBridgeMaxAge       int32  //The value that all bridges use for MaxAge when this bridge is acting as the root.  Note that 802.1D-1998 specifies that the range for this parameter is related to the value of dot1dStpBridgeHelloTime.  The granularity of this timer is specified by 802.1D-1998 to be 1 second.  An agent may return a badValue error if a set is attempted to a value that is not a whole number of seconds.
	Dot1dStpBridgeHelloTime    int32  //The value that all bridges use for HelloTime when this bridge is acting as the root.  The granularity of this timer is specified by 802.1D-1998 to be 1 second.  An agent may return a badValue error if a set is attempted    to a value that is not a whole number of seconds.
	Dot1dStpBridgeForwardDelay int32  //The value that all bridges use for ForwardDelay when this bridge is acting as the root.  Note that 802.1D-1998 specifies that the range for this parameter is related to the value of dot1dStpBridgeMaxAge.  The granularity of this timer is specified by 802.1D-1998 to be 1 second.  An agent may return a badValue error if a set is attempted to a value that is not a whole number of seconds.
	Dot1dStpBridgeForceVersion int32  //TODO
	Dot1dStpBridgeTxHoldCount  int32  //TODO
}

type Dot1dStpBridgeState struct {
	BaseObj
	Dot1dStpBridgeForceVersion      int32  //TODO
	Dot1dBridgeAddress              string `SNAPROUTE: KEY` //The bridge identifier of the root of the spanning tree, as determined by the Spanning Tree Protocol, as executed by this node.  This value is used as the Root Identifier parameter in all Configuration Bridge PDUs originated by this node.
	Dot1dStpBridgeHelloTime         int32  //The value that all bridges use for HelloTime when this bridge is acting as the root.  The granularity of this timer is specified by 802.1D-1998 to be 1 second.  An agent may return a badValue error if a set is attempted    to a value that is not a whole number of seconds.
	Dot1dStpBridgeTxHoldCount       int32  //TODO
	Dot1dStpVlan                    uint16 `SNAPROUTE: KEY` //Each bridge is associated with a domain.  Typically this domain is represented as the vlan; The default domain is typically 1
	Dot1dStpBridgeForwardDelay      int32  //The value that all bridges use for ForwardDelay when this bridge is acting as the root.  Note that 802.1D-1998 specifies that the range for this parameter is related to the value of dot1dStpBridgeMaxAge.  The granularity of this timer is specified by 802.1D-1998 to be 1 second.  An agent may return a badValue error if a set is attempted to a value that is not a whole number of seconds.
	Dot1dStpBridgeMaxAge            int32  //The value that all bridges use for MaxAge when this bridge is acting as the root.  Note that 802.1D-1998 specifies that the range for this parameter is related to the value of dot1dStpBridgeHelloTime.  The granularity of this timer is specified by 802.1D-1998 to be 1 second.  An agent may return a badValue error if a set is attempted to a value that is not a whole number of seconds.
	Dot1dStpPriority                int32  `SNAPROUTE: KEY` //The value of the write-able portion of the Bridge ID (i.e., the first two octets of the (8 octet long) Bridge ID).  The other (last) 6 octets of the Bridge ID are given by the value of dot1dBridgeAddress. On bridges supporting IEEE 802.1t or IEEE 802.1w, permissible values are 0-61440, in steps of 4096.
	Dot1dBrgIfIndex                 int32  //The value of the instance of the ifIndex object,  for the bridge
	Dot1dStpProtocolSpecification   int32  //An indication of what version of the Spanning Tree Protocol is being run.  The value 'decLb100(2)' indicates the DEC LANbridge 100 Spanning Tree protocol. IEEE 802.1D implementations will return 'ieee8021d(3)'. If future versions of the IEEE Spanning Tree Protocol that are incompatible with the current version are released a new value will be defined.
	Dot1dStpTimeSinceTopologyChange uint32 //The time (in hundredths of a second) since the last time a topology change was detected by the bridge entity. For RSTP, this reports the time since the tcWhile timer for any port on this Bridge was nonzero.
	Dot1dStpTopChanges              uint32 //The total number of topology changes detected by this bridge since the management entity was last reset or initialized.
	Dot1dStpDesignatedRoot          string //The bridge identifier of the root of the spanning tree, as determined by the Spanning Tree Protocol, as executed by this node.  This value is used as the Root Identifier parameter in all Configuration Bridge PDUs originated by this node.
	Dot1dStpRootCost                int32  //The cost of the path to the root as seen from this bridge.
	Dot1dStpRootPort                int32  //The port number of the port that offers the lowest cost path from this bridge to the root bridge.
	Dot1dStpMaxAge                  int32  //The maximum age of Spanning Tree Protocol information learned from the network on any port before it is discarded, in units of hundredths of a second.  This is the actual value that this bridge is currently using.
	Dot1dStpHelloTime               int32  //The amount of time between the transmission of Configuration bridge PDUs by this node on any port when it is the root of the spanning tree, or trying to become so, in units of hundredths of a second.  This is the actual value that this bridge is currently using.
	Dot1dStpHoldTime                int32  //This time value determines the interval length during which no more than two Configuration bridge PDUs shall be transmitted by this node, in units of hundredths of a second.
	Dot1dStpForwardDelay            int32  //This time value, measured in units of hundredths of a second, controls how fast a port changes its spanning state when moving towards the Forwarding state.  The value determines how long the port stays in each of the Listening and Learning states, which precede the Forwarding state.  This value is also used when a topology change has been detected and is underway, to age all dynamic entries in the Forwarding Database. [Note that this value is the one that this bridge is currently using, in contrast to dot1dStpBridgeForwardDelay, which is the value that this bridge and all others would start using if/when this bridge were to become the root.]
}


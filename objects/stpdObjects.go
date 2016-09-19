package objects

type StpGlobal struct {
	baseObj
	Vrf        string `SNAPROUTE: "KEY", ACCESS:"w",  MULTIPLICITY:"1", AUTOCREATE: "true", DEFAULT: "default", DESCRIPTION: global system object defining the global state of STPD.`
	AdminState string `DESCRIPTION: Administrative state of STPD, UP will allow for stp/rstp/pvst+ configuration to be applied, DOWN will disallow and de-provision from daemon, except for default STP instance. STRLEN:"4", SELECTION: UP/DOWN, DEFAULT: UP`
}

type StpPort struct {
	baseObj
	Vlan              int32  `SNAPROUTE: "KEY", ACCESS:"rw", MULTIPLICITY:"*", AUTODISCOVER:"true", DESCRIPTION: The value of instance of the vlan object,  for the bridge corresponding to this port., MIN: "0" ,  MAX: "4094"`
	IntfRef           string `SNAPROUTE: "KEY", ACCESS:"rw", DESCRIPTION: The port number of the port for which this entry contains Spanning Tree Protocol management information. `
	Priority          int32  `DESCRIPTION: The value of the priority field that is contained in the first in network byte order octet of the 2 octet long Port ID.  The other octet of the Port ID is given by the value of StpPort. On bridges supporting IEEE 802.1t or IEEE 802.1w, permissible values are 0-240, in steps of 16., MIN: "0" ,  MAX: "255", DEFAULT: 128`
	AdminState        string `DESCRIPTION: The enabled/disabled status of the port., SELECTION: UP/DOWN, DEFAULT: UP`
	PathCost          int32  `DESCRIPTION: The contribution of this port to the path cost of paths towards the spanning tree root which include this port.  802.1D-1998 recommends that the default value of this parameter be in inverse proportion to the speed of the attached LAN.  New implementations should support PathCost32. If the port path costs exceeds the maximum value of this object then this object should report the maximum value; namely 65535.  Applications should try to read the PathCost32 object if this object reports the maximum value.  Value of 1 will force node to auto discover the value        based on the ports capabilities., MIN: "1" ,  MAX: "65535", DEFAULT: 1`
	PathCost32        int32  `DESCRIPTION: The contribution of this port to the path cost of paths towards the spanning tree root which include this port.  802.1D-1998 recommends that the default value of this parameter be in inverse proportion to the speed of the attached LAN.  This object replaces PathCost to support IEEE 802.1t. Value of 1 will force node to auto discover the value        based on the ports capabilities., MIN: "1" ,  MAX: "200000000", DEFAULT: 1`
	ProtocolMigration int32  `DESCRIPTION: When operating in RSTP (version 2) mode writing true(1) to this object forces this port to transmit RSTP BPDUs. Any other operation on this object has no effect and it always returns false(2) when read., SELECTION: false(2)/true(1), DEFAULT: 1`
	AdminPointToPoint int32  `DESCRIPTION: The administrative point-to-point status of the LAN segment attached to this port using the enumeration values of the IEEE 802.1w clause.  A value of forceTrue(0) indicates that this port should always be treated as if it is connected to a point-to-point link.  A value of forceFalse(1) indicates that this port should be treated as having a shared media connection.  A value of auto(2) indicates that this port is considered to have a point-to-point link if it is an Aggregator and all of its    members are aggregatable or if the MAC entity is configured for full duplex operation, either through auto-negotiation or by management means.  Manipulating this object changes the underlying adminPortToPortMAC.  The value of this object MUST be retained across reinitializations of the management system., SELECTION: forceTrue(0)/forceFalse(1)/auto(2), DEFAULT: 2`
	AdminEdgePort     int32  `DESCRIPTION: The administrative value of the Edge Port parameter.  A value of true(1) indicates that this port should be assumed as an edge-port and a value of false(2) indicates that this port should be assumed as a non-edge-port.  Setting this object will also cause the corresponding instance of OperEdgePort to change to the same value.  Note that even when this object's value is true the value of the corresponding instance of OperEdgePort can be false if a BPDU has been received.  The value of this object MUST be retained across reinitializations of the management system., SELECTION: false(2)/true(1), DEFAULT: 2`
	AdminPathCost     int32  `DESCRIPTION: The administratively assigned value for the contribution of this port to the path cost of paths toward the spanning tree root.  Writing a value of '0' assigns the automatically calculated default Path Cost value to the port.  If the default Path Cost is being used this object returns '0' when read.  This complements the object PathCost or PathCost32 which returns the operational value of the path cost.    The value of this object MUST be retained across reinitializations of the management system., MIN: "0" ,  MAX: "200000000", DEFAULT: 200000`
	BpduGuard         int32  `DESCRIPTION: A Port as OperEdge which receives BPDU with BpduGuard enabled will shut the port down., SELECTION: false(2)/true(1), DEFAULT: 2`
	BpduGuardInterval int32  `DESCRIPTION: The interval time to which a port will try to recover from BPDU Guard err-disable state.  If no BPDU frames are detected after this timeout plus 3 Times Hello Time then the port will transition back to Up state.  If condition is cleared manually then this operation is ignored.  If set to zero then timer is inactive and recovery is based on manual intervention., DEFAULT: 15`
	BridgeAssurance   int32  `DESCRIPTION: When enabled BPDUs will be transmitted out of all stp ports regardless of state.  When an stp port fails to receive a BPDU the port should  transition to a Blocked state.  Upon reception of BDPU after shutdown  should transition port into the bridge., SELECTION: false(2)/true(1), DEFAULT: 2`
}

type StpPortState struct {
	baseObj
	Vlan                        int32  `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: The value of instance of the vlan object,  for the bridge corresponding to this port., MIN: "0" ,  MAX: "4094"`
	IntfRef                     string `SNAPROUTE: "KEY", ACCESS:"r", DESCRIPTION: The port number of the port for which this entry contains Spanning Tree Protocol management information. `
	Priority                    int32  `DESCRIPTION: The value of the priority field that is contained in the first in network byte order octet of the 2 octet long Port ID.  The other octet of the Port ID is given by the value of StpPort. On bridges supporting IEEE 802.1t or IEEE 802.1w permissible values are 0-240 in steps of 16., MIN: "0" ,  MAX: "255"`
	Enable                      int32  `DESCRIPTION: The enabled/disabled status of the port., SELECTION: disabled(2)/enabled(1), DEFAULT: 1`
	PathCost                    int32  `DESCRIPTION: The contribution of this port to the path cost of paths towards the spanning tree root which include this port.  802.1D-1998 recommends that the default value of this parameter be in inverse proportion to    the speed of the attached LAN.  New implementations should support PathCost32. If the port path costs exceeds the maximum value of this object then this object should report the maximum value namely 65535.  Applications should try to read the PathCost32 object if this object reports the maximum value., MIN: "1" ,  MAX: "65535"`
	PathCost32                  int32  `DESCRIPTION: The contribution of this port to the path cost of paths towards the spanning tree root which include this port.  802.1D-1998 recommends that the default value of this parameter be in inverse proportion to the speed of the attached LAN.  This object replaces PathCost to support IEEE 802.1t., MIN: "1" ,  MAX: "200000000"`
	State                       int32  `DESCRIPTION: The port's current state as defined by application of the Spanning Tree Protocol.  This state controls what action a port takes on reception of a frame.  If the bridge has detected a port that is malfunctioning it will place that port into the broken(6) state.  For ports that are disabled (see Enable) this object will have a value of disabled(1)., SELECTION: listening(3)/disabled(1)/broken(6)/learning(4)/forwarding(5)/blocking(2)`
	DesignatedRoot              string `DESCRIPTION: The unique Bridge Identifier of the Bridge recorded as the Root in the Configuration BPDUs transmitted by the Designated Bridge for the segment to which the port is attached., LEN : "8"`
	DesignatedCost              int32  `DESCRIPTION: The path cost of the Designated Port of the segment connected to this port.  This value is compared to the Root Path Cost field in received bridge PDUs.`
	DesignatedBridge            string `DESCRIPTION: The Bridge Identifier of the bridge that this port considers to be the Designated Bridge for this port's segment., LEN : "8"`
	DesignatedPort              string `DESCRIPTION: The Port Identifier of the port on the Designated Bridge for this port's segment., LEN : "2"`
	ForwardTransitions          uint32 `DESCRIPTION: The number of times this port has transitioned from the Learning state to the Forwarding state.`
	AdminEdgePort               int32  `DESCRIPTION: The administrative value of the Edge Port parameter.  A value of true(1) indicates that this port should be assumed as an edge-port and a value of false(2) indicates that this port should be assumed as a non-edge-port.    Setting this object will also cause the corresponding instance of OperEdgePort to change to the same value.  Note that even when this object's value is true the value of the corresponding instance of OperEdgePort can be false if a BPDU has been received.  The value of this object MUST be retained across reinitializations of the management system., SELECTION: false(2)/true(1)`
	AdminPathCost               int32  `DESCRIPTION: The administratively assigned value for the contribution of this port to the path cost of paths toward the spanning tree root.  Writing a value of '0' assigns the automatically calculated default Path Cost value to the port.  If the default Path Cost is being used this object returns '0' when read.  This complements the object PathCost or PathCost32 which returns the operational value of the path cost.    The value of this object MUST be retained across reinitializations of the management system., MIN: "0" ,  MAX: "200000000"`
	OperEdgePort                int32  `DESCRIPTION: The operational value of the Edge Port parameter.  The object is initialized to the value of the corresponding instance of AdminEdgePort.  When the corresponding instance of AdminEdgePort is set this object will be changed as well.  This object will also be changed to false on reception of a BPDU., SELECTION: false(2)/true(1)`
	OperPointToPoint            int32  `DESCRIPTION: The operational point-to-point status of the LAN segment attached to this port.  It indicates whether a port is considered to have a point-to-point connection. If adminPointToPointMAC is set to auto(2) then the value of operPointToPointMAC is determined in accordance with the specific procedures defined for the MAC entity concerned as defined in IEEE 802.1w clause 6.5.  The value is determined dynamically; that is it is re-evaluated whenever the value of adminPointToPointMAC changes and whenever the specific procedures defined for the MAC entity evaluate a change in its point-to-point status., SELECTION: false(2)/true(1)`
	MaxAge                      int32  `DESCRIPTION: The value that all bridges use for MaxAge as advertised by the root bridge.  Note that 802.1D-1998 specifies that the range for this parameter is related to the value of BridgeHelloTime.  The granularity of this timer is specified by 802.1D-1998 to be 1 second.  An agent may return a badValue error if a set is attempted to a value that is not a whole number of seconds.`
	HelloTime                   int32  `DESCRIPTION: The value that all bridges use for HelloTime as advertised by the root bridge.  The granularity of this timer is specified by 802.1D-1998 to be 1 second.  An agent may return a badValue error if a set is attempted    to a value that is not a whole number of seconds.`
	ForwardDelay                int32  `DESCRIPTION: The value that all bridges use for ForwardDelay as advertised by the root bridge.  Note that 802.1D-1998 specifies that the range for this parameter is related to the value of dot1dStpBridgeMaxAge.  The granularity of this timer is specified by 802.1D-1998 to be 1 second.  An agent may return a badValue error if a set is attempted to a value that is not a whole number of seconds.`
	BridgeAssurance             int32  `DESCRIPTION: Used to make sure that a neighboring switch does not malfunction  and begin forwarding frames when it should not.  It does this by monitoring receipt of BPDUs on point-to-point links.  When the  BPDUs stop being received the port is put into blocking state  (actually a port inconsistent state which stops forwarding).   When BPDUs restart the port resumes normal RSTP or MST modes.   This handles unidirectional links as well as the malfunction of a  neighboring switch where STP stops sending BPDUs but the switch  continues to forward frames. , SELECTION: false(2)/true(1)`
	BridgeAssuranceInconsistant int32  `DESCRIPTION: When port stops receiving BPDU on a Bridge Assurance enabled port then this will be set., SELECTION: false(2)/true(1)`
	BpduGuard                   int32  `DESCRIPTION: Used in conjuction with AdminEdge to shutdown a port when a BPDU is received.  Protects against loops in the network, SELECTION: false(2)/true(1)`
	BpduGuardInterval           int32  `DESCRIPTION: The interval time to which a port will try to recover from BPDU Guard err-disable state.  If no BPDU frames are detected after this timeout plus 3 Times Hello Time then the port will transition back to Up state.  If condition is cleared manually then this operation is ignored.  If set to zero then timer is inactive and recovery is based on manual intervention., DEFAULT: 30`
	BpduGuardDetected           int32  `DESCRIPTION: Indicates whether a BPDU frame was received on this STP port if the port  is and Edge Port and BPDU Guard is enabled, SELECTION: false(2)/true(1)`
	StpInPkts                   uint64 `DESCRIPTION: Number of STP PDUs received`
	StpOutPkts                  uint64 `DESCRIPTION: Number of STP BPDUs transmitted`
	RstpInPkts                  uint64 `DESCRIPTION: Number of RSTP BPDUs received`
	RstpOutPkts                 uint64 `DESCRIPTION: Number of RSTP BPDUs transmitted`
	TcInPkts                    uint64 `DESCRIPTION: Number of TC BPDUs received`
	TcOutPkts                   uint64 `DESCRIPTION: Number of TC BPDUs transmitted`
	TcAckInPkts                 uint64 `DESCRIPTION: Number of TC Ack BPDUs received`
	TcAckOutPkts                uint64 `DESCRIPTION: Number of TC Ack BPDUs transmitted`
	PvstInPkts                  uint64 `DESCRIPTION: Number of PVST BPDUs received`
	PvstOutPkts                 uint64 `DESCRIPTION: Number of PVST BPDUs transmitted`
	BpduInPkts                  uint64 `DESCRIPTION: Number of BPDUs received`
	BpduOutPkts                 uint64 `DESCRIPTION: Number of BPDUs transmitted`
	PimPrevState                string `DESCRIPTION: PIM previous fsm state`
	PimCurrState                string `DESCRIPTION: PIM current fsm state`
	PrtmPrevState               string `DESCRIPTION: PRTM previous fsm state`
	PrtmCurrState               string `DESCRIPTION: PRTM current fsm state`
	PrxmPrevState               string `DESCRIPTION: PRXM previous fsm state`
	PrxmCurrState               string `DESCRIPTION: PRXM current fsm state`
	PstmPrevState               string `DESCRIPTION: PSTM previous fsm state`
	PstmCurrState               string `DESCRIPTION: PSTM current fsm state`
	TcmPrevState                string `DESCRIPTION: TCM previous fsm state`
	TcmCurrState                string `DESCRIPTION: TCM current fsm state`
	PpmPrevState                string `DESCRIPTION: PPM previous fsm state`
	PpmCurrState                string `DESCRIPTION: PPM current fsm state`
	PtxmPrevState               string `DESCRIPTION: PTXM previous fsm state`
	PtxmCurrState               string `DESCRIPTION: PTXM current fsm state`
	PtimPrevState               string `DESCRIPTION: PTIM previous fsm state`
	PtimCurrState               string `DESCRIPTION: PTIM current fsm state`
	BdmPrevState                string `DESCRIPTION: BDM previous fsm state`
	BdmCurrState                string `DESCRIPTION: BDM current fsm state`
	EdgeDelayWhile              int32  `DESCRIPTION: The Edge Delay timer. The time remaining in the absence of a received BPDU before this port is identified as an operEdgePort.`
	FdWhile                     int32  `DESCRIPTION: The Forward Delay timer. Used to delay Port State transitions until other Bridges have received spanning tree information`
	HelloWhen                   int32  `DESCRIPTION: The Hello timer. Used to ensure that at least one BPDU is transmitted by a Designated Port in each HelloTime period.`
	MdelayWhile                 int32  `DESCRIPTION: The Migration Delay timer. Used by the Port Protocol Migration state machine to allow time for another RSTP Bridge on the same LAN to synchronize its migration state with this Port before the receipt of a BPDU can cause this Port to change the BPDU types it transmits. Initialized to MigrateTime (17.13.9).`
	RbWhile                     int32  `DESCRIPTION: The Recent Backup timer. Maintained at its initial value twice HelloTime while the Port is a Backup Port.`
	RcvdInfoWhile               int32  `DESCRIPTION: The Received Info timer. The time remaining before the spanning tree information received by this Port [portPriority (17.19.21) and portTimes (17.19.22)] is aged out if not refreshed by the receipt of a further Configuration Message.`
	RrWhile                     int32  `DESCRIPTION: The Recent Root timer.`
	TcWhile                     int32  `DESCRIPTION: The Topology Change timer. TCN Messages are sent while this timer is running`
	BaWhile                     int32  `DESCRIPTION: Bridge Assurance timer 3 * Hello Timer`
}

type StpBridgeInstance struct {
	baseObj
	Vlan         uint16 `SNAPROUTE: "KEY",  ACCESS:"rw",  MULTIPLICITY:"*", AUTODISCOVER: "true", DESCRIPTION: Each bridge is associated with a domain.  Typically this domain is represented as the vlan; The default domain is 4095, MIN: "1" ,  MAX: "4095"`
	Address      string `DESCRIPTION: The bridge identifier of the root of the spanning tree as determined by the Spanning Tree Protocol as executed by this node.  This value is used as the Root Identifier parameter in all Configuration Bridge PDUs originated by this node.,  DEFAULT: "00:00:00:00:00:00"`
	Priority     int32  `DESCRIPTION: The value of the write-able portion of the Bridge ID i.e. the first two octets of the 8 octet long Bridge ID.  The other last 6 octets of the Bridge ID are given by the value of Address. On bridges supporting IEEE 802.1t or IEEE 802.1w permissible values are 0-61440 in steps of 4096.  Extended Priority is enabled when the lower 12 bits are set using the Bridges VLAN id, MIN: "0" ,  MAX: "65535", DEFAULT: 32768`
	MaxAge       int32  `DESCRIPTION: The value that all bridges use for MaxAge when this bridge is acting as the root.  Note that 802.1D-1998 specifies that the range for this parameter is related to the value of HelloTime.  The granularity of this timer is specified by 802.1D-1998 to be 1 second.  An agent may return a badValue error if a set is attempted to a value that is not a whole number of seconds., MIN: "6" ,  MAX: "40", DEFAULT: 20`
	HelloTime    int32  `DESCRIPTION: The value that all bridges use for HelloTime when this bridge is acting as the root.  The granularity of this timer is specified by 802.1D-1998 to be 1 second.  An agent may return a badValue error if a set is attempted    to a value that is not a whole number of seconds., MIN: "1" ,  MAX: "2", DEFAULT: 2`
	ForwardDelay int32  `DESCRIPTION: The value that all bridges use for ForwardDelay when this bridge is acting as the root.  Note that 802.1D-1998 specifies that the range for this parameter is related to the value of MaxAge.  The granularity of this timer is specified by 802.1D-1998 to be 1 second.  An agent may return a badValue error if a set is attempted to a value that is not a whole number of seconds., MIN: "3" ,  MAX: "30", DEFAULT: 15`
	ForceVersion int32  `DESCRIPTION: Stp Version, SELECTION: stp(1)/rstp-pvst(2)/mstp(3), DEFAULT: 2`
	TxHoldCount  int32  `DESCRIPTION: Configures the number of BPDUs that can be sent before pausing for 1 second., MIN: "1" ,  MAX: "10", DEFAULT: 6`
}

type StpBridgeInstanceState struct {
	baseObj
	Vlan                    uint16 `SNAPROUTE: "KEY",  ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: Each bridge is associated with a domain.  Typically this domain is represented as the vlan; The default domain is typically 1, MIN: "1" ,  MAX: "4095"`
	IfIndex                 int32  `DESCRIPTION: The value of the instance of the ifIndex object for the bridge, MIN: "1" ,  MAX: "2147483647"`
	Address                 string `DESCRIPTION: The bridge identifier of the root of the spanning tree as determined by the Spanning Tree Protocol as executed by this node.  This value is used as the Root Identifier parameter in all Configuration Bridge PDUs originated by this node."`
	Priority                int32  `DESCRIPTION: The value of the write-able portion of the Bridge ID i.e. the first two octets of the 8 octet long Bridge ID.  The other last 6 octets of the Bridge ID are given by the value of Address. On bridges supporting IEEE 802.1t or IEEE 802.1w permissible values are 0-61440 in steps of 4096., MIN: "0" ,  MAX: "65535"`
	ProtocolSpecification   int32  `DESCRIPTION: An indication of what version of the Spanning Tree Protocol is being run.  The value 'decLb100(2)' indicates the DEC LANbridge 100 Spanning Tree protocol. IEEE 802.1D implementations will return 'ieee8021d(3)'. If future versions of the IEEE Spanning Tree Protocol that are incompatible with the current version are released a new value will be defined., SELECTION: ieee8021d(3)/unknown(1)/decLb100(2)`
	TimeSinceTopologyChange uint32 `DESCRIPTION: The time (in hundredths of a second) since the last time a topology change was detected by the bridge entity. For RSTP this reports the time since the tcWhile timer for any port on this Bridge was nonzero.`
	TopChanges              uint32 `DESCRIPTION: The total number of topology changes detected by this bridge since the management entity was last reset or initialized.`
	DesignatedRoot          string `DESCRIPTION: The bridge identifier of the root of the spanning tree as determined by the Spanning Tree Protocol as executed by this node.  This value is used as the Root Identifier parameter in all Configuration Bridge PDUs originated by this node., LEN : "8"`
	RootCost                int32  `DESCRIPTION: The cost of the path to the root as seen from this bridge.`
	RootPort                int32  `DESCRIPTION: The port number of the port that offers the lowest cost path from this bridge to the root bridge.`
	MaxAge                  int32  `DESCRIPTION: The maximum age of Spanning Tree Protocol information learned from the network on any port before it is discarded in units of hundredths of a second.  This is the actual value that this bridge is currently using.`
	HelloTime               int32  `DESCRIPTION: The amount of time between the transmission of Configuration bridge PDUs by this node on any port when it is the root of the spanning tree or trying to become so in units of hundredths of a second.  This is the actual value that this bridge is currently using.`
	HoldTime                int32  `DESCRIPTION: This time value determines the interval length during which no more than two Configuration bridge PDUs shall be transmitted by this node in units of hundredths of a second.`
	ForwardDelay            int32  `DESCRIPTION: This time value measured in units of hundredths of a second controls how fast a port changes its spanning state when moving towards the Forwarding state.  The value determines how long the port stays in each of the Listening and Learning states which precede the Forwarding state.  This value is also used when a topology change has been detected and is underway to age all dynamic entries in the Forwarding Database. [Note that this value is the one that this bridge is currently using in contrast to ForwardDelay which is the value that this bridge and all others would start using if/when this bridge were to become the root.]`
	BridgeMaxAge            int32  `DESCRIPTION: The maximum age of Spanning Tree Protocol information learned from the network on any port before it is discarded in units of hundredths of a second.  This is the provisioned value of the local bridge.`
	BridgeHelloTime         int32  `DESCRIPTION: The amount of time between the transmission of Configuration bridge PDUs by this node on any port when it is the root of the spanning tree or trying to become so in units of hundredths of a second.  This is the provisioned value of the local bridge   .`
	BridgeHoldTime          int32  `DESCRIPTION: This time value determines the interval length during which no more than two Configuration bridge PDUs shall be transmitted by this node in units of hundredths of a second. This is the provisioned value of the local bridge`
	BridgeForwardDelay      int32  `DESCRIPTION: This time value measured in units of hundredths of a second controls how fast a port changes its spanning state when moving towards the Forwarding state.  The value determines how long the port stays in each of the Listening and Learning states which precede the Forwarding state.  This value is also used when a topology change has been detected and is underway to age all dynamic entries in the Forwarding Database. [Note This is the provisioned value of the local bridge in contrast to ForwardDelay which is the value that this bridge and all others would start using if/when this bridge were to become the root.]`
	TxHoldCount             int32  `DESCRIPTION: TODO`
}

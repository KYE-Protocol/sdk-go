// KYE Protocol™ — public vocabulary constants.
//
// Mirrors public/vocabulary/ in the spec repo. These values are stable
// across the v1.x series and used as discriminators in every KYE
// Protocol™ schema.
package kye

// EntityClass* are the canonical entity-class constants used in KYE URNs.
const (
	EntityClassHuman      = "human"
	EntityClassBusiness   = "business"
	EntityClassAgent      = "agent"
	EntityClassService    = "service"
	EntityClassModel      = "model"
	EntityClassTool       = "tool"
	EntityClassWorkflow   = "workflow"
	EntityClassResource   = "resource"
	EntityClassCapability = "capability"
	EntityClassCredential = "credential"
	EntityClassTrust      = "trust"
	EntityClassCard       = "card"
	EntityClassWallet     = "wallet"
)

// EntityClasses is the canonical list of public entity classes.
var EntityClasses = []string{
	EntityClassHuman,
	EntityClassBusiness,
	EntityClassAgent,
	EntityClassService,
	EntityClassModel,
	EntityClassTool,
	EntityClassWorkflow,
	EntityClassResource,
	EntityClassCapability,
	EntityClassCredential,
	EntityClassTrust,
	EntityClassCard,
	EntityClassWallet,
}

// Decision* are the three canonical decision codes returned by
// POST /v1/runtime/authorize.
const (
	DecisionAllowWithConstraints = "allow_with_constraints"
	DecisionRequireApproval      = "require_approval"
	DecisionDeny                 = "deny"
)

// Signal* are the cascade-bus signal types.
const (
	SignalStop       = "stop"
	SignalQuarantine = "quarantine"
	SignalRevoke     = "revoke"
	SignalRestore    = "restore"
	SignalReplay     = "replay"
)

// State dimensions composed at every authorize call.
const (
	StateLifecycle  = "lifecycle"
	StateAuthority  = "authority"
	StateDelegation = "delegation"
	StateCredential = "credential"
	StateRecovery   = "recovery"
	StateRisk       = "risk"
)

// Conformance* are the 5-tier KYE Conformance & Certification Programme™
// ladder values.
const (
	ConformanceL0Declared      = "L0_declared"
	ConformanceL1SelfTested    = "L1_self_tested"
	ConformanceL2SelfAttested  = "L2_self_attested"
	ConformanceL3Conformant    = "L3_conformant"
	ConformanceL4Certified     = "L4_certified"
)

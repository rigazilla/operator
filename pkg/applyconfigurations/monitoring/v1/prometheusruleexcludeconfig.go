// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// PrometheusRuleExcludeConfigApplyConfiguration represents an declarative configuration of the PrometheusRuleExcludeConfig type for use
// with apply.
type PrometheusRuleExcludeConfigApplyConfiguration struct {
	RuleNamespace *string `json:"ruleNamespace,omitempty"`
	RuleName      *string `json:"ruleName,omitempty"`
}

// PrometheusRuleExcludeConfigApplyConfiguration constructs an declarative configuration of the PrometheusRuleExcludeConfig type for use with
// apply.
func PrometheusRuleExcludeConfig() *PrometheusRuleExcludeConfigApplyConfiguration {
	return &PrometheusRuleExcludeConfigApplyConfiguration{}
}

// WithRuleNamespace sets the RuleNamespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RuleNamespace field is set to the value of the last call.
func (b *PrometheusRuleExcludeConfigApplyConfiguration) WithRuleNamespace(value string) *PrometheusRuleExcludeConfigApplyConfiguration {
	b.RuleNamespace = &value
	return b
}

// WithRuleName sets the RuleName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RuleName field is set to the value of the last call.
func (b *PrometheusRuleExcludeConfigApplyConfiguration) WithRuleName(value string) *PrometheusRuleExcludeConfigApplyConfiguration {
	b.RuleName = &value
	return b
}

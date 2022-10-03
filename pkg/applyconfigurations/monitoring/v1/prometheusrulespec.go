// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// PrometheusRuleSpecApplyConfiguration represents an declarative configuration of the PrometheusRuleSpec type for use
// with apply.
type PrometheusRuleSpecApplyConfiguration struct {
	Groups []RuleGroupApplyConfiguration `json:"groups,omitempty"`
}

// PrometheusRuleSpecApplyConfiguration constructs an declarative configuration of the PrometheusRuleSpec type for use with
// apply.
func PrometheusRuleSpec() *PrometheusRuleSpecApplyConfiguration {
	return &PrometheusRuleSpecApplyConfiguration{}
}

// WithGroups adds the given value to the Groups field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Groups field.
func (b *PrometheusRuleSpecApplyConfiguration) WithGroups(values ...*RuleGroupApplyConfiguration) *PrometheusRuleSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithGroups")
		}
		b.Groups = append(b.Groups, *values[i])
	}
	return b
}
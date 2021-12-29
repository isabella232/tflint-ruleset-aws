// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsNetworkfirewallFirewallInvalidDescriptionRule checks the pattern is valid
type AwsNetworkfirewallFirewallInvalidDescriptionRule struct {
	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsNetworkfirewallFirewallInvalidDescriptionRule returns new rule with default attributes
func NewAwsNetworkfirewallFirewallInvalidDescriptionRule() *AwsNetworkfirewallFirewallInvalidDescriptionRule {
	return &AwsNetworkfirewallFirewallInvalidDescriptionRule{
		resourceType:  "aws_networkfirewall_firewall",
		attributeName: "description",
		max:           512,
		pattern:       regexp.MustCompile(`^.*$`),
	}
}

// Name returns the rule name
func (r *AwsNetworkfirewallFirewallInvalidDescriptionRule) Name() string {
	return "aws_networkfirewall_firewall_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsNetworkfirewallFirewallInvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsNetworkfirewallFirewallInvalidDescriptionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsNetworkfirewallFirewallInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsNetworkfirewallFirewallInvalidDescriptionRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"description must be 512 characters or less",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^.*$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsWafv2IPSetInvalidIPAddressVersionRule checks the pattern is valid
type AwsWafv2IPSetInvalidIPAddressVersionRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsWafv2IPSetInvalidIPAddressVersionRule returns new rule with default attributes
func NewAwsWafv2IPSetInvalidIPAddressVersionRule() *AwsWafv2IPSetInvalidIPAddressVersionRule {
	return &AwsWafv2IPSetInvalidIPAddressVersionRule{
		resourceType:  "aws_wafv2_ip_set",
		attributeName: "ip_address_version",
		enum: []string{
			"IPV4",
			"IPV6",
		},
	}
}

// Name returns the rule name
func (r *AwsWafv2IPSetInvalidIPAddressVersionRule) Name() string {
	return "aws_wafv2_ip_set_invalid_ip_address_version"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsWafv2IPSetInvalidIPAddressVersionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsWafv2IPSetInvalidIPAddressVersionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsWafv2IPSetInvalidIPAddressVersionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsWafv2IPSetInvalidIPAddressVersionRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" is an invalid value as ip_address_version`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
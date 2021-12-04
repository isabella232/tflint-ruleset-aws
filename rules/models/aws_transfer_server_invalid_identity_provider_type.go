// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsTransferServerInvalidIdentityProviderTypeRule checks the pattern is valid
type AwsTransferServerInvalidIdentityProviderTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsTransferServerInvalidIdentityProviderTypeRule returns new rule with default attributes
func NewAwsTransferServerInvalidIdentityProviderTypeRule() *AwsTransferServerInvalidIdentityProviderTypeRule {
	return &AwsTransferServerInvalidIdentityProviderTypeRule{
		resourceType:  "aws_transfer_server",
		attributeName: "identity_provider_type",
		enum: []string{
			"SERVICE_MANAGED",
			"API_GATEWAY",
			"AWS_DIRECTORY_SERVICE",
			"AWS_LAMBDA",
		},
	}
}

// Name returns the rule name
func (r *AwsTransferServerInvalidIdentityProviderTypeRule) Name() string {
	return "aws_transfer_server_invalid_identity_provider_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsTransferServerInvalidIdentityProviderTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsTransferServerInvalidIdentityProviderTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsTransferServerInvalidIdentityProviderTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsTransferServerInvalidIdentityProviderTypeRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as identity_provider_type`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}

// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsStoragegatewayNfsFileShareInvalidSquashRule checks the pattern is valid
type AwsStoragegatewayNfsFileShareInvalidSquashRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsStoragegatewayNfsFileShareInvalidSquashRule returns new rule with default attributes
func NewAwsStoragegatewayNfsFileShareInvalidSquashRule() *AwsStoragegatewayNfsFileShareInvalidSquashRule {
	return &AwsStoragegatewayNfsFileShareInvalidSquashRule{
		resourceType:  "aws_storagegateway_nfs_file_share",
		attributeName: "squash",
		max:           15,
		min:           5,
	}
}

// Name returns the rule name
func (r *AwsStoragegatewayNfsFileShareInvalidSquashRule) Name() string {
	return "aws_storagegateway_nfs_file_share_invalid_squash"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsStoragegatewayNfsFileShareInvalidSquashRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsStoragegatewayNfsFileShareInvalidSquashRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsStoragegatewayNfsFileShareInvalidSquashRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsStoragegatewayNfsFileShareInvalidSquashRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"squash must be 15 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"squash must be 5 characters or higher",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
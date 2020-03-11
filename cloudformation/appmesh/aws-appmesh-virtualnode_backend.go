package appmesh

import (
	"github.com/fmechant/goformation/v4/cloudformation/policies"
)

// VirtualNode_Backend AWS CloudFormation Resource (AWS::AppMesh::VirtualNode.Backend)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-backend.html
type VirtualNode_Backend struct {

	// VirtualService AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualnode-backend.html#cfn-appmesh-virtualnode-backend-virtualservice
	VirtualService *VirtualNode_VirtualServiceBackend `json:"VirtualService,omitempty"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationUpdateReplacePolicy represents a CloudFormation UpdateReplacePolicy
	AWSCloudFormationUpdateReplacePolicy policies.UpdateReplacePolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`

	// AWSCloudFormationCondition stores the logical ID of the condition that must be satisfied for this resource to be created
	AWSCloudFormationCondition string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *VirtualNode_Backend) AWSCloudFormationType() string {
	return "AWS::AppMesh::VirtualNode.Backend"
}

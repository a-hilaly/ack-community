package gencopy

import (
	ackv1alpha1 "github.com/aws/aws-controllers-k8s/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type CreateBucketConfiguration struct {
	LocationConstraint *string `json:"locationConstraint,omitempty"`
}

type BucketSpec struct {
	ACL                        *string                    `json:"acl,omitempty"`
	CreateBucketConfiguration  *CreateBucketConfiguration `json:"createBucketConfiguration,omitempty"`
	GrantFullControl           *string                    `json:"grantFullControl,omitempty"`
	GrantRead                  *string                    `json:"grantRead,omitempty"`
	GrantReadACP               *string                    `json:"grantReadACP,omitempty"`
	GrantWrite                 *string                    `json:"grantWrite,omitempty"`
	GrantWriteACP              *string                    `json:"grantWriteACP,omitempty"`
	Name                       *string                    `json:"name"`
	ObjectLockEnabledForBucket *bool                      `json:"objectLockEnabledForBucket,omitempty"`
}

type BucketStatus struct {
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	Conditions          []*ackv1alpha1.Condition      `json:"conditions"`
	Location            *string                       `json:"location,omitempty"`
}

type Bucket struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BucketSpec   `json:"spec,omitempty"`
	Status            BucketStatus `json:"status,omitempty"`
}

func (b1 *Bucket) Equal(b2 *Bucket) (bool, []string) {
	c := Comparator{}
	c.stringEqualSoft("Spec.ACL", b1.Spec.ACL, b2.Spec.ACL)
	if b1.Spec.CreateBucketConfiguration == nil {

	}
	c.stringEqualSoft("Spec.GrantFullControl", b1.Spec.GrantFullControl, b2.Spec.GrantFullControl)
	c.stringEqualSoft("Spec.GrantRead", b1.Spec.GrantRead, b2.Spec.GrantRead)
	c.stringEqualSoft("Spec.GrandReacACP", b1.Spec.GrantReadACP, b2.Spec.GrantReadACP)
	c.stringEqualSoft("Spec.ACL", b1.Spec.GrantWrite, b2.Spec.GrantWrite)
	c.stringEqualSoft("Spec.ACL", b1.Spec.GrantReadACP, b2.Spec.GrantWriteACP)
	c.stringEqualStrict("Spec.Name", b1.Spec.Name, b2.Spec.Name)

	diffs := c.Diffs()
	return len(diffs) == 0, diffs
}

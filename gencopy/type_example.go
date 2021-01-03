package gencopy

import (
	ackv1alpha1 "github.com/aws/aws-controllers-k8s/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type SubObjectType struct {
	StringType *string
}

type ObjectSpec struct {
	StringType *string
	SliceType  []*string
	MapType    map[string]*string
	SubType    *SubObjectType
}

type ObjectStatus struct {
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata
	Conditions          []*ackv1alpha1.Condition
	Location            *string
}

type Object struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   ObjectSpec
	Status ObjectStatus
}

/* func (b1 *Object) Equal(b2 *Object) (bool, []string) {
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
} */

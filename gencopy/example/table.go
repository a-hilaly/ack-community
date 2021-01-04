package example

import (
	ackv1alpha1 "github.com/aws/aws-controllers-k8s/apis/core/v1alpha1"
	"github.com/aws/aws-controllers-k8s/gencopy/pkg/compare"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type TableSpec struct {
	BillingMode      *string             `json:"billingMode,omitempty"`
	KeySchema        []*KeySchemaElement `json:"keySchema"`
	SSESpecification *SSESpecification   `json:"sseSpecification,omitempty"`
	TableName        *string             `json:"tableName"`
}

type TableStatus struct {
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	Conditions          []*ackv1alpha1.Condition      `json:"conditions"`
	ArchivalSummary     *ArchivalSummary              `json:"archivalSummary,omitempty"`
	BillingModeSummary  *BillingModeSummary           `json:"billingModeSummary,omitempty"`
	CreationDateTime    *metav1.Time                  `json:"creationDateTime,omitempty"`
	GlobalTableVersion  *string                       `json:"globalTableVersion,omitempty"`
	ItemCount           *int64                        `json:"itemCount,omitempty"`
	LatestStreamARN     *string                       `json:"latestStreamARN,omitempty"`
	LatestStreamLabel   *string                       `json:"latestStreamLabel,omitempty"`
	Replicas            []*ReplicaDescription         `json:"replicas,omitempty"`
	RestoreSummary      *RestoreSummary               `json:"restoreSummary,omitempty"`
	SSEDescription      *SSEDescription               `json:"sseDescription,omitempty"`
	TableID             *string                       `json:"tableID,omitempty"`
	TableSizeBytes      *int64                        `json:"tableSizeBytes,omitempty"`
	TableStatus         *string                       `json:"tableStatus,omitempty"`
}

type Table struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TableSpec   `json:"spec,omitempty"`
	Status            TableStatus `json:"status,omitempty"`
}

// These are the functions we should generation
func (t1 *Table) Diffs(t2 *Table) []string {
	c := compare.New()

	// Compare simple fields
	c.StringEqualStrict("Spec.TableName", t1.Spec.TableName, t1.Spec.TableName)
	c.StringEqualSoft("Spec.BillingMode", t1.Spec.BillingMode, t1.Spec.BillingMode)

	// Compare embedded structs
	{
		// soft compare
		{
			if t1.Spec.SSESpecification == nil && t2.Spec.SSESpecification != nil {
				c.Append(t2.Spec.SSESpecification.Diffs(&SSESpecification{})...)
			} else if t2.Spec.SSESpecification == nil && t1.Spec.SSESpecification != nil {
				c.Append(t1.Spec.SSESpecification.Diffs(&SSESpecification{})...)
			} else if t1.Spec.SSESpecification != nil && t2.Spec.SSESpecification != nil {
				c.Append(t1.Spec.SSESpecification.Diffs(t2.Spec.SSESpecification)...)
			}
		}

		// strict
		if (t1.Spec.SSESpecification == nil && t2.Spec.SSESpecification != nil) ||
			(t2.Spec.SSESpecification == nil && t1.Spec.SSESpecification != nil) {
			c.Append("Spec.SSESpecification")
		} else if t1.Spec.SSESpecification != nil && t2.Spec.SSESpecification != nil {

		}

		// soft
		if t1.Spec.SSESpecification != nil || t2.Spec.SSESpecification != nil {
			c.Append(t1.Spec.SSESpecification.Diffs(t2.Spec.SSESpecification)...)
		}
	}

	// Compare array fields
	if len(t1.Spec.KeySchema) != len(t2.Spec.KeySchema) {
		c.Append("Spec.KeySchema")
	} else {
		for _, e1 := range t1.Spec.KeySchema {
			foundEqual := false
			for _, e2 := range t2.Spec.KeySchema {
				if len(e1.Diffs(e2)) == 0 {
					foundEqual = true
					break
				}
			}
			if !foundEqual {
				c.Append("Spec.KeySchema")
				break
			}
		}
	}

	// Compare map fields
	// TODO

	return c.Diffs()
}

func (sses1 *SSESpecification) Diffs(sses2 *SSESpecification) []string {
	c := compare.New()

	c.BoolEqualStrict("SSESpecification.Enabled", sses1.Enabled, sses1.Enabled)
	c.StringEqualStrict("SSESpecification.AttributeName", sses1.SSEType, sses2.SSEType)
	c.StringEqualStrict("SSESpecification.KMSMasterKeyID", sses1.KMSMasterKeyID, sses2.KMSMasterKeyID)

	return c.Diffs()
}

func (kse1 *KeySchemaElement) Diffs(kse2 *KeySchemaElement) []string {
	c := compare.New()

	c.StringEqualStrict("KeySchemaElement.KeyType", kse1.KeyType, kse2.KeyType)
	c.StringEqualStrict("KeySchemaElement.AttributeName", kse1.AttributeName, kse2.AttributeName)

	return c.Diffs()
}

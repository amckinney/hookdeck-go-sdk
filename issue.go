// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
)

// Issue
type Issue struct {
	typeName            string
	DeliveryIssue       *DeliveryIssue
	TransformationIssue *TransformationIssue
}

func NewIssueFromDeliveryIssue(value *DeliveryIssue) *Issue {
	return &Issue{typeName: "deliveryIssue", DeliveryIssue: value}
}

func NewIssueFromTransformationIssue(value *TransformationIssue) *Issue {
	return &Issue{typeName: "transformationIssue", TransformationIssue: value}
}

func (i *Issue) UnmarshalJSON(data []byte) error {
	valueDeliveryIssue := new(DeliveryIssue)
	if err := json.Unmarshal(data, &valueDeliveryIssue); err == nil {
		i.typeName = "deliveryIssue"
		i.DeliveryIssue = valueDeliveryIssue
		return nil
	}
	valueTransformationIssue := new(TransformationIssue)
	if err := json.Unmarshal(data, &valueTransformationIssue); err == nil {
		i.typeName = "transformationIssue"
		i.TransformationIssue = valueTransformationIssue
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, i)
}

func (i Issue) MarshalJSON() ([]byte, error) {
	switch i.typeName {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", i.typeName, i)
	case "deliveryIssue":
		return json.Marshal(i.DeliveryIssue)
	case "transformationIssue":
		return json.Marshal(i.TransformationIssue)
	}
}

type IssueVisitor interface {
	VisitDeliveryIssue(*DeliveryIssue) error
	VisitTransformationIssue(*TransformationIssue) error
}

func (i *Issue) Accept(v IssueVisitor) error {
	switch i.typeName {
	default:
		return fmt.Errorf("invalid type %s in %T", i.typeName, i)
	case "deliveryIssue":
		return v.VisitDeliveryIssue(i.DeliveryIssue)
	case "transformationIssue":
		return v.VisitTransformationIssue(i.TransformationIssue)
	}
}

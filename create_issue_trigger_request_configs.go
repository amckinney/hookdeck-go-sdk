// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
)

// Configuration object for the specific issue type selected
type CreateIssueTriggerRequestConfigs struct {
	typeName                          string
	IssueTriggerDeliveryConfigs       *IssueTriggerDeliveryConfigs
	IssueTriggerTransformationConfigs *IssueTriggerTransformationConfigs
	IssueTriggerBackpressureConfigs   *IssueTriggerBackpressureConfigs
}

func NewCreateIssueTriggerRequestConfigsFromIssueTriggerDeliveryConfigs(value *IssueTriggerDeliveryConfigs) *CreateIssueTriggerRequestConfigs {
	return &CreateIssueTriggerRequestConfigs{typeName: "issueTriggerDeliveryConfigs", IssueTriggerDeliveryConfigs: value}
}

func NewCreateIssueTriggerRequestConfigsFromIssueTriggerTransformationConfigs(value *IssueTriggerTransformationConfigs) *CreateIssueTriggerRequestConfigs {
	return &CreateIssueTriggerRequestConfigs{typeName: "issueTriggerTransformationConfigs", IssueTriggerTransformationConfigs: value}
}

func NewCreateIssueTriggerRequestConfigsFromIssueTriggerBackpressureConfigs(value *IssueTriggerBackpressureConfigs) *CreateIssueTriggerRequestConfigs {
	return &CreateIssueTriggerRequestConfigs{typeName: "issueTriggerBackpressureConfigs", IssueTriggerBackpressureConfigs: value}
}

func (c *CreateIssueTriggerRequestConfigs) UnmarshalJSON(data []byte) error {
	valueIssueTriggerDeliveryConfigs := new(IssueTriggerDeliveryConfigs)
	if err := json.Unmarshal(data, &valueIssueTriggerDeliveryConfigs); err == nil {
		c.typeName = "issueTriggerDeliveryConfigs"
		c.IssueTriggerDeliveryConfigs = valueIssueTriggerDeliveryConfigs
		return nil
	}
	valueIssueTriggerTransformationConfigs := new(IssueTriggerTransformationConfigs)
	if err := json.Unmarshal(data, &valueIssueTriggerTransformationConfigs); err == nil {
		c.typeName = "issueTriggerTransformationConfigs"
		c.IssueTriggerTransformationConfigs = valueIssueTriggerTransformationConfigs
		return nil
	}
	valueIssueTriggerBackpressureConfigs := new(IssueTriggerBackpressureConfigs)
	if err := json.Unmarshal(data, &valueIssueTriggerBackpressureConfigs); err == nil {
		c.typeName = "issueTriggerBackpressureConfigs"
		c.IssueTriggerBackpressureConfigs = valueIssueTriggerBackpressureConfigs
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, c)
}

func (c CreateIssueTriggerRequestConfigs) MarshalJSON() ([]byte, error) {
	switch c.typeName {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", c.typeName, c)
	case "issueTriggerDeliveryConfigs":
		return json.Marshal(c.IssueTriggerDeliveryConfigs)
	case "issueTriggerTransformationConfigs":
		return json.Marshal(c.IssueTriggerTransformationConfigs)
	case "issueTriggerBackpressureConfigs":
		return json.Marshal(c.IssueTriggerBackpressureConfigs)
	}
}

type CreateIssueTriggerRequestConfigsVisitor interface {
	VisitIssueTriggerDeliveryConfigs(*IssueTriggerDeliveryConfigs) error
	VisitIssueTriggerTransformationConfigs(*IssueTriggerTransformationConfigs) error
	VisitIssueTriggerBackpressureConfigs(*IssueTriggerBackpressureConfigs) error
}

func (c *CreateIssueTriggerRequestConfigs) Accept(v CreateIssueTriggerRequestConfigsVisitor) error {
	switch c.typeName {
	default:
		return fmt.Errorf("invalid type %s in %T", c.typeName, c)
	case "issueTriggerDeliveryConfigs":
		return v.VisitIssueTriggerDeliveryConfigs(c.IssueTriggerDeliveryConfigs)
	case "issueTriggerTransformationConfigs":
		return v.VisitIssueTriggerTransformationConfigs(c.IssueTriggerTransformationConfigs)
	case "issueTriggerBackpressureConfigs":
		return v.VisitIssueTriggerBackpressureConfigs(c.IssueTriggerBackpressureConfigs)
	}
}

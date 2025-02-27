//
// Automatically generated. DO NOT EDIT.
//

package resources

import (
	"encoding/json"
	"math/big"

	"github.com/Skilldus/contrail-go-api"
)

const (
	security_logging_object_security_logging_object_rules int = iota
	security_logging_object_security_logging_object_rate
	security_logging_object_id_perms
	security_logging_object_perms2
	security_logging_object_annotations
	security_logging_object_display_name
	security_logging_object_network_policy_refs
	security_logging_object_security_group_refs
	security_logging_object_tag_refs
	security_logging_object_virtual_network_back_refs
	security_logging_object_virtual_machine_interface_back_refs
	security_logging_object_firewall_policy_back_refs
	security_logging_object_firewall_rule_back_refs
)

type SecurityLoggingObject struct {
	contrail.ObjectBase
	security_logging_object_rules       SecurityLoggingObjectRuleListType
	security_logging_object_rate        int
	id_perms                            IdPermsType
	perms2                              PermType2
	annotations                         KeyValuePairs
	display_name                        string
	network_policy_refs                 contrail.ReferenceList
	security_group_refs                 contrail.ReferenceList
	tag_refs                            contrail.ReferenceList
	virtual_network_back_refs           contrail.ReferenceList
	virtual_machine_interface_back_refs contrail.ReferenceList
	firewall_policy_back_refs           contrail.ReferenceList
	firewall_rule_back_refs             contrail.ReferenceList
	valid                               big.Int
	modified                            big.Int
	baseMap                             map[string]contrail.ReferenceList
}

func (obj *SecurityLoggingObject) GetType() string {
	return "security-logging-object"
}

func (obj *SecurityLoggingObject) GetDefaultParent() []string {
	name := []string{"default-global-system-config", "default-global-vrouter-config"}
	return name
}

func (obj *SecurityLoggingObject) GetDefaultParentType() string {
	return "global-vrouter-config"
}

func (obj *SecurityLoggingObject) SetName(name string) {
	obj.VSetName(obj, name)
}

func (obj *SecurityLoggingObject) SetParent(parent contrail.IObject) {
	obj.VSetParent(obj, parent)
}

func (obj *SecurityLoggingObject) storeReferenceBase(
	name string, refList contrail.ReferenceList) {
	if obj.baseMap == nil {
		obj.baseMap = make(map[string]contrail.ReferenceList)
	}
	refCopy := make(contrail.ReferenceList, len(refList))
	copy(refCopy, refList)
	obj.baseMap[name] = refCopy
}

func (obj *SecurityLoggingObject) hasReferenceBase(name string) bool {
	if obj.baseMap == nil {
		return false
	}
	_, exists := obj.baseMap[name]
	return exists
}

func (obj *SecurityLoggingObject) UpdateDone() {
	obj.modified.SetInt64(0)
	obj.baseMap = nil
}

func (obj *SecurityLoggingObject) GetSecurityLoggingObjectRules() SecurityLoggingObjectRuleListType {
	return obj.security_logging_object_rules
}

func (obj *SecurityLoggingObject) SetSecurityLoggingObjectRules(value *SecurityLoggingObjectRuleListType) {
	obj.security_logging_object_rules = *value
	obj.modified.SetBit(&obj.modified, security_logging_object_security_logging_object_rules, 1)
}

func (obj *SecurityLoggingObject) GetSecurityLoggingObjectRate() int {
	return obj.security_logging_object_rate
}

func (obj *SecurityLoggingObject) SetSecurityLoggingObjectRate(value int) {
	obj.security_logging_object_rate = value
	obj.modified.SetBit(&obj.modified, security_logging_object_security_logging_object_rate, 1)
}

func (obj *SecurityLoggingObject) GetIdPerms() IdPermsType {
	return obj.id_perms
}

func (obj *SecurityLoggingObject) SetIdPerms(value *IdPermsType) {
	obj.id_perms = *value
	obj.modified.SetBit(&obj.modified, security_logging_object_id_perms, 1)
}

func (obj *SecurityLoggingObject) GetPerms2() PermType2 {
	return obj.perms2
}

func (obj *SecurityLoggingObject) SetPerms2(value *PermType2) {
	obj.perms2 = *value
	obj.modified.SetBit(&obj.modified, security_logging_object_perms2, 1)
}

func (obj *SecurityLoggingObject) GetAnnotations() KeyValuePairs {
	return obj.annotations
}

func (obj *SecurityLoggingObject) SetAnnotations(value *KeyValuePairs) {
	obj.annotations = *value
	obj.modified.SetBit(&obj.modified, security_logging_object_annotations, 1)
}

func (obj *SecurityLoggingObject) GetDisplayName() string {
	return obj.display_name
}

func (obj *SecurityLoggingObject) SetDisplayName(value string) {
	obj.display_name = value
	obj.modified.SetBit(&obj.modified, security_logging_object_display_name, 1)
}

func (obj *SecurityLoggingObject) readNetworkPolicyRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(security_logging_object_network_policy_refs) == 0) {
		err := obj.GetField(obj, "network_policy_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *SecurityLoggingObject) GetNetworkPolicyRefs() (
	contrail.ReferenceList, error) {
	err := obj.readNetworkPolicyRefs()
	if err != nil {
		return nil, err
	}
	return obj.network_policy_refs, nil
}

func (obj *SecurityLoggingObject) AddNetworkPolicy(
	rhs *NetworkPolicy, data SecurityLoggingObjectRuleListType) error {
	err := obj.readNetworkPolicyRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(security_logging_object_network_policy_refs) == 0 {
		obj.storeReferenceBase("network-policy", obj.network_policy_refs)
	}

	ref := contrail.Reference{
		rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), data}
	obj.network_policy_refs = append(obj.network_policy_refs, ref)
	obj.modified.SetBit(&obj.modified, security_logging_object_network_policy_refs, 1)
	return nil
}

func (obj *SecurityLoggingObject) DeleteNetworkPolicy(uuid string) error {
	err := obj.readNetworkPolicyRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(security_logging_object_network_policy_refs) == 0 {
		obj.storeReferenceBase("network-policy", obj.network_policy_refs)
	}

	for i, ref := range obj.network_policy_refs {
		if ref.Uuid == uuid {
			obj.network_policy_refs = append(
				obj.network_policy_refs[:i],
				obj.network_policy_refs[i+1:]...)
			break
		}
	}
	obj.modified.SetBit(&obj.modified, security_logging_object_network_policy_refs, 1)
	return nil
}

func (obj *SecurityLoggingObject) ClearNetworkPolicy() {
	if (obj.valid.Bit(security_logging_object_network_policy_refs) != 0) &&
		(obj.modified.Bit(security_logging_object_network_policy_refs) == 0) {
		obj.storeReferenceBase("network-policy", obj.network_policy_refs)
	}
	obj.network_policy_refs = make([]contrail.Reference, 0)
	obj.valid.SetBit(&obj.valid, security_logging_object_network_policy_refs, 1)
	obj.modified.SetBit(&obj.modified, security_logging_object_network_policy_refs, 1)
}

func (obj *SecurityLoggingObject) SetNetworkPolicyList(
	refList []contrail.ReferencePair) {
	obj.ClearNetworkPolicy()
	obj.network_policy_refs = make([]contrail.Reference, len(refList))
	for i, pair := range refList {
		obj.network_policy_refs[i] = contrail.Reference{
			pair.Object.GetFQName(),
			pair.Object.GetUuid(),
			pair.Object.GetHref(),
			pair.Attribute,
		}
	}
}

func (obj *SecurityLoggingObject) readSecurityGroupRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(security_logging_object_security_group_refs) == 0) {
		err := obj.GetField(obj, "security_group_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *SecurityLoggingObject) GetSecurityGroupRefs() (
	contrail.ReferenceList, error) {
	err := obj.readSecurityGroupRefs()
	if err != nil {
		return nil, err
	}
	return obj.security_group_refs, nil
}

func (obj *SecurityLoggingObject) AddSecurityGroup(
	rhs *SecurityGroup, data SecurityLoggingObjectRuleListType) error {
	err := obj.readSecurityGroupRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(security_logging_object_security_group_refs) == 0 {
		obj.storeReferenceBase("security-group", obj.security_group_refs)
	}

	ref := contrail.Reference{
		rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), data}
	obj.security_group_refs = append(obj.security_group_refs, ref)
	obj.modified.SetBit(&obj.modified, security_logging_object_security_group_refs, 1)
	return nil
}

func (obj *SecurityLoggingObject) DeleteSecurityGroup(uuid string) error {
	err := obj.readSecurityGroupRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(security_logging_object_security_group_refs) == 0 {
		obj.storeReferenceBase("security-group", obj.security_group_refs)
	}

	for i, ref := range obj.security_group_refs {
		if ref.Uuid == uuid {
			obj.security_group_refs = append(
				obj.security_group_refs[:i],
				obj.security_group_refs[i+1:]...)
			break
		}
	}
	obj.modified.SetBit(&obj.modified, security_logging_object_security_group_refs, 1)
	return nil
}

func (obj *SecurityLoggingObject) ClearSecurityGroup() {
	if (obj.valid.Bit(security_logging_object_security_group_refs) != 0) &&
		(obj.modified.Bit(security_logging_object_security_group_refs) == 0) {
		obj.storeReferenceBase("security-group", obj.security_group_refs)
	}
	obj.security_group_refs = make([]contrail.Reference, 0)
	obj.valid.SetBit(&obj.valid, security_logging_object_security_group_refs, 1)
	obj.modified.SetBit(&obj.modified, security_logging_object_security_group_refs, 1)
}

func (obj *SecurityLoggingObject) SetSecurityGroupList(
	refList []contrail.ReferencePair) {
	obj.ClearSecurityGroup()
	obj.security_group_refs = make([]contrail.Reference, len(refList))
	for i, pair := range refList {
		obj.security_group_refs[i] = contrail.Reference{
			pair.Object.GetFQName(),
			pair.Object.GetUuid(),
			pair.Object.GetHref(),
			pair.Attribute,
		}
	}
}

func (obj *SecurityLoggingObject) readTagRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(security_logging_object_tag_refs) == 0) {
		err := obj.GetField(obj, "tag_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *SecurityLoggingObject) GetTagRefs() (
	contrail.ReferenceList, error) {
	err := obj.readTagRefs()
	if err != nil {
		return nil, err
	}
	return obj.tag_refs, nil
}

func (obj *SecurityLoggingObject) AddTag(
	rhs *Tag) error {
	err := obj.readTagRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(security_logging_object_tag_refs) == 0 {
		obj.storeReferenceBase("tag", obj.tag_refs)
	}

	ref := contrail.Reference{
		rhs.GetFQName(), rhs.GetUuid(), rhs.GetHref(), nil}
	obj.tag_refs = append(obj.tag_refs, ref)
	obj.modified.SetBit(&obj.modified, security_logging_object_tag_refs, 1)
	return nil
}

func (obj *SecurityLoggingObject) DeleteTag(uuid string) error {
	err := obj.readTagRefs()
	if err != nil {
		return err
	}

	if obj.modified.Bit(security_logging_object_tag_refs) == 0 {
		obj.storeReferenceBase("tag", obj.tag_refs)
	}

	for i, ref := range obj.tag_refs {
		if ref.Uuid == uuid {
			obj.tag_refs = append(
				obj.tag_refs[:i],
				obj.tag_refs[i+1:]...)
			break
		}
	}
	obj.modified.SetBit(&obj.modified, security_logging_object_tag_refs, 1)
	return nil
}

func (obj *SecurityLoggingObject) ClearTag() {
	if (obj.valid.Bit(security_logging_object_tag_refs) != 0) &&
		(obj.modified.Bit(security_logging_object_tag_refs) == 0) {
		obj.storeReferenceBase("tag", obj.tag_refs)
	}
	obj.tag_refs = make([]contrail.Reference, 0)
	obj.valid.SetBit(&obj.valid, security_logging_object_tag_refs, 1)
	obj.modified.SetBit(&obj.modified, security_logging_object_tag_refs, 1)
}

func (obj *SecurityLoggingObject) SetTagList(
	refList []contrail.ReferencePair) {
	obj.ClearTag()
	obj.tag_refs = make([]contrail.Reference, len(refList))
	for i, pair := range refList {
		obj.tag_refs[i] = contrail.Reference{
			pair.Object.GetFQName(),
			pair.Object.GetUuid(),
			pair.Object.GetHref(),
			pair.Attribute,
		}
	}
}

func (obj *SecurityLoggingObject) readVirtualNetworkBackRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(security_logging_object_virtual_network_back_refs) == 0) {
		err := obj.GetField(obj, "virtual_network_back_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *SecurityLoggingObject) GetVirtualNetworkBackRefs() (
	contrail.ReferenceList, error) {
	err := obj.readVirtualNetworkBackRefs()
	if err != nil {
		return nil, err
	}
	return obj.virtual_network_back_refs, nil
}

func (obj *SecurityLoggingObject) readVirtualMachineInterfaceBackRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(security_logging_object_virtual_machine_interface_back_refs) == 0) {
		err := obj.GetField(obj, "virtual_machine_interface_back_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *SecurityLoggingObject) GetVirtualMachineInterfaceBackRefs() (
	contrail.ReferenceList, error) {
	err := obj.readVirtualMachineInterfaceBackRefs()
	if err != nil {
		return nil, err
	}
	return obj.virtual_machine_interface_back_refs, nil
}

func (obj *SecurityLoggingObject) readFirewallPolicyBackRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(security_logging_object_firewall_policy_back_refs) == 0) {
		err := obj.GetField(obj, "firewall_policy_back_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *SecurityLoggingObject) GetFirewallPolicyBackRefs() (
	contrail.ReferenceList, error) {
	err := obj.readFirewallPolicyBackRefs()
	if err != nil {
		return nil, err
	}
	return obj.firewall_policy_back_refs, nil
}

func (obj *SecurityLoggingObject) readFirewallRuleBackRefs() error {
	if !obj.IsTransient() &&
		(obj.valid.Bit(security_logging_object_firewall_rule_back_refs) == 0) {
		err := obj.GetField(obj, "firewall_rule_back_refs")
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *SecurityLoggingObject) GetFirewallRuleBackRefs() (
	contrail.ReferenceList, error) {
	err := obj.readFirewallRuleBackRefs()
	if err != nil {
		return nil, err
	}
	return obj.firewall_rule_back_refs, nil
}

func (obj *SecurityLoggingObject) MarshalJSON() ([]byte, error) {
	msg := map[string]*json.RawMessage{}
	err := obj.MarshalCommon(msg)
	if err != nil {
		return nil, err
	}

	if obj.modified.Bit(security_logging_object_security_logging_object_rules) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.security_logging_object_rules)
		if err != nil {
			return nil, err
		}
		msg["security_logging_object_rules"] = &value
	}

	if obj.modified.Bit(security_logging_object_security_logging_object_rate) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.security_logging_object_rate)
		if err != nil {
			return nil, err
		}
		msg["security_logging_object_rate"] = &value
	}

	if obj.modified.Bit(security_logging_object_id_perms) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.id_perms)
		if err != nil {
			return nil, err
		}
		msg["id_perms"] = &value
	}

	if obj.modified.Bit(security_logging_object_perms2) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.perms2)
		if err != nil {
			return nil, err
		}
		msg["perms2"] = &value
	}

	if obj.modified.Bit(security_logging_object_annotations) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.annotations)
		if err != nil {
			return nil, err
		}
		msg["annotations"] = &value
	}

	if obj.modified.Bit(security_logging_object_display_name) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.display_name)
		if err != nil {
			return nil, err
		}
		msg["display_name"] = &value
	}

	if len(obj.network_policy_refs) > 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.network_policy_refs)
		if err != nil {
			return nil, err
		}
		msg["network_policy_refs"] = &value
	}

	if len(obj.security_group_refs) > 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.security_group_refs)
		if err != nil {
			return nil, err
		}
		msg["security_group_refs"] = &value
	}

	if len(obj.tag_refs) > 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.tag_refs)
		if err != nil {
			return nil, err
		}
		msg["tag_refs"] = &value
	}

	return json.Marshal(msg)
}

func (obj *SecurityLoggingObject) UnmarshalJSON(body []byte) error {
	var m map[string]json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	err = obj.UnmarshalCommon(m)
	if err != nil {
		return err
	}
	for key, value := range m {
		switch key {
		case "security_logging_object_rules":
			err = json.Unmarshal(value, &obj.security_logging_object_rules)
			if err == nil {
				obj.valid.SetBit(&obj.valid, security_logging_object_security_logging_object_rules, 1)
			}
			break
		case "security_logging_object_rate":
			err = json.Unmarshal(value, &obj.security_logging_object_rate)
			if err == nil {
				obj.valid.SetBit(&obj.valid, security_logging_object_security_logging_object_rate, 1)
			}
			break
		case "id_perms":
			err = json.Unmarshal(value, &obj.id_perms)
			if err == nil {
				obj.valid.SetBit(&obj.valid, security_logging_object_id_perms, 1)
			}
			break
		case "perms2":
			err = json.Unmarshal(value, &obj.perms2)
			if err == nil {
				obj.valid.SetBit(&obj.valid, security_logging_object_perms2, 1)
			}
			break
		case "annotations":
			err = json.Unmarshal(value, &obj.annotations)
			if err == nil {
				obj.valid.SetBit(&obj.valid, security_logging_object_annotations, 1)
			}
			break
		case "display_name":
			err = json.Unmarshal(value, &obj.display_name)
			if err == nil {
				obj.valid.SetBit(&obj.valid, security_logging_object_display_name, 1)
			}
			break
		case "tag_refs":
			err = json.Unmarshal(value, &obj.tag_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, security_logging_object_tag_refs, 1)
			}
			break
		case "virtual_network_back_refs":
			err = json.Unmarshal(value, &obj.virtual_network_back_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, security_logging_object_virtual_network_back_refs, 1)
			}
			break
		case "virtual_machine_interface_back_refs":
			err = json.Unmarshal(value, &obj.virtual_machine_interface_back_refs)
			if err == nil {
				obj.valid.SetBit(&obj.valid, security_logging_object_virtual_machine_interface_back_refs, 1)
			}
			break
		case "network_policy_refs":
			{
				type ReferenceElement struct {
					To   []string
					Uuid string
					Href string
					Attr SecurityLoggingObjectRuleListType
				}
				var array []ReferenceElement
				err = json.Unmarshal(value, &array)
				if err != nil {
					break
				}
				obj.valid.SetBit(&obj.valid, security_logging_object_network_policy_refs, 1)
				obj.network_policy_refs = make(contrail.ReferenceList, 0)
				for _, element := range array {
					ref := contrail.Reference{
						element.To,
						element.Uuid,
						element.Href,
						element.Attr,
					}
					obj.network_policy_refs = append(obj.network_policy_refs, ref)
				}
				break
			}
		case "security_group_refs":
			{
				type ReferenceElement struct {
					To   []string
					Uuid string
					Href string
					Attr SecurityLoggingObjectRuleListType
				}
				var array []ReferenceElement
				err = json.Unmarshal(value, &array)
				if err != nil {
					break
				}
				obj.valid.SetBit(&obj.valid, security_logging_object_security_group_refs, 1)
				obj.security_group_refs = make(contrail.ReferenceList, 0)
				for _, element := range array {
					ref := contrail.Reference{
						element.To,
						element.Uuid,
						element.Href,
						element.Attr,
					}
					obj.security_group_refs = append(obj.security_group_refs, ref)
				}
				break
			}
		case "firewall_policy_back_refs":
			{
				type ReferenceElement struct {
					To   []string
					Uuid string
					Href string
					Attr SloRateType
				}
				var array []ReferenceElement
				err = json.Unmarshal(value, &array)
				if err != nil {
					break
				}
				obj.valid.SetBit(&obj.valid, security_logging_object_firewall_policy_back_refs, 1)
				obj.firewall_policy_back_refs = make(contrail.ReferenceList, 0)
				for _, element := range array {
					ref := contrail.Reference{
						element.To,
						element.Uuid,
						element.Href,
						element.Attr,
					}
					obj.firewall_policy_back_refs = append(obj.firewall_policy_back_refs, ref)
				}
				break
			}
		case "firewall_rule_back_refs":
			{
				type ReferenceElement struct {
					To   []string
					Uuid string
					Href string
					Attr SloRateType
				}
				var array []ReferenceElement
				err = json.Unmarshal(value, &array)
				if err != nil {
					break
				}
				obj.valid.SetBit(&obj.valid, security_logging_object_firewall_rule_back_refs, 1)
				obj.firewall_rule_back_refs = make(contrail.ReferenceList, 0)
				for _, element := range array {
					ref := contrail.Reference{
						element.To,
						element.Uuid,
						element.Href,
						element.Attr,
					}
					obj.firewall_rule_back_refs = append(obj.firewall_rule_back_refs, ref)
				}
				break
			}
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (obj *SecurityLoggingObject) UpdateObject() ([]byte, error) {
	msg := map[string]*json.RawMessage{}
	err := obj.MarshalId(msg)
	if err != nil {
		return nil, err
	}

	if obj.modified.Bit(security_logging_object_security_logging_object_rules) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.security_logging_object_rules)
		if err != nil {
			return nil, err
		}
		msg["security_logging_object_rules"] = &value
	}

	if obj.modified.Bit(security_logging_object_security_logging_object_rate) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.security_logging_object_rate)
		if err != nil {
			return nil, err
		}
		msg["security_logging_object_rate"] = &value
	}

	if obj.modified.Bit(security_logging_object_id_perms) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.id_perms)
		if err != nil {
			return nil, err
		}
		msg["id_perms"] = &value
	}

	if obj.modified.Bit(security_logging_object_perms2) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.perms2)
		if err != nil {
			return nil, err
		}
		msg["perms2"] = &value
	}

	if obj.modified.Bit(security_logging_object_annotations) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.annotations)
		if err != nil {
			return nil, err
		}
		msg["annotations"] = &value
	}

	if obj.modified.Bit(security_logging_object_display_name) != 0 {
		var value json.RawMessage
		value, err := json.Marshal(&obj.display_name)
		if err != nil {
			return nil, err
		}
		msg["display_name"] = &value
	}

	if obj.modified.Bit(security_logging_object_network_policy_refs) != 0 {
		if len(obj.network_policy_refs) == 0 {
			var value json.RawMessage
			value, err := json.Marshal(
				make([]contrail.Reference, 0))
			if err != nil {
				return nil, err
			}
			msg["network_policy_refs"] = &value
		} else if !obj.hasReferenceBase("network-policy") {
			var value json.RawMessage
			value, err := json.Marshal(&obj.network_policy_refs)
			if err != nil {
				return nil, err
			}
			msg["network_policy_refs"] = &value
		}
	}

	if obj.modified.Bit(security_logging_object_security_group_refs) != 0 {
		if len(obj.security_group_refs) == 0 {
			var value json.RawMessage
			value, err := json.Marshal(
				make([]contrail.Reference, 0))
			if err != nil {
				return nil, err
			}
			msg["security_group_refs"] = &value
		} else if !obj.hasReferenceBase("security-group") {
			var value json.RawMessage
			value, err := json.Marshal(&obj.security_group_refs)
			if err != nil {
				return nil, err
			}
			msg["security_group_refs"] = &value
		}
	}

	if obj.modified.Bit(security_logging_object_tag_refs) != 0 {
		if len(obj.tag_refs) == 0 {
			var value json.RawMessage
			value, err := json.Marshal(
				make([]contrail.Reference, 0))
			if err != nil {
				return nil, err
			}
			msg["tag_refs"] = &value
		} else if !obj.hasReferenceBase("tag") {
			var value json.RawMessage
			value, err := json.Marshal(&obj.tag_refs)
			if err != nil {
				return nil, err
			}
			msg["tag_refs"] = &value
		}
	}

	return json.Marshal(msg)
}

func (obj *SecurityLoggingObject) UpdateReferences() error {

	if (obj.modified.Bit(security_logging_object_network_policy_refs) != 0) &&
		len(obj.network_policy_refs) > 0 &&
		obj.hasReferenceBase("network-policy") {
		err := obj.UpdateReference(
			obj, "network-policy",
			obj.network_policy_refs,
			obj.baseMap["network-policy"])
		if err != nil {
			return err
		}
	}

	if (obj.modified.Bit(security_logging_object_security_group_refs) != 0) &&
		len(obj.security_group_refs) > 0 &&
		obj.hasReferenceBase("security-group") {
		err := obj.UpdateReference(
			obj, "security-group",
			obj.security_group_refs,
			obj.baseMap["security-group"])
		if err != nil {
			return err
		}
	}

	if (obj.modified.Bit(security_logging_object_tag_refs) != 0) &&
		len(obj.tag_refs) > 0 &&
		obj.hasReferenceBase("tag") {
		err := obj.UpdateReference(
			obj, "tag",
			obj.tag_refs,
			obj.baseMap["tag"])
		if err != nil {
			return err
		}
	}

	return nil
}

func SecurityLoggingObjectByName(c contrail.ApiClient, fqn string) (*SecurityLoggingObject, error) {
	obj, err := c.FindByName("security-logging-object", fqn)
	if err != nil {
		return nil, err
	}
	return obj.(*SecurityLoggingObject), nil
}

func SecurityLoggingObjectByUuid(c contrail.ApiClient, uuid string) (*SecurityLoggingObject, error) {
	obj, err := c.FindByUuid("security-logging-object", uuid)
	if err != nil {
		return nil, err
	}
	return obj.(*SecurityLoggingObject), nil
}

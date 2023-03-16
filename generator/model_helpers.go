package main

import (
	"fmt"
	"strings"
)

func (m Model) AllowedFieldsOnCreateClause() string {
	allowedFields := []string{}
	for _, f := range m.Fields {
		switch f.OnCreate {
		case OnEnumAllow:
			fallthrough
		case OnEnumRequire:
			allowedFields = append(allowedFields, f.Name)
			break
		case OnEnumDeny:
		default:
			break
		}
	}
	return fmt.Sprintf("incomingData().keys().hasOnly([%s])", toQuotedString(allowedFields))
}
func (m Model) RequiredFieldsOnCreateClause() string {
	requiredFields := []string{}
	for _, f := range m.Fields {
		switch f.OnCreate {
		case OnEnumRequire:
			requiredFields = append(requiredFields, f.Name)
			break
		case OnEnumAllow:
			fallthrough
		case OnEnumDeny:
			fallthrough
		default:
			break
		}
	}
	return fmt.Sprintf("incomingData().keys().hasAll([%s])", toQuotedString(requiredFields))
}
func (m Model) AllowedFieldsOnUpdateClause() string {
	allowedFields := []string{}
	for _, f := range m.Fields {
		switch f.OnUpdate {
		case OnEnumAllow:
			fallthrough
		case OnEnumRequire:
			allowedFields = append(allowedFields, f.Name)
			break
		case OnEnumDeny:
			fallthrough
		default:
			break
		}
	}
	return fmt.Sprintf("delta().hasOnly([%s])", toQuotedString(allowedFields))
}
func (m Model) RequiredFieldsOnUpdateClause() string {
	requiredFields := []string{}
	for _, f := range m.Fields {
		switch f.OnUpdate {
		case OnEnumRequire:
			requiredFields = append(requiredFields, f.Name)
			break
		case OnEnumAllow:
			fallthrough
		case OnEnumDeny:
		default:
			break
		}
	}
	return fmt.Sprintf("delta().hasAll([%s])", toQuotedString(requiredFields))
}
func (m Model) TypeCheckFieldsClauses() string {
	fieldClauses := []string{}

	for i, f := range m.Fields {
		clause := ""
		switch f.Datatype {
		case DataTypeEnumArray:
			clause = fmt.Sprintf("typeCheckList(\"%s\", %s, %s, %s)", f.Name, toBoolString(f.Nullable), toIntOrNull(f.Min), toIntOrNull(f.Max))
			break
		case DataTypeEnumBoolean:
			clause = fmt.Sprintf("typeCheckBoolean(\"%s\", %s)", f.Name, toBoolString(f.Nullable))
			break
		case DataTypeEnumBytes:
			clause = fmt.Sprintf("typeCheckBytes(\"%s\", %s)", f.Name, toBoolString(f.Nullable))
			break
		case DataTypeEnumDateTime:
			clause = fmt.Sprintf("typeCheckTimestamp(\"%s\", %s)", f.Name, toBoolString(f.Nullable))
			break
		case DataTypeEnumFloat:
			clause = fmt.Sprintf("typeCheckFloat(\"%s\", %s, %s, %s)", f.Name, toBoolString(f.Nullable), toIntOrNull(f.Min), toIntOrNull(f.Max))
			break
		case DataTypeEnumInteger:
			clause = fmt.Sprintf("typeCheckInt(\"%s\", %s, %s, %s)", f.Name, toBoolString(f.Nullable), toIntOrNull(f.Min), toIntOrNull(f.Max))
			break
		case DataTypeEnumLatLng:
			// clause := fmt.Sprintf("typeCheckBoolean(\"%s\", %s)", f.Name, toBoolString(f.Nullable))
			// fieldClauses = append(fieldClauses, clause)
			break
		case DataTypeEnumMap:
			// clause := fmt.Sprintf("typeCheckBoolean(\"%s\", %s)", f.Name, toBoolString(f.Nullable))
			// fieldClauses = append(fieldClauses, clause)
			break
		case DataTypeEnumReference:
			// clause := fmt.Sprintf("typeCheckBoolean(\"%s\", %s)", f.Name, toBoolString(f.Nullable))
			// fieldClauses = append(fieldClauses, clause)
			break
		case DataTypeEnumText:
			clause = fmt.Sprintf("typeCheckString(\"%s\", %s, %s)", f.Name, toBoolString(f.Nullable), toEscapedString(f.Match))
			break
		}

		if clause == "" {
			continue
		}
		if i == 0 {
			fieldClauses = append(fieldClauses, clause)
		} else {
			fieldClauses = append(fieldClauses, fmt.Sprintf("        %s", clause))
		}
	}

	return strings.Join(fieldClauses, " && \n")
}

func toEscapedString(s *string) string {
	if s == nil {
		return "null"
	} else {
		//Convert abc => "abc", ab"c => "ab\"c", ab\c => "ab\\c"
		return "\"" + strings.ReplaceAll(strings.ReplaceAll(*s, "\\", "\\\\"), "\"", "\\\"") + "\""
	}
}

func toIntOrNull(i *int) string {
	if i == nil {
		return "null"
	} else {
		return fmt.Sprint(*i)
	}
}

func toQuotedString(arr []string) string {
	for i := 0; i < len(arr); i++ {
		arr[i] = fmt.Sprintf("\"%s\"", arr[i])
	}
	return strings.Join(arr, ",")
}

func toBoolString(b *bool) string {
	if b != nil && *b == true {
		return "true"
	} else {
		return "false"
	}
}

package main

type (
	DocumentRoot struct {
		Models []*Model `json:"models"`
	}

	Model struct {
		Match  string   `json:"match"`
		Fields []*Field `json:"fields"`
	}

	Field struct {
		Name     string       `json:"name"`
		Datatype DataTypeEnum `json:"datatype"`
		Min      *int         `json:"min,omitempty"`
		Max      *int         `json:"max,omitempty"`
		Nullable *bool        `json:"nullable,omitempty"`
		Match    *string      `json:"match"`
		OnCreate OnEnum       `json:"onCreate,omitempty"`
		OnUpdate OnEnum       `json:"onUpdate,omitempty"`
		Fields   []*Field     `json:"fields,omitempty"`
	}
)

type DataTypeEnum string

const (
	DataTypeEnumArray     = "array"
	DataTypeEnumBoolean   = "boolean"
	DataTypeEnumBytes     = "bytes"
	DataTypeEnumDateTime  = "datetime"
	DataTypeEnumFloat     = "float"
	DataTypeEnumLatLng    = "latlng"
	DataTypeEnumInteger   = "integer"
	DataTypeEnumMap       = "map"
	DataTypeEnumReference = "reference"
	DataTypeEnumText      = "text"
)

type OnEnum string

const (
	OnEnumRequire = "require"
	OnEnumDeny    = "deny"
	OnEnumAllow   = "allow"
)

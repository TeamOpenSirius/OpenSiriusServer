// Code generated from dump.cs. DO NOT EDIT.
// Source namespace: SiriusApi.Shared.Annotations
package dto

type ForeignKeyAttribute struct {
	MapsToPropertyName string `json:"mapsToPropertyName" msgpack:"mapsToPropertyName"`
}

type InversePropertyAttribute struct {
	MapsFromPropertyName string `json:"mapsFromPropertyName" msgpack:"mapsFromPropertyName"`
}

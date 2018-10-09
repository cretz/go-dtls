package model

type ExtensionType uint16
type Extension interface {
	Type() ExtensionType
}

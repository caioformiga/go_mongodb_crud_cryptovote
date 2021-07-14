package model

type Cryptocurrency struct {
	Name   string `bson:"name,omitempty"`
	Symbol string `bson:"symbol,omitempty"`
}

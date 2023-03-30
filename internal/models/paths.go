package models

type Path struct {
	ActiveLink  string `bson:"active_link" json:"active_link"`
	HistoryLink string `bson:"history_link" json:"history_link"`
}

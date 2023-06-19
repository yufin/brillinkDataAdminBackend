package utils

import uuid2 "github.com/google/uuid"

func NewRandomUUID() string {
	uuid, _ := uuid2.NewRandom()
	return uuid.String()
}

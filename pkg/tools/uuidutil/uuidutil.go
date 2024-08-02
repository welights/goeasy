package uuidutil

import (
	"strings"

	"github.com/google/uuid"
)

func Generate() string {
	return uuid.New().String()
}

func GenerateWithoutHyphen() string {
	return strings.ReplaceAll(Generate(), "-", "")
}

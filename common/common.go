package common

import (
	"strings"

	"github.com/google/uuid"
)

func NewUuid() string {
	uuid := uuid.New()
	return strings.Replace(uuid.String(), "-", "", -1)
}

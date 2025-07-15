package models

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
)

const ErrGenerateUniqueID = "failed to generate unique ID"

type ICodeGenerator interface {
	GenerateUniqueID() (string, error)
}

type CodeGenerator struct{}

func NewCodeGenerator() *CodeGenerator {
	return &CodeGenerator{}
}

func (cg *CodeGenerator) GenerateUniqueID() (string, error) {
	randomBytes := make([]byte, 6)
	_, err := rand.Read(randomBytes)

	if err != nil {
		return "", errors.New(ErrGenerateUniqueID)
	}

	return base64.URLEncoding.EncodeToString(randomBytes)[:6], nil
}

package utils

import "errors"

func NewResourceNotFoundError(resourceName string) error {
	return errors.New(resourceName + " not found")
}

package services

import (
	"fmt"
)

type Service struct {
	Error error
}

func (service *Service) AddError(err error) error {
	if service.Error == nil {
		service.Error = err
	} else if err != nil {
		service.Error = fmt.Errorf("%v; %w", service.Error, err)
	}
	return service.Error
}

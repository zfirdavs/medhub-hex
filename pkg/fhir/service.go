package fhir

import (
	"context"

	schema "github.com/MedHubUz/fhirschema"
)

type service struct {
	repo Repository
}

// NewService create new service
func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) Read(ctx context.Context, resource *Resource) (interface{}, error) {
	return s.repo.Read(ctx, resource)
}

func (s *service) ReadAll(ctx context.Context, resourceType string) (*schema.Bundle, error) {
	return s.repo.ReadAll(ctx, resourceType)
}

func (s *service) Create(ctx context.Context, resource *Resource) (*Resource, error) {
	// resourseJSON, err := json.Marshal(fhirResource)
	// if err != nil {
	// 	return nil, err
	// }

	// resource := Resource{
	// 	ResourceType:   resourceType,
	// 	ResourceID:     "1",
	// 	Data:           string(resourseJSON),
	// 	CreatedAt:      time.Now().UTC().Format("2006-01-02 15:04:05"),
	// 	PractitionerID: "1",
	// }

	return s.repo.Create(ctx, resource)
}

func (s *service) Update(ctx context.Context, resourceType, resourceID string) (interface{}, error) {
	return s.repo.Update(ctx, resourceType, resourceID)
}

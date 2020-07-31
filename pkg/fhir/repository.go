package fhir

import (
	"context"

	schema "github.com/MedHubUz/fhirschema"
)

// Repository provides fhir standard spec methods
type Repository interface {
	Read(ctx context.Context, resource *Resource) (interface{}, error)
	ReadAll(ctx context.Context, resourceType string) (*schema.Bundle, error)
	Create(ctx context.Context, resource *Resource) (*Resource, error)
	Update(ctx context.Context, resourceType, resourceID string) (interface{}, error)
}

// Service provides fhir standard spec methods
type Service interface {
	Read(ctx context.Context, resource *Resource) (interface{}, error)
	ReadAll(ctx context.Context, resourceType string) (*schema.Bundle, error)
	Create(ctx context.Context, resource *Resource) (*Resource, error)
	Update(ctx context.Context, resourceType, resourceID string) (interface{}, error)
}

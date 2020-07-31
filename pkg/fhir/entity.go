package fhir

// Resource provides fhir standard implementation common resource
type Resource struct {
	ID             uint64 `json:"id,omitempty"`
	ResourceID     string `json:"resource_id,omitempty" db:"resource_id"`
	PractitionerID string `json:"practitioner_id,omitempty" db:"practitioner_id"`
	ResourceType   string `json:"resource_type,omitempty" db:"resource_type"`
	Clinic         string `json:"clinic,omitempty"`
	CreatedAt      string `json:"created_at,omitempty" db:"created_at"`
	Data           string `json:"data,omitempty"`
}

// type resourceJSON struct {
// 	Data interface{} `json:"data,omitempty"`
// }

// func (r resourceJSON) Value() (driver.Value, error) {
// 	return json.Marshal(r.Data)
// }

// func (d resourceJSON) Scan(value interface{}) error {
// 	switch v := value.(type) {
// 	case []byte:
// 		return json.Unmarshal(v, &d)
// 	case string:
// 		return json.Unmarshal([]byte(v), &d)
// 	case nil:
// 		return nil
// 	}
// 	return errors.New("type assertion to []byte, string and nil failed")
// }

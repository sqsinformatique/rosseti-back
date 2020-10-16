package models

type TemplateItem struct {
	Type        string `json:"item_type,omitempty" bson:"item_type"`
	Name        string `json:"item_name,omitempty" bson:"item_name"`
	Description string `json:"item_description,omitempty" bson:"item_description"`
	Limit       struct {
		Max string `json:"item_max,omitempty" bson:"item_max"`
		Min string `json:"item_min,omitempty" bson:"item_min"`
	} `json:"item_limit" bson:"item_limit"`
	Requried bool     `json:"item_requried,omitempty" bson:"item_requried"`
	Values   []string `json:"item_values,omitempty" bson:"item_values"`
}

type Template struct {
	Name string                   `json:"template_name" bson:"template_name"`
	Body map[string]*TemplateItem `json:"template_body" bson:"template_body"`
	Timestamp
}

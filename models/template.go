package models

type TemplateItem struct {
	Type        string `json:"item_type" bson:"item_type"`
	Description string `json:"item_description" bson:"item_description"`
	Limit       struct {
		Max string `json:"item_max" bson:"item_max"`
		Min string `json:"item_min" bson:"item_min"`
	} `json:"item_limit" bson:"item_limit"`
	Requried bool     `json:"item_requried" bson:"item_requried"`
	Values   []string `json:"item_values" bson:"item_values"`
}

type Template struct {
	Name string                   `json:"template_name" bson:"template_name"`
	Body map[string]*TemplateItem `json:"template_body" bson:"template_body"`
	Timestamp
}

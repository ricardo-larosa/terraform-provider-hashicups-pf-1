package hashicups

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Order -
type Order struct {
	ID          types.String `tfsdk:"id"`
	Items       []OrderItem  `tfsdk:"items"`
	LastUpdated types.String `tfsdk:"last_updated"`
}

// OrderItem -
type OrderItem struct {
	Coffee   Coffee `tfsdk:"coffee"`
	Quantity int    `tfsdk:"quantity"`
}

// Coffee -
type Coffee struct {
	ID          int          `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Teaser      types.String `tfsdk:"teaser"`
	Description types.String `tfsdk:"description"`
	Price       types.Number `tfsdk:"price"`
	Image       types.String `tfsdk:"image"`
	Ingredients []Ingredient `tfsdk:"ingredients"`
}

// Ingredient -
type Ingredient struct {
	ID int `tfsdk:"ingredient_id"`
}

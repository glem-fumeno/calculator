package schemas

type RecipeItemType string

const (
	ItemTypeIngredient RecipeItemType = "ingredient"
	ItemTypeProduct    RecipeItemType = "product"
)

type DBItem struct {
	ItemName string
	Unit     string
}

type DBRecipe struct {
	RecipeName string
}

type DBRecipeItem struct {
	RecipeName string
	ItemName   string
	ItemType   RecipeItemType
	Quantity   int
}

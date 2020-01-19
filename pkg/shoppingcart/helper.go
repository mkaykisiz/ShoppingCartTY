package shoppingcart

func HasElement(elementName string, elements []string) bool {
	for _, element := range elements {
		if element == elementName {
			return true
		}
	}
	return false
}

func GetSelfAndParentsCategories(category Category, categories []Category) []Category {
	// get selected category and parents for search campaign
	if category.ParentCategory == nil {
		return append([]Category{category}, categories...)
	}
	return append([]Category{category}, GetSelfAndParentsCategories(*category.ParentCategory, categories)...)
}

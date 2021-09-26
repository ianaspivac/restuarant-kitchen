package components
type Food struct{
	id int
	name string
	preparationTime float32
	complexity int
	cookingApparatus string
}
var Menu = []Food {
	{
		id: 1,
		name: "pizza",
		preparationTime: 20 ,
		complexity: 2 ,
		cookingApparatus: "oven",
	},
	{
		id: 2,
		name: "salad",
		preparationTime: 10 ,
		complexity: 1 ,
		cookingApparatus: "",
	},
	{
		id: 3,
		name: "zeama",
		preparationTime: 7 ,
		complexity: 1 ,
		cookingApparatus: "stove",
	},
	{
		id: 4,
		name: "Scallop Sashimi with Meyer Lemon Confit",
		preparationTime: 32 ,
		complexity: 3 ,
		cookingApparatus: "",
	},
	{
		id: 5,
		name: "Island Duck with Mulberry Mustard",
		preparationTime: 35 ,
		complexity: 3 ,
		cookingApparatus: "oven",
	},
	{
		id: 6,
		name: "Waffles",
		preparationTime: 10 ,
		complexity: 1 ,
		cookingApparatus: "stove",
	},
	{
		id: 7,
		name: "Aubergine",
		preparationTime: 20 ,
		complexity: 2 ,
		cookingApparatus: "",
	},
	{
		id: 8,
		name: "Lasagna",
		preparationTime: 30 ,
		complexity: 2 ,
		cookingApparatus: "oven",
	},
	{
		id: 9,
		name: "Burger",
		preparationTime: 15 ,
		complexity: 1 ,
		cookingApparatus: "oven",
	},
	{
		id: 10,
		name: "Gyros",
		preparationTime: 15 ,
		complexity: 1 ,
		cookingApparatus: "",
	},
}

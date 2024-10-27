package data


type Products struct {
  Title string
  CurrentPrice float64
  PreviousPrice float64
  Image string
}

func GetAll() []Products{
	return list
}

var list = []Products{
	{
		Title: "Hand Base Lamp",
		CurrentPrice: 35.00,
	    PreviousPrice: 55.00,
	    Image: "hand-base-lamp.png",
	},
	{
		Title: "Vintage Chair",
		CurrentPrice: 65.00,
	    PreviousPrice: 95.00,
	    Image: "vintage-chair.png",
	},
 	{
		Title: "Lamp tool",
		CurrentPrice: 35.00,
	    PreviousPrice: 45.00,
	    Image: "lamp-tool.png",
	},
    {
		Title: "Stylish Chair",
		CurrentPrice: 45.00,
	    PreviousPrice: 55.00,
	    Image: "stylish-chair.png",
	},
	{
		Title: "Vintage Chair",
		CurrentPrice: 65.00,
	    PreviousPrice: 95.00,
	    Image: "vintage-chair2.png",
	},
	{
		Title: "Stackable Chair ",
		CurrentPrice: 87.00,
	    PreviousPrice: 97.00,
	    Image: "stackable-chair.png",
	},
}
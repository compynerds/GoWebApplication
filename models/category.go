package models

import ()

type Category struct {
	imageUrl      string
	title         string
	description   string
	isOrientRight bool
	id            int
}

func (c *Category) ImageUrl() string {
	return c.imageUrl
}

func (c *Category) Title() string {
	return c.title
}

func (c *Category) Description() string {
	return c.description
}

func (c *Category) IsOrientRight() bool {
	return c.isOrientRight
}

func (c *Category) Id() int {
	return c.id
}

func (c *Category) SetImageUrl(value string) {
	c.imageUrl = value
}

func (c *Category) SetTitle(value string) {
	c.title = value
}

func (c *Category) SetDescription(value string) {
	c.description = value
}

func (c *Category) SetIsOrientRight(value bool) {
	c.isOrientRight = value
}

func (c *Category) SetId(value int) {
	c.id = value
}

func GetCategories() []Category {
	result := []Category{
		Category{
			id:       1,
			imageUrl: "lemon.png",
			title:    "Juices and Mixes",
			description: `Explore our wide assortment of juices and mixes expected by
                today's lemonade stand clientelle. Now featuring a full line of
                organic juices that are guaranteed to be obtained from trees that
                have never been treated with pesticides or artificial
                fertilizers.`,
			isOrientRight: false,
		},
		Category{
			id:       2,
			imageUrl: "kiwi.png",
			title:    "Cups, Straws, and Other Supplies",
			description: `From paper cups to bio-degradable plastic to straws and
              napkins, LSS is your source for the sundries that keep your stand
              running smoothly.`,
			isOrientRight: true,
		},
		Category{
			id:       3,
			imageUrl: "pineapple.png",
			title:    "Signs and Advertising",
			description: `Sure, you could just wait for people to find your stand
              along the side of the road, but if you want to take it to the next
              level, our premium line of advertising supplies.`,
			isOrientRight: false,
		},
	}
  return result
}

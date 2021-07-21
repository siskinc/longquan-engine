package property

type PropertyMap struct {
	Property
	Children []*PropertyMap
}

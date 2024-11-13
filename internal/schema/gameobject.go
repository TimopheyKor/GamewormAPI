package schema

// GameObject holds the columns that a row for one game in the GamewormDB
// will have. It should be used to interface between receiving API calls to
// GamewormAPI with game data and pushing or editing game data in Google Sheets.
type GameObject struct {
	Id        string
	Title     string
	Image     string
	Developer string
	Publisher string
}

func (g *GameObject) NewGameObject(title, img, dev, pub string) *GameObject {
	adjTitle := TrimExtendStr(title)
	adjDev := TrimExtendStr(dev)
	idx := len(adjTitle) - 3
	genId := adjTitle[:3] + adjTitle[idx:] + 
	return &GameObject{
		Id:    title[],
		Title: title,
	}
}

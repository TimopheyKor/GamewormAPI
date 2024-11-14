package schema

// GameObject holds the columns that a row for one game in the GamewormDB
// will have. It should be used to interface between receiving API calls to
// GamewormAPI with game data and pushing or editing game data in Google Sheets.
type GameObject struct {
	id        string
	Title     string
	Image     string
	Developer string
	Publisher string
}

// NewGameObject requires a title and dev string, then takes additional option
// functions, and returns a pointer to a new game object constructed with
// the given parameters.
func (g *GameObject) NewGameObject(title, dev string, opts ...GameOption) *GameObject {
	// Create an ID from the required parameters:
	adjTitle, adjDev := TrimExtendStr(title), TrimExtendStr(dev)
	idx := len(adjTitle) - 3
	genId := adjTitle[:3] + adjTitle[idx:] + 
	return &GameObject{
		Id:    title[],
		Title: title,
	}
}

// GetID returns the id of the GameObject.
func (g *GameObject) GetID() string {
	return g.id
}

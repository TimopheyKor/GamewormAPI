package schema

// GameObject holds the columns that a row for one game in the GamewormDB
// will have. It should be used to interface between receiving API calls to
// GamewormAPI with game data and pushing or editing game data in Google Sheets.
type GameObject struct {
	id        string
	title     string
	Image     string
	developer string
	Publisher string
}

// NewGameObject requires a title and dev string, then takes additional option
// functions, and returns a pointer to a new game object constructed with
// the given parameters.
func NewGameObject(title, dev string, opts ...GameOption) *GameObject {
	newGameObj := &GameObject{
		id:        GenerateGameID(title, dev),
		title:     title,
		developer: dev,
	}
	for _, opt := range opts {
		opt(newGameObj)
	}
	return newGameObj
}

// GetID returns the id of the GameObject.
func (g *GameObject) GetID() string {
	return g.id
}

func (g *GameObject) ToSlice() []any {
	return []any{g.id, g.title, g.Image, g.developer, g.Publisher}
}

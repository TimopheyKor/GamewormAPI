package schema

import "github.com/TimopheyKor/GamewormAPI/internal/static"

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

// NewGameObjectFromDB takes a row from the Games table in the GamewormDB sheet
// and converts it to a new GameObject.
func NewGameObjectFromDB(row [5]string) *GameObject {
	newGameObj := &GameObject{
		id:        row[static.GamePK],
		title:     row[static.GTitle],
		Image:     row[static.GImage],
		developer: row[static.GDev],
		Publisher: row[static.GPub],
	}
	return newGameObj
}

// GetID returns the id of the GameObject.
func (g *GameObject) GetID() string {
	return g.id
}

// ToSlice is a method on GameObject that returns a slice of any type containing
// the primary fields of the GameObject. This is primarily for the use of
// passing game data to the Google Sheets API in a recognizable format that
// adheres to the GamewormDB Schema.
func (g *GameObject) ToSlice() []any {
	return []any{g.id, g.title, g.Image, g.developer, g.Publisher}
}

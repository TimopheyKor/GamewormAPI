package schema

// GameOption "Options" type is a function that takes a pointer to a GameObject.
type GameOption func(*GameObject)

// WithImg(url) is used to pass a URL into the paramaters of NewGameObject to
// populate the GameObject's Image field.
func WithImg(url string) GameOption {
	return func(g *GameObject) {
		g.Image = url
	}
}

// WithPub(publisher) is used to pass the game's publisher as a string into the
// parameter of NewGameObject to populate the GameObject's Publisher field.
func WithPub(pub string) GameOption {
	return func(g *GameObject) {
		g.Publisher = pub
	}
}

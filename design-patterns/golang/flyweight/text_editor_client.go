package flyweight

// TextEditorClient represents a client that uses the flyweight objects
type TextEditorClient struct {
	factory *CharacterFlyweightFactory
}

// NewTextEditorClient creates a new TextEditorClient instance
func NewTextEditorClient() *TextEditorClient {
	return &TextEditorClient{
		factory: NewCharacterFlyweightFactory(),
	}
}

// RenderText renders a text string using flyweight objects
func (c *TextEditorClient) RenderText(text string, x, y int, fontFamily string, fontSize int, color string) {
	column := 0
	for _, char := range text {
		flyweight := c.factory.GetFlyweight(char, fontFamily, fontSize, color)
		flyweight.Draw(x+column*fontSize, y)
		column++
	}
}

// GetUniqueCharacterCount returns the number of unique character flyweights
func (c *TextEditorClient) GetUniqueCharacterCount() int {
	return c.factory.GetFlyweightCount()
}

package model


// A TopCourse is a Course with no name - it's the root of the tree (benelux).
type TopCourse struct {
	path     FilePath
	children []Tutorial
}

func NewTopCourse(p FilePath, c []Tutorial) *TopCourse { return &TopCourse{p, c} }
func (t *TopCourse) Accept(v TutVisitor)               { v.VisitTopCourse(t) }
func (t *TopCourse) Name() string                      { return "" }
func (t *TopCourse) Path() FilePath                    { return t.path }
func (t *TopCourse) Children() []Tutorial              { return t.children }

// A Course, or directory, has a name but no content, and an ordered list of
// Lessons and Courses. If the list is empty, the Course is dropped (hah!).
type Course struct {
	patrh    FilePath
	children []Tutorial
}

func NewCourse(p FilePath, c []Tutorial) *Course { return &Course{p, c} }
func (c *Course) Accept(v TutVisitor)            { v.VisitCourse(c) }
func (c *Course) Name() string                   { return c.patrh.Base() }
func (c *Course) Path() FilePath                 { return c.patrh }
func (c *Course) Children() []Tutorial           { return c.children }

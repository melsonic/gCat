package util

/// context object with two fields
/// field number => whether '-n' flag is set
/// number_blankline => whether '-b' flag is set
type Context struct {
	number           bool
	number_blankline bool
}

/// get whether '-n' flag is passed or not
func (c *Context) GetNumberFlagState() bool {
  return c.number
}

/// set whether '-n' flag is passed or not
func (c *Context) SetNumberFlagState(numberFlag bool) {
	c.number = numberFlag
}

/// get whether '-b' flag is passed or not 
func (c *Context) GetNumberBlanklineFlagState() bool {
  return c.number_blankline
}

/// set whether '-b' flag is passed or not 
func (c *Context) SetNumberBlanklineFlagState(numberBlankLineFlag bool) {
	c.number_blankline = numberBlankLineFlag
}

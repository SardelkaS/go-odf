package style

import (
	"fmt"
	"github.com/SardelkaS/go-odf/odt/content/style/props"
	"github.com/SardelkaS/go-odf/odt/content/style/types"
	"strconv"
	"sync/atomic"
)

var nameIter = atomic.Uint64{}

const _defaultParentStyleName = "Обычный"

type Style struct {
	name            string
	parentStyleName string
	family          types.Family
	Props           props.Props
}

// New creates new Style with default values
func New() Style {
	iter := nameIter.Load()
	nameIter.Add(1)
	if iter == 0 {
		iter = 1
		nameIter.Add(1)
	}

	return Style{
		name:            fmt.Sprintf("P%s", strconv.FormatUint(iter, 10)),
		parentStyleName: _defaultParentStyleName,
		family:          types.Family_Paragraph,
		Props:           props.New(),
	}
}

// GetName returns style name
func (s Style) GetName() string {
	return s.name
}

// Generate generates xml code
func (s Style) Generate() string {
	propsXml := s.Props.Generate()
	return fmt.Sprintf(`<style style:name="%s" style:parent-style-name="%s" style:family="%s">%s</style>`,
		s.name, s.parentStyleName, s.family, propsXml)
}

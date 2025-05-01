package odt

import (
	"fmt"
	"strconv"
	"sync/atomic"
)

var hsNameIter = atomic.Uint64{}

type headingStyle struct {
	name  string
	level int64
}

func newHeadingStyle(level int64) *headingStyle {
	iter := hsNameIter.Load()
	hsNameIter.Add(1)
	if iter == 0 {
		iter = 2
		hsNameIter.Add(2)
	}

	return &headingStyle{
		name:  fmt.Sprintf("Contents_%s", strconv.FormatUint(iter, 10)),
		level: level,
	}
}

func (h *headingStyle) generate() string {
	return fmt.Sprintf(`<style:style style:name="%s"
            style:family="paragraph" style:parent-style-name="Index" style:class="index">
            <style:paragraph-properties fo:margin-left="%.2fcm" fo:text-indent="0cm"
                style:auto-text-indent="false">
                <style:tab-stops>
                    <style:tab-stop style:position="%.2fcm" style:type="right"
                        style:leader-style="dotted" style:leader-text="." />
                </style:tab-stops>
            </style:paragraph-properties>
        </style:style>`,
		h.name, float64(h.level-1)*0.5, float64(17)-float64(h.level-1)*0.5)
}

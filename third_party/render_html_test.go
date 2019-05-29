package devrel_tutorial

import (
	"testing"
	"github.com/googlecodelabs/tools/claat/util"
)

func TestRenderHtmlHeading(t *testing.T) {
	h2 := &Heading{
		Level: 2,
		Text:  "<script>some _very_ bad code!</script>",
	}
	h3 := &Heading{
		Level: 3,
		Text:  "D@ ?òü ǝ$çâpæ urlquery? '>__<' {&]",
	}
	h4 := &Heading{
		Level: 4,
		Text:  "**__Markdown not ![esca](ped)__**",
	}

	tests := []*util.TestingBatch{
		{h2.Html(), "<h2>&lt;script&gt;some _very_ bad code!&lt;/script&gt;</h2>"},
		{h3.Html(), "<h3>D@ ?òü ǝ$çâpæ urlquery? &#39;&gt;__&lt;&#39; {&amp;]</h3>"},
		{h4.Html(), "<h4>**__Markdown not ![esca](ped)__**</h4>"},
	}
	util.TestBatch(tests, t)
}

package echopongo2

import (
	"fmt"
	"io"
	"io/fs"
	"myapp/core/libs/helpers"
	"strings"

	"github.com/flosch/pongo2/v5"
	"github.com/labstack/echo/v4"
)

type Renderer struct {
	tplSet *pongo2.TemplateSet
}

// NewRendererOverFs 基于fs接口的renderer
// viewsFs参数说明：
//      外部首先使用embed包，将模板目录封装到embed.FS，例如
//           //go:embed views
//           var embedFs embed.FS
//           renderer := NewRenderer(fs.Sub(embedFs, "views"))
func NewRendererOverFs(name string, viewsFs fs.FS) *Renderer {
	tplSet := pongo2.NewSet(name, pongo2.NewFSLoader(viewsFs))
	return &Renderer{
		tplSet: tplSet,
	}
}

// NewRendererOverDir 基于本地文件目录的renderer
func NewRendererOverDir(name string, baseDir string) *Renderer {
	tplSet := pongo2.NewSet(name, helpers.Must(pongo2.NewLocalFileSystemLoader(baseDir)))
	return &Renderer{
		tplSet: tplSet,
	}
}

func (renderer *Renderer) Render(w io.Writer, tplName string, data interface{}, c echo.Context) error {
	if !strings.HasSuffix(tplName, ".html") {
		tplName = tplName + ".html"
	}
	tpl, err := renderer.tplSet.FromCache(tplName)
	if err != nil {
		return err
	}
	var pongoContext pongo2.Context
	if data != nil {
		switch v := data.(type) {
		case echo.Map:
			pongoContext = pongo2.Context(v)
		case pongo2.Context:
			pongoContext = v
		default:
			return fmt.Errorf("can not convert %T to pongo2.Context", data)
		}
	}
	return tpl.ExecuteWriter(pongoContext, w)
}

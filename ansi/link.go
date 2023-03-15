package ansi

import (
	"github.com/muesli/termenv"
	"io"
	"net/url"
)

// A LinkElement is used to render hyperlinks.
type LinkElement struct {
	Text    string
	BaseURL string
	URL     string
}

func (link_element *LinkElement) Render(w io.Writer, ctx RenderContext) error {
	if ctx.options.HyperLinks {
		return (&BaseElement{
			Token: termenv.Hyperlink(link_element.URL, link_element.Text),
			Style: ctx.options.Styles.Link,
		}).Render(w, ctx)
	} else {
		var rendered_text = false
		if len(link_element.Text) > 0 && link_element.Text != link_element.URL {
			err := (&BaseElement{
				Token: link_element.Text,
				Style: ctx.options.Styles.LinkText,
			}).Render(w, ctx)
			if err != nil {
				return err
			} else {
				rendered_text = true
			}
		}

		url, err := url.Parse(link_element.URL)
		// if the URL only consists of an anchor, ignore it
		if err != nil && "#"+url.Fragment != link_element.URL {
			padding := ""
			if rendered_text {
				padding = " "
			}
			return (&BaseElement{
				Token:  resolveRelativeURL(link_element.BaseURL, link_element.URL),
				Prefix: padding,
				Style:  ctx.options.Styles.Link,
			}).Render(w, ctx)
		} else {
			return nil
		}
	}
}

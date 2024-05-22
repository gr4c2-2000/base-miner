package crawl

import (
	"context"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/chromedp"
)

type Command struct {
	Website      string
	ResponseChan chan *goquery.Document
}

type Headless struct {
	internalChan chan Command
	globalCtx    context.Context
	globalCancel context.CancelFunc
}

func InitHeadless(ctx context.Context, chanbuff int) *Headless {
	h := Headless{}
	h.globalCtx, h.globalCancel = chromedp.NewContext(context.Background())
	h.internalChan = make(chan Command, chanbuff)
	return &h
}

func (h *Headless) worker() {

	for command := range h.internalChan {
		res, _ := h.ParseWebApp(command.Website)
		command.ResponseChan <- res
	}

}

func (h *Headless) PushCommand(c Command) {
	h.internalChan <- c
}

func (h *Headless) Cancel() {
	h.globalCancel()
}

func (h *Headless) ParseWebApp(url string) (*goquery.Document, error) {
	var outterHTML string
	ctx, cancel := chromedp.NewContext(h.globalCtx)
	defer cancel()

	if err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(url),

		chromedp.WaitReady(":root"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			node, err := dom.GetDocument().Do(ctx)
			if err != nil {
				return err
			}

			outterHTML, err = dom.GetOuterHTML().WithNodeID(node.NodeID).Do(ctx)
			return err
		}),
	}); err != nil {
		return nil, fmt.Errorf("ParseWebApp(): ActionFunc(): %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(outterHTML))
	if err != nil {
		return nil, fmt.Errorf("ParseWebApp(): goquery.NewDocumentFromReader(): %w", err)
	}

	return doc, nil
}

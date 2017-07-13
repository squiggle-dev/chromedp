// Package page provides the Chrome Debugging Protocol
// commands, types, and events for the Page domain.
//
// Actions and events related to the inspected page belong to the page
// domain.
//
// Generated by the chromedp-gen command.
package page

// Code generated by chromedp-gen. DO NOT EDIT.

import (
	"context"
	"encoding/base64"

	cdp "github.com/knq/chromedp/cdp"
	"github.com/knq/chromedp/cdp/debugger"
	"github.com/knq/chromedp/cdp/dom"
	"github.com/knq/chromedp/cdp/runtime"
)

// EnableParams enables page domain notifications.
type EnableParams struct{}

// Enable enables page domain notifications.
func Enable() *EnableParams {
	return &EnableParams{}
}

// Do executes Page.enable against the provided context and
// target handler.
func (p *EnableParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	return h.Execute(ctxt, cdp.CommandPageEnable, nil, nil)
}

// DisableParams disables page domain notifications.
type DisableParams struct{}

// Disable disables page domain notifications.
func Disable() *DisableParams {
	return &DisableParams{}
}

// Do executes Page.disable against the provided context and
// target handler.
func (p *DisableParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	return h.Execute(ctxt, cdp.CommandPageDisable, nil, nil)
}

// AddScriptToEvaluateOnNewDocumentParams evaluates given script in every
// frame upon creation (before loading frame's scripts).
type AddScriptToEvaluateOnNewDocumentParams struct {
	Source string `json:"source"`
}

// AddScriptToEvaluateOnNewDocument evaluates given script in every frame
// upon creation (before loading frame's scripts).
//
// parameters:
//   source
func AddScriptToEvaluateOnNewDocument(source string) *AddScriptToEvaluateOnNewDocumentParams {
	return &AddScriptToEvaluateOnNewDocumentParams{
		Source: source,
	}
}

// AddScriptToEvaluateOnNewDocumentReturns return values.
type AddScriptToEvaluateOnNewDocumentReturns struct {
	Identifier ScriptIdentifier `json:"identifier,omitempty"` // Identifier of the added script.
}

// Do executes Page.addScriptToEvaluateOnNewDocument against the provided context and
// target handler.
//
// returns:
//   identifier - Identifier of the added script.
func (p *AddScriptToEvaluateOnNewDocumentParams) Do(ctxt context.Context, h cdp.Handler) (identifier ScriptIdentifier, err error) {
	// execute
	var res AddScriptToEvaluateOnNewDocumentReturns
	err = h.Execute(ctxt, cdp.CommandPageAddScriptToEvaluateOnNewDocument, p, &res)
	if err != nil {
		return "", err
	}

	return res.Identifier, nil
}

// RemoveScriptToEvaluateOnNewDocumentParams removes given script from the
// list.
type RemoveScriptToEvaluateOnNewDocumentParams struct {
	Identifier ScriptIdentifier `json:"identifier"`
}

// RemoveScriptToEvaluateOnNewDocument removes given script from the list.
//
// parameters:
//   identifier
func RemoveScriptToEvaluateOnNewDocument(identifier ScriptIdentifier) *RemoveScriptToEvaluateOnNewDocumentParams {
	return &RemoveScriptToEvaluateOnNewDocumentParams{
		Identifier: identifier,
	}
}

// Do executes Page.removeScriptToEvaluateOnNewDocument against the provided context and
// target handler.
func (p *RemoveScriptToEvaluateOnNewDocumentParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	return h.Execute(ctxt, cdp.CommandPageRemoveScriptToEvaluateOnNewDocument, p, nil)
}

// SetAutoAttachToCreatedPagesParams controls whether browser will open a new
// inspector window for connected pages.
type SetAutoAttachToCreatedPagesParams struct {
	AutoAttach bool `json:"autoAttach"` // If true, browser will open a new inspector window for every page created from this one.
}

// SetAutoAttachToCreatedPages controls whether browser will open a new
// inspector window for connected pages.
//
// parameters:
//   autoAttach - If true, browser will open a new inspector window for every page created from this one.
func SetAutoAttachToCreatedPages(autoAttach bool) *SetAutoAttachToCreatedPagesParams {
	return &SetAutoAttachToCreatedPagesParams{
		AutoAttach: autoAttach,
	}
}

// Do executes Page.setAutoAttachToCreatedPages against the provided context and
// target handler.
func (p *SetAutoAttachToCreatedPagesParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	return h.Execute(ctxt, cdp.CommandPageSetAutoAttachToCreatedPages, p, nil)
}

// ReloadParams reloads given page optionally ignoring the cache.
type ReloadParams struct {
	IgnoreCache            bool   `json:"ignoreCache,omitempty"`            // If true, browser cache is ignored (as if the user pressed Shift+refresh).
	ScriptToEvaluateOnLoad string `json:"scriptToEvaluateOnLoad,omitempty"` // If set, the script will be injected into all frames of the inspected page after reload.
}

// Reload reloads given page optionally ignoring the cache.
//
// parameters:
func Reload() *ReloadParams {
	return &ReloadParams{}
}

// WithIgnoreCache if true, browser cache is ignored (as if the user pressed
// Shift+refresh).
func (p ReloadParams) WithIgnoreCache(ignoreCache bool) *ReloadParams {
	p.IgnoreCache = ignoreCache
	return &p
}

// WithScriptToEvaluateOnLoad if set, the script will be injected into all
// frames of the inspected page after reload.
func (p ReloadParams) WithScriptToEvaluateOnLoad(scriptToEvaluateOnLoad string) *ReloadParams {
	p.ScriptToEvaluateOnLoad = scriptToEvaluateOnLoad
	return &p
}

// Do executes Page.reload against the provided context and
// target handler.
func (p *ReloadParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	return h.Execute(ctxt, cdp.CommandPageReload, p, nil)
}

// NavigateParams navigates current page to the given URL.
type NavigateParams struct {
	URL            string         `json:"url"`                      // URL to navigate the page to.
	Referrer       string         `json:"referrer,omitempty"`       // Referrer URL.
	TransitionType TransitionType `json:"transitionType,omitempty"` // Intended transition type.
}

// Navigate navigates current page to the given URL.
//
// parameters:
//   url - URL to navigate the page to.
func Navigate(url string) *NavigateParams {
	return &NavigateParams{
		URL: url,
	}
}

// WithReferrer referrer URL.
func (p NavigateParams) WithReferrer(referrer string) *NavigateParams {
	p.Referrer = referrer
	return &p
}

// WithTransitionType intended transition type.
func (p NavigateParams) WithTransitionType(transitionType TransitionType) *NavigateParams {
	p.TransitionType = transitionType
	return &p
}

// NavigateReturns return values.
type NavigateReturns struct {
	FrameID cdp.FrameID `json:"frameId,omitempty"` // Frame id that will be navigated.
}

// Do executes Page.navigate against the provided context and
// target handler.
//
// returns:
//   frameID - Frame id that will be navigated.
func (p *NavigateParams) Do(ctxt context.Context, h cdp.Handler) (frameID cdp.FrameID, err error) {
	// execute
	var res NavigateReturns
	err = h.Execute(ctxt, cdp.CommandPageNavigate, p, &res)
	if err != nil {
		return "", err
	}

	return res.FrameID, nil
}

// StopLoadingParams force the page stop all navigations and pending resource
// fetches.
type StopLoadingParams struct{}

// StopLoading force the page stop all navigations and pending resource
// fetches.
func StopLoading() *StopLoadingParams {
	return &StopLoadingParams{}
}

// Do executes Page.stopLoading against the provided context and
// target handler.
func (p *StopLoadingParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	return h.Execute(ctxt, cdp.CommandPageStopLoading, nil, nil)
}

// GetNavigationHistoryParams returns navigation history for the current
// page.
type GetNavigationHistoryParams struct{}

// GetNavigationHistory returns navigation history for the current page.
func GetNavigationHistory() *GetNavigationHistoryParams {
	return &GetNavigationHistoryParams{}
}

// GetNavigationHistoryReturns return values.
type GetNavigationHistoryReturns struct {
	CurrentIndex int64              `json:"currentIndex,omitempty"` // Index of the current navigation history entry.
	Entries      []*NavigationEntry `json:"entries,omitempty"`      // Array of navigation history entries.
}

// Do executes Page.getNavigationHistory against the provided context and
// target handler.
//
// returns:
//   currentIndex - Index of the current navigation history entry.
//   entries - Array of navigation history entries.
func (p *GetNavigationHistoryParams) Do(ctxt context.Context, h cdp.Handler) (currentIndex int64, entries []*NavigationEntry, err error) {
	// execute
	var res GetNavigationHistoryReturns
	err = h.Execute(ctxt, cdp.CommandPageGetNavigationHistory, nil, &res)
	if err != nil {
		return 0, nil, err
	}

	return res.CurrentIndex, res.Entries, nil
}

// NavigateToHistoryEntryParams navigates current page to the given history
// entry.
type NavigateToHistoryEntryParams struct {
	EntryID int64 `json:"entryId"` // Unique id of the entry to navigate to.
}

// NavigateToHistoryEntry navigates current page to the given history entry.
//
// parameters:
//   entryID - Unique id of the entry to navigate to.
func NavigateToHistoryEntry(entryID int64) *NavigateToHistoryEntryParams {
	return &NavigateToHistoryEntryParams{
		EntryID: entryID,
	}
}

// Do executes Page.navigateToHistoryEntry against the provided context and
// target handler.
func (p *NavigateToHistoryEntryParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	return h.Execute(ctxt, cdp.CommandPageNavigateToHistoryEntry, p, nil)
}

// GetResourceTreeParams returns present frame / resource tree structure.
type GetResourceTreeParams struct{}

// GetResourceTree returns present frame / resource tree structure.
func GetResourceTree() *GetResourceTreeParams {
	return &GetResourceTreeParams{}
}

// GetResourceTreeReturns return values.
type GetResourceTreeReturns struct {
	FrameTree *FrameResourceTree `json:"frameTree,omitempty"` // Present frame / resource tree structure.
}

// Do executes Page.getResourceTree against the provided context and
// target handler.
//
// returns:
//   frameTree - Present frame / resource tree structure.
func (p *GetResourceTreeParams) Do(ctxt context.Context, h cdp.Handler) (frameTree *FrameResourceTree, err error) {
	// execute
	var res GetResourceTreeReturns
	err = h.Execute(ctxt, cdp.CommandPageGetResourceTree, nil, &res)
	if err != nil {
		return nil, err
	}

	return res.FrameTree, nil
}

// GetResourceContentParams returns content of the given resource.
type GetResourceContentParams struct {
	FrameID cdp.FrameID `json:"frameId"` // Frame id to get resource for.
	URL     string      `json:"url"`     // URL of the resource to get content for.
}

// GetResourceContent returns content of the given resource.
//
// parameters:
//   frameID - Frame id to get resource for.
//   url - URL of the resource to get content for.
func GetResourceContent(frameID cdp.FrameID, url string) *GetResourceContentParams {
	return &GetResourceContentParams{
		FrameID: frameID,
		URL:     url,
	}
}

// GetResourceContentReturns return values.
type GetResourceContentReturns struct {
	Content       string `json:"content,omitempty"`       // Resource content.
	Base64encoded bool   `json:"base64Encoded,omitempty"` // True, if content was served as base64.
}

// Do executes Page.getResourceContent against the provided context and
// target handler.
//
// returns:
//   content - Resource content.
func (p *GetResourceContentParams) Do(ctxt context.Context, h cdp.Handler) (content []byte, err error) {
	// execute
	var res GetResourceContentReturns
	err = h.Execute(ctxt, cdp.CommandPageGetResourceContent, p, &res)
	if err != nil {
		return nil, err
	}

	// decode
	var dec []byte
	if res.Base64encoded {
		dec, err = base64.StdEncoding.DecodeString(res.Content)
		if err != nil {
			return nil, err
		}
	} else {
		dec = []byte(res.Content)
	}
	return dec, nil
}

// SearchInResourceParams searches for given string in resource content.
type SearchInResourceParams struct {
	FrameID       cdp.FrameID `json:"frameId"`                 // Frame id for resource to search in.
	URL           string      `json:"url"`                     // URL of the resource to search in.
	Query         string      `json:"query"`                   // String to search for.
	CaseSensitive bool        `json:"caseSensitive,omitempty"` // If true, search is case sensitive.
	IsRegex       bool        `json:"isRegex,omitempty"`       // If true, treats string parameter as regex.
}

// SearchInResource searches for given string in resource content.
//
// parameters:
//   frameID - Frame id for resource to search in.
//   url - URL of the resource to search in.
//   query - String to search for.
func SearchInResource(frameID cdp.FrameID, url string, query string) *SearchInResourceParams {
	return &SearchInResourceParams{
		FrameID: frameID,
		URL:     url,
		Query:   query,
	}
}

// WithCaseSensitive if true, search is case sensitive.
func (p SearchInResourceParams) WithCaseSensitive(caseSensitive bool) *SearchInResourceParams {
	p.CaseSensitive = caseSensitive
	return &p
}

// WithIsRegex if true, treats string parameter as regex.
func (p SearchInResourceParams) WithIsRegex(isRegex bool) *SearchInResourceParams {
	p.IsRegex = isRegex
	return &p
}

// SearchInResourceReturns return values.
type SearchInResourceReturns struct {
	Result []*debugger.SearchMatch `json:"result,omitempty"` // List of search matches.
}

// Do executes Page.searchInResource against the provided context and
// target handler.
//
// returns:
//   result - List of search matches.
func (p *SearchInResourceParams) Do(ctxt context.Context, h cdp.Handler) (result []*debugger.SearchMatch, err error) {
	// execute
	var res SearchInResourceReturns
	err = h.Execute(ctxt, cdp.CommandPageSearchInResource, p, &res)
	if err != nil {
		return nil, err
	}

	return res.Result, nil
}

// SetDocumentContentParams sets given markup as the document's HTML.
type SetDocumentContentParams struct {
	FrameID cdp.FrameID `json:"frameId"` // Frame id to set HTML for.
	HTML    string      `json:"html"`    // HTML content to set.
}

// SetDocumentContent sets given markup as the document's HTML.
//
// parameters:
//   frameID - Frame id to set HTML for.
//   html - HTML content to set.
func SetDocumentContent(frameID cdp.FrameID, html string) *SetDocumentContentParams {
	return &SetDocumentContentParams{
		FrameID: frameID,
		HTML:    html,
	}
}

// Do executes Page.setDocumentContent against the provided context and
// target handler.
func (p *SetDocumentContentParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	return h.Execute(ctxt, cdp.CommandPageSetDocumentContent, p, nil)
}

// CaptureScreenshotParams capture page screenshot.
type CaptureScreenshotParams struct {
	Format      CaptureScreenshotFormat `json:"format,omitempty"`      // Image compression format (defaults to png).
	Quality     int64                   `json:"quality,omitempty"`     // Compression quality from range [0..100] (jpeg only).
	Clip        *Viewport               `json:"clip,omitempty"`        // Capture the screenshot of a given region only.
	FromSurface bool                    `json:"fromSurface,omitempty"` // Capture the screenshot from the surface, rather than the view. Defaults to true.
}

// CaptureScreenshot capture page screenshot.
//
// parameters:
func CaptureScreenshot() *CaptureScreenshotParams {
	return &CaptureScreenshotParams{}
}

// WithFormat image compression format (defaults to png).
func (p CaptureScreenshotParams) WithFormat(format CaptureScreenshotFormat) *CaptureScreenshotParams {
	p.Format = format
	return &p
}

// WithQuality compression quality from range [0..100] (jpeg only).
func (p CaptureScreenshotParams) WithQuality(quality int64) *CaptureScreenshotParams {
	p.Quality = quality
	return &p
}

// WithClip capture the screenshot of a given region only.
func (p CaptureScreenshotParams) WithClip(clip *Viewport) *CaptureScreenshotParams {
	p.Clip = clip
	return &p
}

// WithFromSurface capture the screenshot from the surface, rather than the
// view. Defaults to true.
func (p CaptureScreenshotParams) WithFromSurface(fromSurface bool) *CaptureScreenshotParams {
	p.FromSurface = fromSurface
	return &p
}

// CaptureScreenshotReturns return values.
type CaptureScreenshotReturns struct {
	Data string `json:"data,omitempty"` // Base64-encoded image data.
}

// Do executes Page.captureScreenshot against the provided context and
// target handler.
//
// returns:
//   data - Base64-encoded image data.
func (p *CaptureScreenshotParams) Do(ctxt context.Context, h cdp.Handler) (data []byte, err error) {
	// execute
	var res CaptureScreenshotReturns
	err = h.Execute(ctxt, cdp.CommandPageCaptureScreenshot, p, &res)
	if err != nil {
		return nil, err
	}

	// decode
	var dec []byte
	dec, err = base64.StdEncoding.DecodeString(res.Data)
	if err != nil {
		return nil, err
	}
	return dec, nil
}

// PrintToPDFParams print page as PDF.
type PrintToPDFParams struct {
	Landscape           bool    `json:"landscape,omitempty"`           // Paper orientation. Defaults to false.
	DisplayHeaderFooter bool    `json:"displayHeaderFooter,omitempty"` // Display header and footer. Defaults to false.
	PrintBackground     bool    `json:"printBackground,omitempty"`     // Print background graphics. Defaults to false.
	Scale               float64 `json:"scale,omitempty"`               // Scale of the webpage rendering. Defaults to 1.
	PaperWidth          float64 `json:"paperWidth,omitempty"`          // Paper width in inches. Defaults to 8.5 inches.
	PaperHeight         float64 `json:"paperHeight,omitempty"`         // Paper height in inches. Defaults to 11 inches.
	MarginTop           float64 `json:"marginTop,omitempty"`           // Top margin in inches. Defaults to 1cm (~0.4 inches).
	MarginBottom        float64 `json:"marginBottom,omitempty"`        // Bottom margin in inches. Defaults to 1cm (~0.4 inches).
	MarginLeft          float64 `json:"marginLeft,omitempty"`          // Left margin in inches. Defaults to 1cm (~0.4 inches).
	MarginRight         float64 `json:"marginRight,omitempty"`         // Right margin in inches. Defaults to 1cm (~0.4 inches).
	PageRanges          string  `json:"pageRanges,omitempty"`          // Paper ranges to print, e.g., '1-5, 8, 11-13'. Defaults to the empty string, which means print all pages.
}

// PrintToPDF print page as PDF.
//
// parameters:
func PrintToPDF() *PrintToPDFParams {
	return &PrintToPDFParams{}
}

// WithLandscape paper orientation. Defaults to false.
func (p PrintToPDFParams) WithLandscape(landscape bool) *PrintToPDFParams {
	p.Landscape = landscape
	return &p
}

// WithDisplayHeaderFooter display header and footer. Defaults to false.
func (p PrintToPDFParams) WithDisplayHeaderFooter(displayHeaderFooter bool) *PrintToPDFParams {
	p.DisplayHeaderFooter = displayHeaderFooter
	return &p
}

// WithPrintBackground print background graphics. Defaults to false.
func (p PrintToPDFParams) WithPrintBackground(printBackground bool) *PrintToPDFParams {
	p.PrintBackground = printBackground
	return &p
}

// WithScale scale of the webpage rendering. Defaults to 1.
func (p PrintToPDFParams) WithScale(scale float64) *PrintToPDFParams {
	p.Scale = scale
	return &p
}

// WithPaperWidth paper width in inches. Defaults to 8.5 inches.
func (p PrintToPDFParams) WithPaperWidth(paperWidth float64) *PrintToPDFParams {
	p.PaperWidth = paperWidth
	return &p
}

// WithPaperHeight paper height in inches. Defaults to 11 inches.
func (p PrintToPDFParams) WithPaperHeight(paperHeight float64) *PrintToPDFParams {
	p.PaperHeight = paperHeight
	return &p
}

// WithMarginTop top margin in inches. Defaults to 1cm (~0.4 inches).
func (p PrintToPDFParams) WithMarginTop(marginTop float64) *PrintToPDFParams {
	p.MarginTop = marginTop
	return &p
}

// WithMarginBottom bottom margin in inches. Defaults to 1cm (~0.4 inches).
func (p PrintToPDFParams) WithMarginBottom(marginBottom float64) *PrintToPDFParams {
	p.MarginBottom = marginBottom
	return &p
}

// WithMarginLeft left margin in inches. Defaults to 1cm (~0.4 inches).
func (p PrintToPDFParams) WithMarginLeft(marginLeft float64) *PrintToPDFParams {
	p.MarginLeft = marginLeft
	return &p
}

// WithMarginRight right margin in inches. Defaults to 1cm (~0.4 inches).
func (p PrintToPDFParams) WithMarginRight(marginRight float64) *PrintToPDFParams {
	p.MarginRight = marginRight
	return &p
}

// WithPageRanges paper ranges to print, e.g., '1-5, 8, 11-13'. Defaults to
// the empty string, which means print all pages.
func (p PrintToPDFParams) WithPageRanges(pageRanges string) *PrintToPDFParams {
	p.PageRanges = pageRanges
	return &p
}

// PrintToPDFReturns return values.
type PrintToPDFReturns struct {
	Data string `json:"data,omitempty"` // Base64-encoded pdf data.
}

// Do executes Page.printToPDF against the provided context and
// target handler.
//
// returns:
//   data - Base64-encoded pdf data.
func (p *PrintToPDFParams) Do(ctxt context.Context, h cdp.Handler) (data []byte, err error) {
	// execute
	var res PrintToPDFReturns
	err = h.Execute(ctxt, cdp.CommandPagePrintToPDF, p, &res)
	if err != nil {
		return nil, err
	}

	// decode
	var dec []byte
	dec, err = base64.StdEncoding.DecodeString(res.Data)
	if err != nil {
		return nil, err
	}
	return dec, nil
}

// StartScreencastParams starts sending each frame using the screencastFrame
// event.
type StartScreencastParams struct {
	Format        ScreencastFormat `json:"format,omitempty"`        // Image compression format.
	Quality       int64            `json:"quality,omitempty"`       // Compression quality from range [0..100].
	MaxWidth      int64            `json:"maxWidth,omitempty"`      // Maximum screenshot width.
	MaxHeight     int64            `json:"maxHeight,omitempty"`     // Maximum screenshot height.
	EveryNthFrame int64            `json:"everyNthFrame,omitempty"` // Send every n-th frame.
}

// StartScreencast starts sending each frame using the screencastFrame event.
//
// parameters:
func StartScreencast() *StartScreencastParams {
	return &StartScreencastParams{}
}

// WithFormat image compression format.
func (p StartScreencastParams) WithFormat(format ScreencastFormat) *StartScreencastParams {
	p.Format = format
	return &p
}

// WithQuality compression quality from range [0..100].
func (p StartScreencastParams) WithQuality(quality int64) *StartScreencastParams {
	p.Quality = quality
	return &p
}

// WithMaxWidth maximum screenshot width.
func (p StartScreencastParams) WithMaxWidth(maxWidth int64) *StartScreencastParams {
	p.MaxWidth = maxWidth
	return &p
}

// WithMaxHeight maximum screenshot height.
func (p StartScreencastParams) WithMaxHeight(maxHeight int64) *StartScreencastParams {
	p.MaxHeight = maxHeight
	return &p
}

// WithEveryNthFrame send every n-th frame.
func (p StartScreencastParams) WithEveryNthFrame(everyNthFrame int64) *StartScreencastParams {
	p.EveryNthFrame = everyNthFrame
	return &p
}

// Do executes Page.startScreencast against the provided context and
// target handler.
func (p *StartScreencastParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	return h.Execute(ctxt, cdp.CommandPageStartScreencast, p, nil)
}

// StopScreencastParams stops sending each frame in the screencastFrame.
type StopScreencastParams struct{}

// StopScreencast stops sending each frame in the screencastFrame.
func StopScreencast() *StopScreencastParams {
	return &StopScreencastParams{}
}

// Do executes Page.stopScreencast against the provided context and
// target handler.
func (p *StopScreencastParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	return h.Execute(ctxt, cdp.CommandPageStopScreencast, nil, nil)
}

// ScreencastFrameAckParams acknowledges that a screencast frame has been
// received by the frontend.
type ScreencastFrameAckParams struct {
	SessionID int64 `json:"sessionId"` // Frame number.
}

// ScreencastFrameAck acknowledges that a screencast frame has been received
// by the frontend.
//
// parameters:
//   sessionID - Frame number.
func ScreencastFrameAck(sessionID int64) *ScreencastFrameAckParams {
	return &ScreencastFrameAckParams{
		SessionID: sessionID,
	}
}

// Do executes Page.screencastFrameAck against the provided context and
// target handler.
func (p *ScreencastFrameAckParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	return h.Execute(ctxt, cdp.CommandPageScreencastFrameAck, p, nil)
}

// HandleJavaScriptDialogParams accepts or dismisses a JavaScript initiated
// dialog (alert, confirm, prompt, or onbeforeunload).
type HandleJavaScriptDialogParams struct {
	Accept     bool   `json:"accept"`               // Whether to accept or dismiss the dialog.
	PromptText string `json:"promptText,omitempty"` // The text to enter into the dialog prompt before accepting. Used only if this is a prompt dialog.
}

// HandleJavaScriptDialog accepts or dismisses a JavaScript initiated dialog
// (alert, confirm, prompt, or onbeforeunload).
//
// parameters:
//   accept - Whether to accept or dismiss the dialog.
func HandleJavaScriptDialog(accept bool) *HandleJavaScriptDialogParams {
	return &HandleJavaScriptDialogParams{
		Accept: accept,
	}
}

// WithPromptText the text to enter into the dialog prompt before accepting.
// Used only if this is a prompt dialog.
func (p HandleJavaScriptDialogParams) WithPromptText(promptText string) *HandleJavaScriptDialogParams {
	p.PromptText = promptText
	return &p
}

// Do executes Page.handleJavaScriptDialog against the provided context and
// target handler.
func (p *HandleJavaScriptDialogParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	return h.Execute(ctxt, cdp.CommandPageHandleJavaScriptDialog, p, nil)
}

// GetAppManifestParams [no description].
type GetAppManifestParams struct{}

// GetAppManifest [no description].
func GetAppManifest() *GetAppManifestParams {
	return &GetAppManifestParams{}
}

// GetAppManifestReturns return values.
type GetAppManifestReturns struct {
	URL    string              `json:"url,omitempty"` // Manifest location.
	Errors []*AppManifestError `json:"errors,omitempty"`
	Data   string              `json:"data,omitempty"` // Manifest content.
}

// Do executes Page.getAppManifest against the provided context and
// target handler.
//
// returns:
//   url - Manifest location.
//   errors
//   data - Manifest content.
func (p *GetAppManifestParams) Do(ctxt context.Context, h cdp.Handler) (url string, errors []*AppManifestError, data string, err error) {
	// execute
	var res GetAppManifestReturns
	err = h.Execute(ctxt, cdp.CommandPageGetAppManifest, nil, &res)
	if err != nil {
		return "", nil, "", err
	}

	return res.URL, res.Errors, res.Data, nil
}

// RequestAppBannerParams [no description].
type RequestAppBannerParams struct{}

// RequestAppBanner [no description].
func RequestAppBanner() *RequestAppBannerParams {
	return &RequestAppBannerParams{}
}

// Do executes Page.requestAppBanner against the provided context and
// target handler.
func (p *RequestAppBannerParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	return h.Execute(ctxt, cdp.CommandPageRequestAppBanner, nil, nil)
}

// SetControlNavigationsParams toggles navigation throttling which allows
// programatic control over navigation and redirect response.
type SetControlNavigationsParams struct {
	Enabled bool `json:"enabled"`
}

// SetControlNavigations toggles navigation throttling which allows
// programatic control over navigation and redirect response.
//
// parameters:
//   enabled
func SetControlNavigations(enabled bool) *SetControlNavigationsParams {
	return &SetControlNavigationsParams{
		Enabled: enabled,
	}
}

// Do executes Page.setControlNavigations against the provided context and
// target handler.
func (p *SetControlNavigationsParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	return h.Execute(ctxt, cdp.CommandPageSetControlNavigations, p, nil)
}

// ProcessNavigationParams should be sent in response to a
// navigationRequested or a redirectRequested event, telling the browser how to
// handle the navigation.
type ProcessNavigationParams struct {
	Response     NavigationResponse `json:"response"`
	NavigationID int64              `json:"navigationId"`
}

// ProcessNavigation should be sent in response to a navigationRequested or a
// redirectRequested event, telling the browser how to handle the navigation.
//
// parameters:
//   response
//   navigationID
func ProcessNavigation(response NavigationResponse, navigationID int64) *ProcessNavigationParams {
	return &ProcessNavigationParams{
		Response:     response,
		NavigationID: navigationID,
	}
}

// Do executes Page.processNavigation against the provided context and
// target handler.
func (p *ProcessNavigationParams) Do(ctxt context.Context, h cdp.Handler) (err error) {
	return h.Execute(ctxt, cdp.CommandPageProcessNavigation, p, nil)
}

// GetLayoutMetricsParams returns metrics relating to the layouting of the
// page, such as viewport bounds/scale.
type GetLayoutMetricsParams struct{}

// GetLayoutMetrics returns metrics relating to the layouting of the page,
// such as viewport bounds/scale.
func GetLayoutMetrics() *GetLayoutMetricsParams {
	return &GetLayoutMetricsParams{}
}

// GetLayoutMetricsReturns return values.
type GetLayoutMetricsReturns struct {
	LayoutViewport *LayoutViewport `json:"layoutViewport,omitempty"` // Metrics relating to the layout viewport.
	VisualViewport *VisualViewport `json:"visualViewport,omitempty"` // Metrics relating to the visual viewport.
	ContentSize    *dom.Rect       `json:"contentSize,omitempty"`    // Size of scrollable area.
}

// Do executes Page.getLayoutMetrics against the provided context and
// target handler.
//
// returns:
//   layoutViewport - Metrics relating to the layout viewport.
//   visualViewport - Metrics relating to the visual viewport.
//   contentSize - Size of scrollable area.
func (p *GetLayoutMetricsParams) Do(ctxt context.Context, h cdp.Handler) (layoutViewport *LayoutViewport, visualViewport *VisualViewport, contentSize *dom.Rect, err error) {
	// execute
	var res GetLayoutMetricsReturns
	err = h.Execute(ctxt, cdp.CommandPageGetLayoutMetrics, nil, &res)
	if err != nil {
		return nil, nil, nil, err
	}

	return res.LayoutViewport, res.VisualViewport, res.ContentSize, nil
}

// CreateIsolatedWorldParams creates an isolated world for the given frame.
type CreateIsolatedWorldParams struct {
	FrameID             cdp.FrameID `json:"frameId"`                       // Id of the frame in which the isolated world should be created.
	WorldName           string      `json:"worldName,omitempty"`           // An optional name which is reported in the Execution Context.
	GrantUniveralAccess bool        `json:"grantUniveralAccess,omitempty"` // Whether or not universal access should be granted to the isolated world. This is a powerful option, use with caution.
}

// CreateIsolatedWorld creates an isolated world for the given frame.
//
// parameters:
//   frameID - Id of the frame in which the isolated world should be created.
func CreateIsolatedWorld(frameID cdp.FrameID) *CreateIsolatedWorldParams {
	return &CreateIsolatedWorldParams{
		FrameID: frameID,
	}
}

// WithWorldName an optional name which is reported in the Execution Context.
func (p CreateIsolatedWorldParams) WithWorldName(worldName string) *CreateIsolatedWorldParams {
	p.WorldName = worldName
	return &p
}

// WithGrantUniveralAccess whether or not universal access should be granted
// to the isolated world. This is a powerful option, use with caution.
func (p CreateIsolatedWorldParams) WithGrantUniveralAccess(grantUniveralAccess bool) *CreateIsolatedWorldParams {
	p.GrantUniveralAccess = grantUniveralAccess
	return &p
}

// CreateIsolatedWorldReturns return values.
type CreateIsolatedWorldReturns struct {
	ExecutionContextID runtime.ExecutionContextID `json:"executionContextId,omitempty"` // Execution context of the isolated world.
}

// Do executes Page.createIsolatedWorld against the provided context and
// target handler.
//
// returns:
//   executionContextID - Execution context of the isolated world.
func (p *CreateIsolatedWorldParams) Do(ctxt context.Context, h cdp.Handler) (executionContextID runtime.ExecutionContextID, err error) {
	// execute
	var res CreateIsolatedWorldReturns
	err = h.Execute(ctxt, cdp.CommandPageCreateIsolatedWorld, p, &res)
	if err != nil {
		return 0, err
	}

	return res.ExecutionContextID, nil
}

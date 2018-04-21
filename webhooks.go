package nationbuilder

import (
	"fmt"
	"net/http"
)

// The webhook data structure
type Webhook struct {
	ID      string `json:"id,omitempty"`
	Version int    `json:"version,omitempty"`
	URL     string `json:"url,omitempty"`
	Event   string `json:"event,omitempty"`
}

// String() returns a textual representation of a webhook
func (w *Webhook) String() string {
	return fmt.Sprintf("Webhook ID: %s, Version: %d, URL: %s, Event: %s", w.ID, w.Version, w.URL, w.Event)
}

// Webhooks is a collection of Webhook objects with pagination support
type Webhooks struct {
	Results []*Webhook `json:"results"`
	Pagination
}

type webhookWrap struct {
	Webhook *Webhook `json:"webhook"`
}

// Retrieve a page of webhooks
func (n *Client) GetWebhooks(options *Options) (webhooks *Webhooks, result *Result) {
	req := n.getRequest("GET", "/webhooks", options)
	result = n.retrieve(req, &webhooks)

	return
}

// Create a webhook
func (n *Client) CreateWebhook(webhook *Webhook, options *Options) (newWebhook *Webhook, result *Result) {
	req := n.getRequest("POST", "/webhooks", options)
	ww := &webhookWrap{}
	result = n.create(&webhookWrap{webhook}, req, ww, http.StatusOK)
	newWebhook = ww.Webhook

	return
}

// Retrieve a webhook
func (n *Client) GetWebhook(id string, options *Options) (webhook *Webhook, result *Result) {
	u := fmt.Sprintf("/webhooks/%s", id)
	req := n.getRequest("GET", u, options)
	ww := &webhookWrap{}
	result = n.retrieve(req, ww)
	webhook = ww.Webhook

	return
}

// Update a webhook
func (n *Client) UpdateWebhook(id string, webhook *Webhook, options *Options) (updatedWebhook *Webhook, result *Result) {
	u := fmt.Sprintf("/webhooks/%s", id)
	req := n.getRequest("PUT", u, options)
	ww := &webhookWrap{}
	result = n.create(&webhookWrap{webhook}, req, ww, http.StatusOK)
	updatedWebhook = ww.Webhook

	return
}

// Delete a webhook
func (n *Client) DeleteWebhook(id string, options *Options) (result *Result) {
	u := fmt.Sprintf("/webhooks/%s", id)
	req := n.getRequest("DELETE", u, options)
	result = n.delete(req)

	return
}

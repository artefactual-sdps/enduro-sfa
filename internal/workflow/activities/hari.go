package activities

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	wferrors "github.com/artefactual-labs/enduro/internal/workflow/errors"
	"github.com/artefactual-labs/enduro/internal/workflow/manager"
)

var hariClient = &http.Client{
	Timeout: 10 * time.Second,
	Transport: &http.Transport{
		DialContext: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout: 5 * time.Second,
	},
}

type UpdateHARIActivity struct {
	manager *manager.Manager
}

func NewUpdateHARIActivity(m *manager.Manager) *UpdateHARIActivity {
	return &UpdateHARIActivity{manager: m}
}

type UpdateHARIActivityParams struct {
	SIPID        string
	Kind         string
	StoredAt     time.Time
	FullPath     string
	PipelineName string
}

func (a UpdateHARIActivity) Execute(ctx context.Context, params *UpdateHARIActivityParams) error {
	var err error
	params.Kind, err = convertKind(params.Kind)
	if err != nil {
		return wferrors.NonRetryableError(fmt.Errorf("error validating kind attribute: %v", err))
	}

	if params.PipelineName == "" {
		params.PipelineName = "<unnamed>"
	}

	apiURL, err := a.url()
	if err != nil {
		return wferrors.NonRetryableError(fmt.Errorf("error in URL construction: %v", err))
	}

	mock, _ := manager.HookAttrBool(a.manager.Hooks, "hari", "mock")
	if mock {
		ts := a.buildMock()
		defer ts.Close()
		apiURL = ts.URL
	}

	var path = a.avlxml(filepath.Join(params.FullPath, params.Kind))
	if path == "" {
		return wferrors.NonRetryableError(fmt.Errorf("error reading AVLXML file: not found"))
	}

	blob, err := ioutil.ReadFile(path)
	if err != nil {
		return wferrors.NonRetryableError(fmt.Errorf("error reading AVLXML file: %v", err))
	}

	if err := a.sendRequest(ctx, blob, apiURL, params); err != nil {
		return fmt.Errorf("error sending request: %v", err)
	}

	return nil
}

// avlxml attempts to find the AVLXML document in multiple known locations.
func (a UpdateHARIActivity) avlxml(prefix string) string {
	locs := []string{"journal/avlxml.xml", "Journal/avlxml.xml"}
	for _, loc := range locs {
		path := filepath.Join(prefix, loc)
		if stat, err := os.Stat(path); err == nil && !stat.IsDir() {
			return path
		}
	}
	return ""
}

// url returns the HARI URL of the API endpoint for AVLXML submission.
func (a UpdateHARIActivity) url() (string, error) {
	p, _ := url.Parse("/v1/hari/avlxml")

	b, err := manager.HookAttrString(a.manager.Hooks, "hari", "baseURL")
	if err != nil {
		return "", fmt.Errorf("error looking up baseURL configuration attribute: %v", err)
	}

	bu, err := url.Parse(b)
	if err != nil {
		return "", fmt.Errorf("error looking up baseURL configuration attribute: %v", err)
	}

	return bu.ResolveReference(p).String(), nil
}

func (a UpdateHARIActivity) sendRequest(ctx context.Context, blob []byte, apiURL string, params *UpdateHARIActivityParams) error {
	payload := &avlRequest{
		XML:       blob,
		Message:   fmt.Sprintf("AVLXML was processed by Archivematica pipeline %s", params.PipelineName),
		Type:      strings.ToLower(params.Kind),
		Timestamp: params.StoredAt,
		AIPID:     params.SIPID,
	}

	var buffer bytes.Buffer
	if err := json.NewEncoder(&buffer).Encode(payload); err != nil {
		return fmt.Errorf("error encoding payload: %v", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", apiURL, &buffer)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "Enduro")

	resp, err := hariClient.Do(req)
	if err != nil {
		return err
	}

	switch {
	case resp.StatusCode >= 200 && resp.StatusCode <= 299:
		err = nil
	default:
		err = fmt.Errorf("unexpected response status: %s", resp.Status)
	}

	return err
}

// buildMock returns a test server used when HARI's API is not available.
func (a UpdateHARIActivity) buildMock() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		a.manager.Logger.Info(
			"Request received",
			"method", r.Method,
			"path", r.URL.Path,
		)

		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"message": "ok"}`)
	}))
}

type avlRequest struct {
	XML       []byte    `json:"xml"`       // AVLXML document encoded using base64.
	Message   string    `json:"message"`   // E.g.: "AVLXML was processed by DPJ Archivematica pipeline"
	Type      string    `json:"type"`      // Lowercase. E.g.: "dpj", "epj", "other" or "avlxml".
	Timestamp time.Time `json:"timestamp"` // RFC3339. E.g. "2006-01-02T15:04:05Z07:00".
	AIPID     string    `json:"aip_id"`
}

var knownKinds = []string{
	"DPJ", "EPJ", "AVLXML", "OTHER",
}

func convertKind(kind string) (string, error) {
	if kind == "" {
		return "", errors.New("empty")
	}

	// Convert into capital letters, e.g. epj-sip => EPJ-SIP.
	kind = strings.ToUpper(kind)

	const suffix = "-SIP"
	if !strings.HasSuffix(kind, suffix) {
		return "", fmt.Errorf("attribute (%s) does not containt suffix (\"-SIP\")", kind)
	}
	kind = strings.TrimSuffix(kind, "-SIP")

	var known bool
	for _, k := range knownKinds {
		if k == kind {
			known = true
			break
		}
	}
	if !known {
		return "", fmt.Errorf("attribute (%s) is unexpected/unknown", kind)
	}

	return kind, nil
}
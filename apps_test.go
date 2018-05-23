package ghost

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestClientGetApps(t *testing.T) {
	mux, server, client := setupTest()
	defer teardownTest(server)

	mux.HandleFunc("/apps", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `
			{
				"_links": {},
				"_meta": {},
				"_items": [{},{}]
			}
		`)
	})

	apps, err := client.GetApps()
	if err != nil {
		t.Errorf("apps.GetApps returned error: %v", err)
	}

	want := Apps{EveCollectionMetadata{
		Links: &struct {
			Parent Link "json:\"parent,omitempty\""
			Self   Link "json:\"self,omitempty\""
		}{},
		Meta: &struct {
			MaxResults int64 `json:"max_results,omitempty"`
			Page       int64 `json:"page,omitempty"`
			Total      int64 `json:"total,omitempty"`
		}{}},
		[]App{{}, {}}}
	if !reflect.DeepEqual(want, apps) {
		t.Errorf("apps.GetApps returned %+v, want %+v", apps, want)
	}
}

func TestClientGetApp(t *testing.T) {
	appID := "5af94fb79cfca46315de31a3"

	tests := []struct {
		url           string
		clientResp    string
		expectedResp  App
		expectedError bool
	}{
		// Valid test case
		{
			url: "/apps/" + appID,
			clientResp: `
			{
				"name": "test"
			}`,
			expectedResp: App{
				Name: "test",
			},
			expectedError: false,
		},
		// Invalid value
		{
			url: "/apps/" + appID,
			clientResp: `
			{
				"name": "test"
			}`,
			expectedResp: App{
				Name: "tes",
			},
			expectedError: true,
		},
		// Invalid AppID
		{
			url:           "/apps/invalid" + appID,
			clientResp:    `{}`,
			expectedResp:  App{},
			expectedError: true,
		},
	}

	for _, test := range tests {
		mux, server, client := setupTest()
		defer teardownTest(server)

		mux.HandleFunc(test.url, func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			fmt.Fprint(w, test.clientResp)
		})

		app, err := client.GetApp(appID)
		if err != nil && !test.expectedError {
			t.Errorf("apps.GetApp returned error: %v", err)
		}
		if !test.expectedError && !reflect.DeepEqual(test.expectedResp, app) {
			t.Errorf("apps.GetApp returned %+v, want %+v", app, test.expectedResp)
		}
	}
}

func TestClientCreateApp(t *testing.T) {
	tests := []struct {
		url           string
		clientResp    string
		expectedResp  EveItemMetadata
		expectedError bool
		app           App
	}{
		// Valid test case
		{
			url: "/apps",
			clientResp: `
			{
				"_id": "test"
			}`,
			expectedResp: EveItemMetadata{
				ID: "test",
			},
			expectedError: false,
			app: App{
				Name: "test",
			},
		},
		// Invalid value
		{
			url: "/apps",
			clientResp: `
			{
				"_id": "test"
			}`,
			expectedResp: EveItemMetadata{
				ID: "te",
			},
			expectedError: true,
			app: App{
				Name: "test",
			},
		},
		// Invalid path
		{
			url: "/apps/invalid",
			clientResp: `
			{
				"_id": "test"
			}`,
			expectedResp: EveItemMetadata{
				ID: "test",
			},
			expectedError: true,
			app: App{
				Name: "test",
			},
		},
	}

	for _, test := range tests {
		mux, server, client := setupTest()
		defer teardownTest(server)

		mux.HandleFunc(test.url, func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "POST")
			fmt.Fprint(w, test.clientResp)
		})

		app, err := client.CreateApp(test.app)
		if err != nil && !test.expectedError {
			t.Errorf("apps.CreateApp returned error: %v", err)
		}
		if !test.expectedError && !reflect.DeepEqual(test.expectedResp, app) {
			t.Errorf("apps.CreateApp returned %+v, want %+v", app, test.expectedResp)
		}
	}
}

func getStringAddr(s string) *string {
	return &s
}

func TestClientUpdateApp(t *testing.T) {
	appID := "5af94fb79cfca46315de31a3"
	appETAG := "1"

	tests := []struct {
		url           string
		clientResp    string
		expectedResp  EveItemMetadata
		expectedError bool
		app           App
	}{
		// Valid test case
		{
			url: "/apps/" + appID,
			clientResp: `
			{
				"_etag": "2"
			}`,
			expectedResp: EveItemMetadata{
				Etag: getStringAddr("2"),
			},
			expectedError: false,
			app: App{
				Name: "test",
			},
		},
		// Invalid value
		{
			url: "/apps/" + appID,
			clientResp: `
			{
				"_etag": "2"
			}`,
			expectedResp: EveItemMetadata{
				Etag: getStringAddr("1"),
			},
			expectedError: true,
			app: App{
				Name: "test",
			},
		},
		// Invalid AppID
		{
			url:           "/apps/invalid",
			clientResp:    `{}`,
			expectedResp:  EveItemMetadata{},
			expectedError: true,
			app: App{
				Name: "test",
			},
		},
	}

	for _, test := range tests {
		mux, server, client := setupTest()
		defer teardownTest(server)

		mux.HandleFunc(test.url, func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "PATCH")
			fmt.Fprint(w, test.clientResp)
		})

		app, err := client.UpdateApp(&test.app, appID, appETAG)
		if err != nil && !test.expectedError {
			t.Errorf("apps.UpdateApp returned error: %v", err)
		}
		if !test.expectedError && !reflect.DeepEqual(test.expectedResp, app) {
			t.Errorf("apps.UpdateApp returned %+v, want %+v", app, test.expectedResp)
		}
	}
}

func TestClientDeleteApp(t *testing.T) {
	appID := "5af94fb79cfca46315de31a3"
	appETAG := "1"

	tests := []struct {
		url           string
		expectedError bool
	}{
		// Valid test case
		{
			url:           "/apps/" + appID,
			expectedError: false,
		},
		// Invalid AppID
		{
			url:           "/apps/invalid",
			expectedError: true,
		},
	}

	for _, test := range tests {
		mux, server, client := setupTest()
		defer teardownTest(server)

		mux.HandleFunc(test.url, func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "DELETE")
			fmt.Fprint(w, "")
		})

		err := client.DeleteApp(appID, appETAG)
		if err != nil && !test.expectedError {
			t.Errorf("apps.DeleteApp returned error: %v", err)
		}
	}
}

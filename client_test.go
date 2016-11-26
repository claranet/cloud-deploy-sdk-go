package ghost

import "testing"

func TestClientGetApps(t *testing.T) {
	c := NewClient("https://demo.ghost.***REMOVED***", "demo", "***REMOVED***")
	err := c.GetApps()
	if err != nil {
		t.Fatalf("error: %v", err)
	}
}

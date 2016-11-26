package ghost

import (
	"fmt"
	"testing"
)

var c = NewClient("https://demo.ghost.***REMOVED***", "demo", "***REMOVED***")

func TestClientGetApps(t *testing.T) {
	fmt.Println("Testing Ghost client get all apps")
	apps, err := c.GetApps()
	if err == nil {
		fmt.Println("All apps retrieved:")
		fmt.Println(apps)
		fmt.Println()
	} else {
		t.Fatalf("error: %v", err)
	}
}

func TestClientGetApp(t *testing.T) {
	fmt.Println("Testing Ghost client get single app")
	app, err := c.GetApp("55cf9ce5fde8dd0521358a19")
	if err == nil {
		fmt.Println("App retrieved:")
		fmt.Println(app)
		fmt.Println()
	} else {
		t.Fatalf("error: %v", err)
	}
}

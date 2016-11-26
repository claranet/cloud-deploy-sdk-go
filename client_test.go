package ghost

import (
	"crypto/rand"
	"fmt"
	"testing"
)

func pseudo_uuid() (uuid string) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err == nil {
		uuid = fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	}
	return
}

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

func TestClientCreateApp(t *testing.T) {
	fmt.Println("Testing Ghost client create app")
	app := App{
		Name: "test-" + pseudo_uuid(),
		Env:  "test",
		Role: "webfront",

		Region:       "eu-west-1",
		InstanceType: "t2.nano",
		VpcID:        "vpc-123456",

		LogNotifications: []string{"ghost-demo@domain.com"},

		BuildInfos: BuildInfos{
			SourceAmi:   "ami-123456",
			SshUsername: "admin",
			SubnetID:    "subnet-123456",
		},

		EnvironmentInfos: EnvironmentInfos{
			InstanceProfile: "test-instance-profile",
			KeyName:         "test-key-name",
			OptionalVolumes: []OptionalVolume{},
			RootBlockDevice: RootBlockDevice{Name: "/dev/xvda"},
			SecurityGroups:  []string{"sg-123456"},
			SubnetIDs:       []string{"subnet-123456"},
		},

		Features: []Feature{
			{
				Name:    "nginx",
				Version: "1.10",
			},
		},
		Modules: []Module{
			{
				Name:    "testmod",
				GitRepo: "git@bitbucket.org/morea/testmod",
				Scope:   "system",
				Path:    "/tmp/path",
			},
		},
	}

	id, err := c.CreateApp(app)
	if err == nil {
		fmt.Println("App created:" + id)
		fmt.Println()
	} else {
		t.Fatalf("error: %v", err)
	}
}

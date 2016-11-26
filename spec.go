package ghost

// Link struct
type Link struct {
	Href  string `json:"href"`
	Title string `json:"title"`
}

// Ghost App's feature struct
type Feature struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// Ghost App's module struct
type Module struct {
	GitRepo     string `json:"git_repo"`
	Initialized bool   `json:"initialized"`
	Name        string `json:"name"`
	Path        string `json:"path"`
	PreDeploy   string `json:"pre_deploy"`
	Scope       string `json:"scope"`
}

// Ghost App struct
type App struct {
	Etag          string `json:"_etag"`
	ID            string `json:"_id"`
	Created       string `json:"_created"`
	Updated       string `json:"_updated"`
	Version       int64  `json:"_version"`
	LatestVersion int64  `json:"_latest_version"`
	User          string `json:"user"`

	Name string `json:"name"`
	Env  string `json:"env"`
	Role string `json:"role"`

	Region       string `json:"region"`
	InstanceType string `json:"instance_type"`
	VpcID        string `json:"vpc_id"`

	LogNotifications []string `json:"log_notifications"`

	BuildInfos struct {
		SourceAmi   string `json:"source_ami"`
		SshUsername string `json:"ssh_username"`
		SubnetID    string `json:"subnet_id"`
	} `json:"build_infos"`

	EnvironmentInfos struct {
		InstanceProfile string        `json:"instance_profile"`
		KeyName         string        `json:"key_name"`
		OptionalVolumes []interface{} `json:"optional_volumes"`
		RootBlockDevice struct {
		} `json:"root_block_device"`
		SecurityGroups []string `json:"security_groups"`
		SubnetIDs      []string `json:"subnet_ids"`
	} `json:"environment_infos"`

	Features []Feature `json:"features"`

	Modules []Module `json:"modules"`

	Links struct {
		Self Link `json:"self"`
	} `json:"_links"`
}

type Apps struct {
	Items []App `json:"_items"`

	Links struct {
		Parent Link `json:"parent"`
		Self   Link `json:"self"`
	} `json:"_links"`

	Meta struct {
		MaxResults int64 `json:"max_results"`
		Page       int64 `json:"page"`
		Total      int64 `json:"total"`
	} `json:"_meta"`
}

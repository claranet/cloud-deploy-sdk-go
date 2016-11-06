package main

type Apps struct {
	Items []struct {
		Created       string `json:"_created"`
		Etag          string `json:"_etag"`
		ID            string `json:"_id"`
		LatestVersion int64  `json:"_latest_version"`
		Links         struct {
			Self struct {
				Href  string `json:"href"`
				Title string `json:"title"`
			} `json:"self"`
		} `json:"_links"`
		Updated   string `json:"_updated"`
		Version   int64  `json:"_version"`
		Autoscale struct {
			Name string `json:"name"`
		} `json:"autoscale"`
		BuildInfos struct {
			SourceAmi   string `json:"source_ami"`
			SshUsername string `json:"ssh_username"`
			SubnetID    string `json:"subnet_id"`
		} `json:"build_infos"`
		Env              string `json:"env"`
		EnvironmentInfos struct {
			InstanceProfile string        `json:"instance_profile"`
			KeyName         string        `json:"key_name"`
			OptionalVolumes []interface{} `json:"optional_volumes"`
			RootBlockDevice struct {
			} `json:"root_block_device"`
			SecurityGroups []string `json:"security_groups"`
			SubnetIDs      []string `json:"subnet_ids"`
		} `json:"environment_infos"`
		Features []struct {
			Name    string  `json:"name"`
			Version float64 `json:"version,string"`
		} `json:"features"`
		InstanceType     string   `json:"instance_type"`
		LogNotifications []string `json:"log_notifications"`
		Modules          []struct {
			GitRepo     string `json:"git_repo"`
			Initialized bool   `json:"initialized"`
			Name        string `json:"name"`
			Path        string `json:"path"`
			PreDeploy   string `json:"pre_deploy"`
			Scope       string `json:"scope"`
		} `json:"modules"`
		Name   string `json:"name"`
		Region string `json:"region"`
		Role   string `json:"role"`
		User   string `json:"user"`
		VpcID  string `json:"vpc_id"`
	} `json:"_items"`
	Links struct {
		Parent struct {
			Href  string `json:"href"`
			Title string `json:"title"`
		} `json:"parent"`
		Self struct {
			Href  string `json:"href"`
			Title string `json:"title"`
		} `json:"self"`
	} `json:"_links"`
	Meta struct {
		MaxResults int64 `json:"max_results"`
		Page       int64 `json:"page"`
		Total      int64 `json:"total"`
	} `json:"_meta"`
}

// App struct describe application in Ghost
type App struct {
	Etag          string `json:"_etag"`
	ID            string `json:"_id"`
	LatestVersion int64  `json:"_latest_version"`
	Autoscale     struct {
		Name string `json:"name"`
	} `json:"autoscale"`
	BuildInfos struct {
		SourceAmi   string `json:"source_ami"`
		SSHUsername string `json:"ssh_username"`
		SubnetID    string `json:"subnet_id"`
	} `json:"build_infos"`
	Env              string `json:"env"`
	EnvironmentInfos struct {
		InstanceProfile string        `json:"instance_profile"`
		KeyName         string        `json:"key_name"`
		OptionalVolumes []interface{} `json:"optional_volumes"`
		RootBlockDevice struct {
		} `json:"root_block_device"`
		SecurityGroups []string `json:"security_groups"`
		SubnetIDs      []string `json:"subnet_ids"`
	} `json:"environment_infos"`
	Features []struct {
		Name    string  `json:"name"`
		Version float64 `json:"version,string"`
	} `json:"features"`
	InstanceType     string   `json:"instance_type"`
	LogNotifications []string `json:"log_notifications"`
	Modules          []struct {
		GitRepo     string `json:"git_repo"`
		Initialized bool   `json:"initialized"`
		Name        string `json:"name"`
		Path        string `json:"path"`
		PreDeploy   string `json:"pre_deploy"`
		Scope       string `json:"scope"`
	} `json:"modules"`
	Name   string `json:"name"`
	Region string `json:"region"`
	Role   string `json:"role"`
	User   string `json:"user"`
	VpcID  string `json:"vpc_id"`
}

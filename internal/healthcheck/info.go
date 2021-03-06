package healthcheck

var (
	CommitHash string
	Branch     string
	BuildTime  string
)

// Info - healthcheck info.
type Info struct {
	App *BuildInfo `json:"app"`
	// rename to your actual db
	DB *Status `json:"db"`
	// add here list of needed services
}

// Status - status of healthcheck entity.
type Status struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

type BuildInfo struct {
	CommitHash string `json:"commit_hash"`
	Branch     string `json:"branch"`
	BuildTime  string `json:"build_time"`
}

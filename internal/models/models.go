package models

// Conf - web gui config
type Conf struct {
	Host       string
	Port       string
	Theme      string
	Color      string
	Icon       string
	ConfPath   string
	DBPath     string
	DBTrimDays string
	YamlPath   string
	NodePath   string
	ColorOn    string
	ColorOff   string
	BtnWidth   string
	WebRefresh string
	Quit       chan bool
}

// Host - panel element
type Host struct {
	Name  string `yaml:"name"`
	Addr  string `yaml:"addr"`
	Port  string `yaml:"port"`
	URL   string `yaml:"url"`
	Icon  string `yaml:"icon"`
	State bool   `yaml:"state"`
}

// Panel - tab element
type Panel struct {
	Name    string       `yaml:"name"`
	Scan    bool         `yaml:"scan"`
	Timeout string       `yaml:"timeout"`
	Hosts   map[int]Host `yaml:"hosts"`
}

// Tab - board element
type Tab struct {
	Name    string         `yaml:"name"`
	Refresh string         `yaml:"refresh"`
	Panels  map[int]string `yaml:"panels"`
}

// MonPanel - uptime element
type MonPanel struct {
	Retries int      `yaml:"retries"`
	Notify  []string `yaml:"notify"`
}

// MonData - for monitoring results
type MonData struct {
	ID     int    `db:"ID"`
	Panel  string `db:"PANEL"`
	Host   string `db:"HOST"`
	Addr   string `db:"ADDR"`
	Port   string `db:"PORT"`
	Date   string `db:"DATE"`
	Time   string `db:"TIME"`
	Color  string `db:"COLOR"`
	State  bool   `db:"STATE"`
	Notify string `db:"NOTIFY"`
}

// Uptime - board element
type Uptime struct {
	Enabled bool                `yaml:"enabled"`
	Show    int                 `yaml:"show"`
	Notify  map[string]string   `yaml:"notify"`
	Panels  map[string]MonPanel `yaml:"panels"`
}

// Links - all links
type Links struct {
	Tabs   map[int]Tab      `yaml:"tabs"`
	Panels map[string]Panel `yaml:"panels"`
	Uptime Uptime           `yaml:"uptime"`
}

// GuiData - web gui data
type GuiData struct {
	Config     Conf
	Themes     []string
	Links      Links
	Panels     map[int]Panel
	CurrentTab string
	TabEdit    int
	UptimeMon  []MonData
	Version    string
}

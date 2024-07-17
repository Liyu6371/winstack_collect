// WinStackCollect, 暂时不考虑有多个 WinStackCloud 实例的情况；暂时只考虑一个的状态下
// 周期调度上报数据到 gse

package winstack

type WinStack struct {
	Concurrency int     `yaml:"concurrency"`
	Period      string  `yaml:"period"`
	Clouds      []Cloud `yaml:"clouds"`
}

type Cloud struct {
	ID       int      `yaml:"id"`
	Account  string   `yaml:"account"`
	Password string   `yaml:"password"`
	Server   string   `yaml:"server"`
	Clusters []string `yaml:"cluster"`
	Hosts    []string `yaml:"host"`
	Storage  []string `yaml:"storage"`
	VM       []string `yaml:"vm"`
}

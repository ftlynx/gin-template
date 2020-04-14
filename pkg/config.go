package pkg


var (
	pkgs []PKG
)

// pkg下面模块配置参数
type PKG interface {
	Config() error
}

func SetConfig(pkg PKG) {
	pkgs = append(pkgs, pkg)
}

func InitConfig() error {
	for _, p := range pkgs {
		if err := p.Config(); err != nil {
			return err
		}
	}

	return nil
}
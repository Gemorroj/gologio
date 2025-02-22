package gologio

type Config struct {
	Apache struct {
		Path   string `yaml:"path"`
		Format struct {
			Date    string `yaml:"date"`
			Type    string `yaml:"type"`
			Message string `yaml:"message"`
			Client  string `yaml:"client"`
		}
		Cast struct {
			Date string `yaml:"date"`
		}
	} `yaml:"apache"`
	Nginx struct {
		Path   string `yaml:"path"`
		Format struct {
			Date    string `yaml:"date"`
			Type    string `yaml:"type"`
			Message string `yaml:"message"`
			Client  string `yaml:"client"`
			Server  string `yaml:"server"`
			Request string `yaml:"request"`
			Host    string `yaml:"host"`
		}
		Cast struct {
			Date string `yaml:"date"`
		}
	} `yaml:"nginx"`
	PhpFpm struct {
		Path   string `yaml:"path"`
		Format struct {
			Date    string `yaml:"date"`
			Type    string `yaml:"type"`
			Pool    string `yaml:"pool"`
			Child   string `yaml:"child"`
			Message string `yaml:"message"`
		}
		Cast struct {
			Date string `yaml:"date"`
		}
	} `yaml:"php_fpm"`
	Php struct {
		Path   string `yaml:"path"`
		Format struct {
			Date    string `yaml:"date"`
			Type    string `yaml:"type"`
			Message string `yaml:"message"`
			File    string `yaml:"file"`
			Line    string `yaml:"line"`
		}
		Cast struct {
			Date string `yaml:"date"`
		}
	} `yaml:"php"`
	Mysql struct {
		Path   string `yaml:"path"`
		Format struct {
			Date    string `yaml:"date"`
			Type    string `yaml:"type"`
			Message string `yaml:"message"`
			Thread  string `yaml:"thread"`
		}
		Cast struct {
			Date string `yaml:"date"`
		}
	} `yaml:"mysql"`
}

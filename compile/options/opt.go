package options

type Options struct {
	options
}

type options struct {
	template, content, output   string
	templateSuffix, logFileName string
}

func NewOptions(template, content, output string, templateSuff, logName string) Options {
	opt := options{template: template, content: content, output: output}
	opt.templateSuffix = templateSuff
	opt.logFileName = logName

	return Options{opt}
}

func (o options) Content() string {
	return o.content
}

func (o options) Template() string {
	return o.template
}

func (o options) Output() string {
	return o.output
}

func (o options) TemplateSuffix() string {
	return o.templateSuffix
}

func (o options) LogFileName() string {
	return o.logFileName
}

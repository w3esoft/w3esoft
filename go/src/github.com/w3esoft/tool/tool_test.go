package tool
import "testing"
func Test(t *testing.T) {
	p:="packages/";
	ps  := map[string]string {
		"@atomicburst/breadcrumb": p + "atomicburst/packages/breadcrumb",
			"@atomicburst/button": p + "atomicburst/packages/button",
			"@atomicburst/calendar": p + "atomicburst/packages/calendar",
			"@atomicburst/chart": p + "atomicburst/packages/chart",
			"@atomicburst/container": p + "atomicburst/packages/container",
			"@atomicburst/container.layout": p + "atomicburst/packages/container.layout",
			"@atomicburst/container.modal": p + "atomicburst/packages/container.modal",
			"@atomicburst/container.panel": p + "atomicburst/packages/container.panel",
			"@atomicburst/container.viewport": p + "atomicburst/packages/container.viewport",
			"@atomicburst/core": p + "atomicburst/packages/core",
			"@atomicburst/data": p + "atomicburst/packages/data",
			"@atomicburst/datatable": p + "atomicburst/packages/datatable",
			"@atomicburst/form.field": p + "atomicburst/packages/form.field",
			"@atomicburst/form.field.checkbox": p + "atomicburst/packages/form.field.checkbox",
			"@atomicburst/form.field.combo": p + "atomicburst/packages/form.field.combo",
			"@atomicburst/form.field.date": p + "atomicburst/packages/form.field.date",
			"@atomicburst/form.field.inputfile": p + "atomicburst/packages/form.field.inputfile",
			"@atomicburst/form.field.picker.date": p + "atomicburst/packages/form.field.picker.date",
			"@atomicburst/form.field.range": p + "atomicburst/packages/form.field.range",
			"@atomicburst/form.field.text": p + "atomicburst/packages/form.field.text",
			"@atomicburst/form.field.textarea": p + "atomicburst/packages/form.field.textarea",
			"@atomicburst/form.mask": p + "atomicburst/packages/form.mask",
			"@atomicburst/form.validation": p + "atomicburst/packages/form.validation",
			"@atomicburst/imagefield": p + "atomicburst/packages/imagefield",
			"@atomicburst/locale.pt_br": p + "atomicburst/packages/locale.pt_br",
			"@atomicburst/menu": p + "atomicburst/packages/menu",
			"@atomicburst/pagination": p + "atomicburst/packages/pagination",
			"@atomicburst/toolbar": p + "atomicburst/packages/toolbar",
			"@damsistemas/core": p + "damsistemas/packages/core",
			"@damsistemas/crm": p + "damsistemas/packages/crm",
			"@damsistemas/damcloud": p + "damsistemas/packages/damcloud",
			"@damsistemas/damgerenciador": p + "damsistemas/packages/damgerenciador",
			"@damsistemas/damonline": p + "damsistemas/packages/damonline",
			"@damsistemas/site": p + "damsistemas/packages/site",
	};
	files :=[]string{}
	for _,k :=range ps  {
		files =append(files,k+ "/**.html")
	}
	rootDir:="E:\\desenvolvimento3\\dam-online\\angular"
	exclude :=[]string{"**/node_modules"}
	Find(rootDir,files,exclude)
}
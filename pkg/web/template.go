package web

import (
	"html/template"
)

func (ws *WebServer) registerTemplates() error {
	ws.registerTemplateComponent("navbar")

	ws.registerViewTemplate("index")
	ws.registerViewTemplate("dashboard")
	ws.registerViewTemplate("projects_list")
	ws.registerViewTemplate("project_create")
	ws.registerViewTemplate("project_details")
	ws.registerViewTemplate("testruns_list")
	ws.registerViewTemplate("suite_create")
	ws.registerViewTemplate("suite_details")

	return nil
}

func (ws *WebServer) registerTemplateComponent(templateName string) {
	ws.templateComponents = append(ws.templateComponents, "template/component/"+templateName+".html")
}

func (ws *WebServer) registerViewTemplate(templateName string) {
	tmpls := []string{"template/" + templateName + ".html", "template/base.html"}
	tmpls = append(tmpls, ws.templateComponents...)
	ws.templates[templateName] = template.Must(template.ParseFS(templateFiles, tmpls...))
}

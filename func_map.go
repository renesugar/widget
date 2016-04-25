package widget

var funcMap = map[string]interface{}{
	"widget_available_scopes": func() []*Scope {
		if len(registeredScopes) > 0 {
			return append([]*Scope{{Name: "Default"}}, registeredScopes...)
		}
		return []*Scope{}
	},
}

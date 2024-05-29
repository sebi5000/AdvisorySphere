// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.696
package project_request

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func ProjectRequest() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<script>\n        function readyToClear() {\n            return document.getElementById('description').value.length === 0\n        }\n    </script><div class=\"row\"><div class=\"col-sm-4\" x-data=\"{ submitStatus: false, clearStatus: false }\"><div class=\"box\"><form id=\"project_request\" hx-post=\"project_request\" hx-trigger=\"submit\" hx-target=\"#project_match\" x-on:submit=\"submitStatus = false\"><fieldset><label>Projektbeschreibung <textarea id=\"description\" name=\"description\" placeholder=\"Beschreibung\" autocomplete=\"off\" aria-label=\"Project Description\" maxlength=\"300\" minlength=\"20\" x-on:input=\"submitStatus = $el.value.length &gt;= 20; clearStatus = $el.value.length &gt;= 1\" x-ref=\"description\" hx-post=\"project_clear\" hx-trigger=\"input[readyToClear()]\" hx-target=\"#profile\" hx-swap=\"outerHTML\"></textarea></label></fieldset><div class=\"grid\"><div><button type=\"submit\" value=\"Match\" form=\"project_request\" x-bind:disabled=\"!submitStatus\">Match <span class=\"fa-solid fa-magnifying-glass\"></span></button></div><div><button type=\"button\" x-on:click=\"$refs.description.value = &#39;&#39;\" x-bind:disabled=\"!clearStatus\">Clear <span class=\"fa-solid fa-eraser\"></span></button></div></div></form></div></div><div class=\"col-lg-8\"><div class=\"box\"><div id=\"project_match\"></div></div></div></div><div id=\"externalProfile\"></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

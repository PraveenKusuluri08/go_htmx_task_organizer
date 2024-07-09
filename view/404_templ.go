// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package view

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func PageNotfound(message string) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"text-center\"><img src=\"https://static.vecteezy.com/system/resources/previews/022/756/256/original/delivery-guy-taking-break-bw-empty-state-illustration-editable-404-not-found-page-for-ux-ui-design-isolated-flat-monochromatic-character-on-white-error-flash-message-for-website-app-vector.jpg\" alt=\"404 Not Found Image\" class=\"mx-auto mb-8 w-1/2 h-auto transition-all filter grayscale hover:grayscale-0 rounded-lg hover:shadow-xl dark:shadow-gray-800 bg-cover bg-fixed\"><p class=\"text-2xl font-semibold text-gray-700 mb-4\">Oops! The page you are looking for cannot be found.</p><p class=\"text-lg font-bold text-blue-500\">Please <a href=\"/login\" class=\"hover:underline\" hx-boost=\"true\">Login</a> or <a href=\"/signup\" class=\"hover:underline\" hx-boost=\"true\">Signup</a> to access this page.</p></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
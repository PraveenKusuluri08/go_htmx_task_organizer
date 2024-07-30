# Task Organizer 

This is task organizer application used to manage the tasks for the invidual users and make the priority of the tasks, notify the tasks based whne the due date is done.

## Tech Stack Involved

**GO** I used go for this project <br/>
**Gin** A web framework I used for this project <br/>
**GORM** An ORM library for to intract with the postgresql database <br/>
**HTMX** Frontend library used along with templ for to manage the UI <br/>
**TailwindCSS** A CSS framework used along with the htmx for to desing user interfaces <br/>
**Docker** Making the application containerized.
**AWS** Deploy docker containers to cloud service Planning to use __AWS ECS, EC2, RDS, S3, Cloud Watch__ and __Github Actions__, 

### Steps to install templ 

```console
go install github.com/a-h/templ/cmd/templ@latest
```

Create a makefile which helps to automate the tasks.

Create a go project with this command

```console
go mod init github.com/<name_what_ever_you_like>
```

Follow the folder structure of the project similar to my folder structre, It's not standard but it is easy to follow and understand how it works.


Use the template renderer in the templ to render the templ file here is the code for that.

> Here is the code, I use gin framework
```go
func TemplateRenderer(c *gin.Context, status int, template templ.Component) error {
	c.Status(status)
	return template.Render(c.Request.Context(), c.Writer)
}
```

To generate the go code for the templ need to use the this command.

```console
templ generate
```

To use tailwindcss need to install the tailwind into for that need to use these commands.

```
npm init
```

Which creates a node package manager and install the tailwindcss into it.

```console
npm install tailwindcss -D
```

Configure the tailwindcss environment

```console
npx tailwindcss init
```

Modify the tailwind confile file with this code

```js
/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./templates/*.templ","./templates/components/*.templ"],
  theme: {
    extend: {},
  },
  plugins: [],
}

```

Create a css file and add these directives to it

```css
@tailwind base;
@tailwind components;
@tailwind utilities;

```

To generate the go code for the templ files use this command.

```console
templ generate
```

That's it all good to go 🚀.
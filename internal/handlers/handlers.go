package handlers

import (
	"net/http"

	"github.com/StevenDStanton/htmx-example/internal/forms"
	"github.com/StevenDStanton/htmx-example/internal/models"
	"github.com/StevenDStanton/htmx-example/internal/templates"
)

var data = models.NewData()

func GetBase(w http.ResponseWriter, r *http.Request) {
	templ := templates.NewTemplates("views/*.html")
	pageData := forms.NewPageData(data, forms.NewFormData())
	if err := templ.Render(w, "index", pageData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CreateContact(w http.ResponseWriter, r *http.Request) {
	templ := templates.NewTemplates("views/*.html")
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}
	username := r.FormValue("user_name")
	email := r.FormValue("user_email")
	emailExists := data.EmailExists(email)
	formData := forms.NewFormData()
	if emailExists {
		formData.Values["name"] = username
		formData.Values["email"] = email
		formData.Errors["email"] = "Email address already exists"
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err := templ.Render(w, "form", formData); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	contact := models.NewContact(username, email)
	data.Contacts = append(data.Contacts, contact)

	if err := templ.Render(w, "form", formData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := templ.Render(w, "oob-contactList", contact); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

package main

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"html/template"
	"net/http"
)

const webServerPort = 9000

const view = `<html>
    <head>
        <title>Template</title>
    </head>
    <body>
        <form method="post" action="/login">
            <div>
                <label>username</label>
                <input type="text" name="username" required/>
            </div>
            <div>
                <label>password</label>
                <input type="password" name="password" required/>
            </div>
            <button type="submit">Login</button>
        </form>
    </body>
</html>`

const (
	ldapServer   = "ldap.forumsys.com"
	ldapPort     = 389
	ldapBindDN   = "cn=read-only-admin,dc=example,dc=com"
	ldapPassword = "password"
	ldapSearchDN = "dc=example,dc=com"
)

type UserLDAPData struct {
	ID       string
	Email    string
	Name     string
	FullName string
}

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		var tmpl = template.Must(template.New("main-template").Parse(view))
		if err := tmpl.Execute(rw, nil); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/login", func(rw http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		username := r.PostFormValue("username")
		password := r.PostFormValue("password")

		//authenticate ldap
		ok, data, err := AuthUsingLDAP(username, password)

		if !ok {
			http.Error(rw, "Invalid username or password", http.StatusUnauthorized)
			return
		}

		if err != nil {
			http.Error(rw, err.Error(), http.StatusUnauthorized)
			return
		}

		message := fmt.Sprintf("welcome %s", data.FullName)
		rw.Write([]byte(message))
	})

	portString := fmt.Sprintf(":%d", webServerPort)
	fmt.Println("server started at", portString)

	err := http.ListenAndServe(portString, nil)

	if err != nil {
		fmt.Printf("error : %v", err.Error())
		return
	}
}

func AuthUsingLDAP(username, password string) (bool, *UserLDAPData, error) {
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", ldapServer, ldapPort))
	if err != nil {
		return false, nil, err
	}
	defer l.Close()

	err = l.Bind(ldapBindDN, ldapPassword)
	if err != nil {
		return false, nil, err
	}

	searchRequest := ldap.NewSearchRequest(
		ldapSearchDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		fmt.Sprintf("(&(objectClass=organizationalPerson)(uid=%s))", username),
		[]string{"dn", "cn", "sn", "mail"},
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		return false, nil, err
	}

	if len(sr.Entries) == 0 {
		return false, nil, fmt.Errorf("user not found")
	}

	entry := sr.Entries[0]

	err = l.Bind(entry.DN, password)
	if err != nil {
		return false, nil, err
	}

	data := new(UserLDAPData)
	data.ID = username

	for _, attribute := range entry.Attributes {
		switch attribute.Name {
		case "sn":
			data.Name = attribute.Values[0]
		case "mail":
			data.Email = attribute.Values[0]
		case "cn":
			data.FullName = attribute.Values[0]
		}
	}

	return false, data, nil
}

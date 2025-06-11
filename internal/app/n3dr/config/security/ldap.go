package security

import (
	"fmt"
	"regexp"
	"time"

	"n3dr/internal/app/n3dr/goswagger/client/security_management_l_d_a_p"
	"n3dr/internal/app/n3dr/goswagger/models"

	log "github.com/sirupsen/logrus"
)

type LDAPParams struct {
	Security *Security
	models.CreateLdapServerXo
}

func (l *LDAPParams) LDAP() error {
	log.Info("configuring LDAP")

	client, err := l.Security.Nexus3.Client()
	if err != nil {
		return err
	}

	createLdapServerParams := security_management_l_d_a_p.CreateLdapServerParams{Body: &l.CreateLdapServerXo}
	createLdapServerParams.WithTimeout(time.Second * 30)

	if _, err := client.SecurityManagementldap.CreateLdapServer(&createLdapServerParams); err != nil {
		fmt.Println(err)
		ldapCreated, errRegex := regexp.MatchString("status 400", err.Error())
		if errRegex != nil {
			return err
		}
		if ldapCreated {
			log.Infof("ldap: '%s' has already been created", *l.CreateLdapServerXo.Name)
			return nil
		}

		return fmt.Errorf("could not create ldap: '%v', err: '%w'", *l.CreateLdapServerXo.Name, err)
	}
	log.Infof("created the following repository: '%v'", *l.CreateLdapServerXo.Name)

	return nil
}

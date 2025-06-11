package security

import (
	"fmt"
	"time"

	"n3dr/internal/app/n3dr/connection"
	"n3dr/internal/app/n3dr/goswagger/client/security_management_anonymous_access"
	"n3dr/internal/app/n3dr/goswagger/models"

	log "github.com/sirupsen/logrus"
)

type Security struct {
	connection.Nexus3
}

func (s *Security) Anonymous(enabled bool) error {
	aasx := models.AnonymousAccessSettingsXO{Enabled: enabled}

	log.Info("changing anonymous access")

	client, err := s.Nexus3.Client()
	if err != nil {
		return err
	}

	anonymousAccess := security_management_anonymous_access.UpdateParams{Body: &aasx}
	anonymousAccess.WithTimeout(time.Second * 30)

	resp, err := client.SecurityManagementAnonymousAccess.Update(&anonymousAccess)
	if err != nil {
		return fmt.Errorf("could not change anonymous access mode: '%w'", err)
	}
	if resp.Payload.Enabled {
		log.Info("anonymous access enabled")
	} else {
		log.Info("anonymous access disabled")
	}

	return nil
}

package cli

import (
	"sync"

	"n3dr/internal/app/n3dr/artifactsv2"
	"n3dr/internal/app/n3dr/artifactsv2/upload"
	"n3dr/internal/app/n3dr/connection"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	otherDockerHosts, otherNexus3URLs, otherNexus3Users, otherNexus3Passwords []string
	otherDockerPorts                                                          []int32
	otherDockerSecurePorts                                                    []bool
)

// syncCmd represents the sync command.
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "sync",
	Long:  `sync`,
	Run: func(cmd *cobra.Command, args []string) {
		numberOfURLs := len(otherNexus3URLs)
		numberOfPasswords := len(otherNexus3Passwords)
		numberOfUsers := len(otherNexus3Users)
		if numberOfURLs != numberOfPasswords || numberOfURLs != numberOfUsers {
			log.Fatal("incorrect number of elements. Ensure that the number of elements is identical")
		}

		n := connection.Nexus3{
			DockerHost:       dockerHost,
			DockerPort:       dockerPort,
			DockerPortSecure: dockerPortSecure,
			DownloadDirName:  downloadDirName,
			FQDN:             n3drURL,
			HTTPS:            &https,
			Pass:             n3drPass,
			User:             n3drUser,
		}
		a := artifactsv2.Nexus3{Nexus3: &n}
		if err := a.Backup(); err != nil {
			log.Fatal(err)
		}

		var wg sync.WaitGroup
		for i, otherNexus3User := range otherNexus3Users {
			n.User = otherNexus3User
			n.Pass = otherNexus3Passwords[i]
			n.FQDN = otherNexus3URLs[i]
			n.DockerHost = otherDockerHosts[i]
			n.DockerPort = otherDockerPorts[i]
			n.DockerPortSecure = otherDockerSecurePorts[i]

			wg.Add(1)
			go func(nPreventDataRace connection.Nexus3) {
				defer wg.Done()

				u := upload.Nexus3{Nexus3: &nPreventDataRace}
				if err := u.Upload(); err != nil {
					panic(err)
				}
			}(n)
		}
		wg.Wait()
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)

	syncCmd.PersistentFlags().StringSliceVarP(&otherNexus3URLs, "otherNexus3URLs", "", nil, "specify the other Nexus3 URLs in a comma separated list")
	if err := syncCmd.MarkPersistentFlagRequired("otherNexus3URLs"); err != nil {
		log.Fatal(err)
	}
	syncCmd.PersistentFlags().StringSliceVarP(&otherNexus3Users, "otherNexus3Users", "", nil, "specify the other Nexus3 users in a comma separated list")
	if err := syncCmd.MarkPersistentFlagRequired("otherNexus3Users"); err != nil {
		log.Fatal(err)
	}
	syncCmd.PersistentFlags().StringSliceVarP(&otherNexus3Passwords, "otherNexus3Passwords", "", nil, "specify the other Nexus3 passwords in a comma separated list")
	if err := syncCmd.MarkPersistentFlagRequired("otherNexus3Passwords"); err != nil {
		log.Fatal(err)
	}

	syncCmd.PersistentFlags().Int32SliceVarP(&otherDockerPorts, "otherDockerPorts", "", nil, "specify the otherDockerPorts in a comma separated list")
	if err := syncCmd.MarkPersistentFlagRequired("otherDockerPorts"); err != nil {
		log.Fatal(err)
	}
	syncCmd.PersistentFlags().StringSliceVarP(&otherDockerHosts, "otherDockerHosts", "", nil, "specify the otherDockerHosts in a comma separated list")
	if err := syncCmd.MarkPersistentFlagRequired("otherDockerHosts"); err != nil {
		log.Fatal(err)
	}
	syncCmd.PersistentFlags().BoolSliceVarP(&otherDockerSecurePorts, "otherDockerSecurePorts", "", nil, "specify the otherDockerSecurePorts in a comma separated list")
	if err := syncCmd.MarkPersistentFlagRequired("otherDockerSecurePorts"); err != nil {
		log.Fatal(err)
	}
}

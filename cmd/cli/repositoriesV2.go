package cli

import (
	"n3dr/internal/app/n3dr/artifactsv2"
	"n3dr/internal/app/n3dr/artifactsv2/count"
	"n3dr/internal/app/n3dr/artifactsv2/name"
	"n3dr/internal/app/n3dr/artifactsv2/upload"
	"n3dr/internal/app/n3dr/connection"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var countV2, namesV2, backupV2, uploadV2, withoutWaitGroups, withoutWaitGroupArtifacts, withoutWaitGroupRepositories bool

// repositoriesCmd represents the repositories command.
var repositoriesV2Cmd = &cobra.Command{
	Use:   "repositoriesV2",
	Short: "Return repository names, number of repositories or backup/upload artifacts.",
	Long: `Count the number of repositories, count the total, return the names or
download/upload artifacts from all repositories.

Examples:
  # Return the number of repositories without logging:
  n3dr repositoriesV2 --count --logLevel=none

  # Return the repository names:
  n3dr repositoriesV2 --names

  # Backup a single repository:
  n3dr repositoriesV2 --backup --n3drRepo some-repo --directory-prefix /tmp/some-dir

  # Backup all artifacts:
  n3dr repositoriesV2 --backup --directory-prefix /tmp/some-dir

  # Backup all artifacts without waitGroupRepositories:
  n3dr repositoriesV2 --backup --directory-prefix /tmp/some-dir --withoutWaitGroupRepositories

  # Backup all artifacts without waitGroupArtifacts:
  n3dr repositoriesV2 --backup --directory-prefix /tmp/some-dir --withoutWaitGroupArtifacts

  # Backup all artifacts without waitGroups:
  n3dr repositoriesV2 --backup --directory-prefix /tmp/some-dir --withoutWaitGroups

  # Backup all artifacts, set log level to trace and write it to a file and syslog:
  n3dr repositoriesV2 --backup --directory-prefix /tmp/some-dir --logFile some-file.log --logLevel trace --syslog

  # Backup all artifacts that reside in a Nexus3 server in a certain dir and store these in a zip file:
  n3dr repositoriesV2 --backup --directory-prefix /tmp/some-dir --directory-prefix-zip /tmp/some-dir/some-zip --zip

  # Backup all artifacts including docker images:
  n3dr repositoriesV2 --backup -u some-user -p some-pass -n localhost:9000 --https=false --directory-prefix /tmp/some-dir --dockerPort 9001 --dockerHost http://localhost

  # Upload all artifacts:
  n3dr repositoriesV2 --upload --directory-prefix /tmp/some-dir

  # Upload all artifacts and skip errors:
  n3dr repositoriesV2 --upload --directory-prefix /tmp/some-dir --skipErrors

  # Upload all artifacts without waitGroups:
  n3dr repositoriesV2 --upload --directory-prefix /tmp/some-dir --withoutWaitGroups

  # Upload artifacts, print errors on stderr and write them to syslog:
  n3dr repositoriesV2 --upload -u some-user -p some-pass -n localhost:9000 --https=false --directory-prefix /tmp/some-dir --logLevel=none --logFile
`,
	Run: func(cmd *cobra.Command, args []string) {
		if !(namesV2 || countV2 || backupV2 || uploadV2) {
			if err := cmd.Help(); err != nil {
				log.Fatal(err)
			}
			log.Fatal("One of the required flags \"names\", \"count\" or \"backup\" not set")
		}
		n := connection.Nexus3{
			AwsBucket:                    awsBucket,
			AwsID:                        awsID,
			AwsRegion:                    awsRegion,
			AwsSecret:                    awsSecret,
			BasePathPrefix:               basePathPrefix,
			DockerHost:                   dockerHost,
			DockerPort:                   dockerPort,
			DockerPortSecure:             dockerPortSecure,
			DownloadDirName:              downloadDirName,
			DownloadDirNameZip:           downloadDirNameZip,
			FQDN:                         n3drURL,
			HTTPS:                        &https,
			Pass:                         n3drPass,
			RepoName:                     n3drRepo,
			SkipErrors:                   skipErrors,
			User:                         n3drUser,
			WithoutWaitGroupArtifacts:    withoutWaitGroupArtifacts,
			WithoutWaitGroupRepositories: withoutWaitGroupRepositories,
			WithoutWaitGroups:            withoutWaitGroups,
			ZIP:                          zip,
		}
		a := artifactsv2.Nexus3{Nexus3: &n}
		c := count.Nexus3{Nexus3: &n, CsvFile: csv}
		nn := name.Nexus3{Nexus3: &n}

		if namesV2 {
			if err := nn.Repositories(); err != nil {
				log.Fatal(err)
			}
		}

		if countV2 {
			if err := c.Repositories(); err != nil {
				log.Fatal(err)
			}
		}

		if backupV2 {
			if n.RepoName != "" {
				if err := a.SingleRepoBackup(); err != nil {
					log.Fatal(err)
				}
			} else {
				if err := a.Backup(); err != nil {
					log.Fatal(err)
				}
			}
		}

		if uploadV2 {
			u := upload.Nexus3{Nexus3: &n}
			if err := u.Upload(); err != nil {
				log.Fatal(err)
			}
		}
	},
	Version: rootCmd.Version,
}

func init() {
	repositoriesV2Cmd.Flags().BoolVarP(&namesV2, "names", "", false, "print all repository names")
	repositoriesV2Cmd.Flags().BoolVarP(&countV2, "count", "", false, "count the number of repositories")
	repositoriesV2Cmd.Flags().BoolVarP(&backupV2, "backup", "", false, "backup artifacts from all repositories")
	repositoriesV2Cmd.Flags().BoolVarP(&uploadV2, "upload", "", false, "upload artifacts from all repositories")
	repositoriesV2Cmd.Flags().StringVarP(&regex, "regex", "x", ".*", "only download artifacts that match a regular expression, e.g. 'some/group42'")
	repositoriesV2Cmd.Flags().StringVar(&n3drRepo, "n3drRepo", "", "backup a single nexus3 repository")
	repositoriesV2Cmd.PersistentFlags().BoolVar(&withoutWaitGroups, "withoutWaitGroups", false, "disable the use of waitGroups")
	repositoriesV2Cmd.PersistentFlags().BoolVar(&withoutWaitGroupArtifacts, "withoutWaitGroupArtifacts", false, "disable the use of waitGroupArtifacts")
	repositoriesV2Cmd.PersistentFlags().BoolVar(&withoutWaitGroupRepositories, "withoutWaitGroupRepositories", false, "disable the use of waitGroupRepositories")

	rootCmd.AddCommand(repositoriesV2Cmd)
}

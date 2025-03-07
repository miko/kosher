package cmd

import (
	"log"
	"os"

	"github.com/cbush06/kosher/integrations"

	"github.com/cbush06/kosher/config"
	"github.com/cbush06/kosher/fs"
	"github.com/spf13/cobra"
)

type jiraCommand struct {
	name        string
	command     *cobra.Command
	useDefaults bool
}

func buildJiraCommand() *jiraCommand {
	cmdJira := &jiraCommand{
		name: "jira",
	}

	cmdJira.command = &cobra.Command{
		Use:   "jira",
		Short: "sends results to a Jira system creating tickets for each failed scenario",
		Long:  `jira creates a new Jira ticket for each failed scenario. The fields of the ticket (e.g. type, labels, summary, description, etc.) may be customized via the settings.json file.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			var (
				err        error
				fileSystem *fs.Fs
			)

			// determine where the executable was called from
			workingDir, _ := os.Getwd()
			if fileSystem, err = fs.NewFs(workingDir); err != nil {
				log.Fatal(err)
			}

			// build the settings file based on the working directory
			settings := config.NewSettings(fileSystem)
			settings.Settings.BindPFlag("useDefaults", cmd.Flags().Lookup("default"))

			if err := integrations.SendTo(integrations.Jira, settings); err != nil {
				log.Fatalln(err)
			}

			return nil
		},
	}

	cmdJira.command.Flags().BoolVarP(&cmdJira.useDefaults, "default", "d", false, "If true, uses default values specified in settings.json file.")

	return cmdJira
}

func (s *jiraCommand) registerWith(cmd *cobra.Command) {
	cmd.AddCommand(s.command)
}

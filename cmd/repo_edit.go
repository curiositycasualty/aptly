package cmd

import (
	"fmt"
	"github.com/gonuts/commander"
	"github.com/gonuts/flag"
)

func aptlyRepoEdit(cmd *commander.Command, args []string) error {
	var err error
	if len(args) != 1 {
		cmd.Usage()
		return err
	}

	repo, err := context.collectionFactory.LocalRepoCollection().ByName(args[0])
	if err != nil {
		return fmt.Errorf("unable to edit: %s", err)
	}

	err = context.collectionFactory.LocalRepoCollection().LoadComplete(repo)
	if err != nil {
		return fmt.Errorf("unable to edit: %s", err)
	}

	if cmd.Flag.Lookup("comment").Value.String() != "" {
		repo.Comment = cmd.Flag.Lookup("comment").Value.String()
	}

	if cmd.Flag.Lookup("distribution").Value.String() != "" {
		repo.DefaultDistribution = cmd.Flag.Lookup("distribution").Value.String()
	}

	if cmd.Flag.Lookup("component").Value.String() != "" {
		repo.DefaultComponent = cmd.Flag.Lookup("component").Value.String()
	}

	err = context.collectionFactory.LocalRepoCollection().Update(repo)
	if err != nil {
		return fmt.Errorf("unable to edit: %s", err)
	}

	fmt.Printf("Local repo %s successfully updated.\n", repo)
	return err
}

func makeCmdRepoEdit() *commander.Command {
	cmd := &commander.Command{
		Run:       aptlyRepoEdit,
		UsageLine: "edit <name>",
		Short:     "edit properties of local repository",
		Long: `
Command edit allows to change metadata of local repository:
comment, default distribution and component.

Example:

  $ aptly repo edit -distribution=wheezy testing
`,
		Flag: *flag.NewFlagSet("aptly-repo-edit", flag.ExitOnError),
	}

	cmd.Flag.String("comment", "", "any text that would be used to described local repository")
	cmd.Flag.String("distribution", "", "default distribution when publishing")
	cmd.Flag.String("component", "", "default component when publishing")

	return cmd
}

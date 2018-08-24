package vagrant

import (
	"testing"
)

func init() {
	successfulOutput["version"] = `
1534905015,,ui,output,Installed Version: 2.1.1
1534905015,,version-installed,2.1.1
1534905015,,ui,output,Latest Version: 2.1.2
1534905015,,version-latest,2.1.2
1534905015,,ui,output, \nTo upgrade to the latest version%!(VAGRANT_COMMA) visit the downloads page and\ndownload and install the latest version of Vagrant from the URL\nbelow:\n\n  https://www.vagrantup.com/downloads.html\n\nIf you're curious what changed in the latest release%!(VAGRANT_COMMA) view the\nCHANGELOG below:\n\n  https://github.com/hashicorp/vagrant/blob/v2.1.2/CHANGELOG.md
`
}

func TestVersionCommand_Run(t *testing.T) {
	client := newMockVagrantClient()
	cmd := client.Version()
	if err := cmd.Run(); err != nil {
		t.Fatalf("Command failed to run: %v", err)
	}
	if cmd.Error != nil {
		t.Fatalf("Command returned error: %v", cmd.Error)
	}
	if cmd.InstalledVersion != "2.1.1" {
		t.Errorf("Expected installed version to be '2.1.1'; got %v", cmd.InstalledVersion)
	}
	if cmd.LatestVersion != "2.1.2" {
		t.Errorf("Expected latest version to be '2.1.2'; got %v", cmd.LatestVersion)
	}
}

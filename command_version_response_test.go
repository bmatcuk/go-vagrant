package vagrant

import (
	"testing"
)

const versionOutput = `
1534905015,,ui,output,Installed Version: 2.1.1
1534905015,,version-installed,2.1.1
1534905015,,ui,output,Latest Version: 2.1.2
1534905015,,version-latest,2.1.2
1534905015,,ui,output, \nTo upgrade to the latest version%!(VAGRANT_COMMA) visit the downloads page and\ndownload and install the latest version of Vagrant from the URL\nbelow:\n\n  https://www.vagrantup.com/downloads.html\n\nIf you're curious what changed in the latest release%!(VAGRANT_COMMA) view the\nCHANGELOG below:\n\n  https://github.com/hashicorp/vagrant/blob/v2.1.2/CHANGELOG.md
`

func TestVersionResponse_handleOutput(t *testing.T) {
	parser := MockOutputParser{}
	data := newVersionResponse()
	parser.Run(versionOutput, &data)

	if data.Error != nil {
		t.Errorf("Successful vagrant status should not have set an error: %v", data.Error)
	}
	if data.InstalledVersion != "2.1.1" {
		t.Errorf("Expected installed version to be '2.1.1'; got %v", data.InstalledVersion)
	}
	if data.LatestVersion != "2.1.2" {
		t.Errorf("Expected latest version to be '2.1.2'; got %v", data.LatestVersion)
	}
}

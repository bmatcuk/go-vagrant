package vagrant

import (
	"testing"
)

func init() {
	successfulOutput["box list"] = `
1630581222,,ui,info,andreiborisov/macos-bigsur-intel (parallels%!(VAGRANT_COMMA) 1.3.0)
1630581222,,box-name,andreiborisov/macos-bigsur-intel
1630581222,,box-provider,parallels
1630581222,,box-version,1.3.0
1630581222,,ui,info,andreiborisov/macos-catalina     (parallels%!(VAGRANT_COMMA) 1.3.1)
1630581222,,box-name,andreiborisov/macos-catalina
1630581222,,box-provider,parallels
1630581222,,box-version,1.3.1
`
}
func TestBoxListCommand_Run(t *testing.T) {
	client := newMockVagrantClient()
	cmd := client.BoxList()
	cmd.Env = newEnvTestOutputKey("box list")
	if err := cmd.Run(); err != nil {
		t.Fatalf("Command failed to run: %v", err)
	}
	if cmd.Error != nil {
		t.Fatalf("Command returned error: %v", cmd.Error)
	}
	if len(cmd.Boxes) != 2 {
		t.Errorf("Expected 2 boxes; got %v", len(cmd.Boxes))
	}
	if cmd.Boxes[0].Version != "1.3.0" {
		t.Errorf("Expected box andreiborisov/macos-bigsur-intel version to be '1.3.0'; got %v", cmd.Boxes[0].Version)
	}
	if cmd.Boxes[0].Name != "andreiborisov/macos-bigsur-intel" {
		t.Errorf("Expected box to have name andreiborisov/macos-bigsur-intel; got %v", cmd.Boxes[0].Name)
	}
	if cmd.Boxes[0].Provider != "parallels" {
		t.Errorf("Expected box andreiborisov/macos-bigsur-intel to have parallels provider; got %v", cmd.Boxes[0].Provider)
	}
}

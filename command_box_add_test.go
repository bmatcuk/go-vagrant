package vagrant

import "testing"

func init() {
	successfulOutput["box add"] = `
1631018118,box,ui,output,==> box: Loading metadata for box 'ubuntu/trusty64'
1631018118,box,ui,detail,    box: URL: https://vagrantcloud.com/ubuntu/trusty64
1631018119,box,ui,output,==> box: Adding box 'ubuntu/trusty64' (v20190514.0.0) for provider: virtualbox
1631018119,box,ui,detail,    box: Downloading: https://vagrantcloud.com/ubuntu/boxes/trusty64/versions/20190514.0.0/providers/virtualbox.box
1631018119,,ui,detail,Progress: 0% (Rate: 0*/s%!(VAGRANT_COMMA) Estimated time remaining: --:--:--)
1631018120,,ui,detail,Progress: 100% (Rate: 215/s%!(VAGRANT_COMMA) Estimated time remaining: --:--:--)
1631018121,,ui,detail,Download redirected to host: cloud-images.ubuntu.com
1631018121,,ui,detail,Progress: 100% (Rate: 121/s%!(VAGRANT_COMMA) Estimated time remaining: --:--:--)
1631018121,,ui,detail,Progress: 0% (Rate: 0*/s%!(VAGRANT_COMMA) Estimated time remaining: --:--:--)
1631018123,,ui,detail,Progress: 1% (Rate: 6429k/s%!(VAGRANT_COMMA) Estimated time remaining: 0:03:05)
1631018124,,ui,detail,Progress: 8% (Rate: 17.4M/s%!(VAGRANT_COMMA) Estimated time remaining: 0:00:42)
1631018125,,ui,detail,Progress: 15% (Rate: 21.2M/s%!(VAGRANT_COMMA) Estimated time remaining: 0:00:27)
1631018126,,ui,detail,Progress: 21% (Rate: 23.5M/s%!(VAGRANT_COMMA) Estimated time remaining: 0:00:21)
1631018127,,ui,detail,Progress: 28% (Rate: 24.7M/s%!(VAGRANT_COMMA) Estimated time remaining: 0:00:17)
1631018128,,ui,detail,Progress: 35% (Rate: 29.2M/s%!(VAGRANT_COMMA) Estimated time remaining: 0:00:14)
1631018129,,ui,detail,Progress: 42% (Rate: 29.2M/s%!(VAGRANT_COMMA) Estimated time remaining: 0:00:12)
1631018130,,ui,detail,Progress: 49% (Rate: 29.6M/s%!(VAGRANT_COMMA) Estimated time remaining: 0:00:10)
1631018131,,ui,detail,Progress: 56% (Rate: 29.3M/s%!(VAGRANT_COMMA) Estimated time remaining: 0:00:09)
1631018132,,ui,detail,Progress: 63% (Rate: 29.5M/s%!(VAGRANT_COMMA) Estimated time remaining: 0:00:07)
1631018133,,ui,detail,Progress: 67% (Rate: 27.6M/s%!(VAGRANT_COMMA) Estimated time remaining: 0:00:06)
1631018134,,ui,detail,Progress: 74% (Rate: 27.4M/s%!(VAGRANT_COMMA) Estimated time remaining: 0:00:05)
1631018135,,ui,detail,Progress: 81% (Rate: 27.4M/s%!(VAGRANT_COMMA) Estimated time remaining: 0:00:04)
1631018136,,ui,detail,Progress: 88% (Rate: 27.2M/s%!(VAGRANT_COMMA) Estimated time remaining: 0:00:02)
1631018137,,ui,detail,Progress: 94% (Rate: 27.1M/s%!(VAGRANT_COMMA) Estimated time remaining: 0:00:01)
1631018138,box,ui,success,==> box: Successfully added box 'ubuntu/trusty64' (v20190514.0.0) for 'virtualbox'!
`
}
func TestBoxAddCommand_buildArguments(t *testing.T) {
	client := newMockVagrantClient()

	t.Run("error-on-missing-location", func(t *testing.T) {
		cmd := client.BoxAdd("")
		_, err := cmd.buildArguments()
		assertErr(t, err)
	})
	t.Run("default", func(t *testing.T) {
		cmd := client.BoxAdd("ubuntu/trusty64")
		args, err := cmd.buildArguments()
		assertNoErr(t, err)
		assertArguments(t, args, "add", "ubuntu/trusty64")
	})
	t.Run("checksum-type", func(t *testing.T) {
		cmd := client.BoxAdd("ubuntu/trusty64")
		cmd.Checksum = "test"
		cmd.CheckSumType = SHA1
		args, err := cmd.buildArguments()
		assertNoErr(t, err)
		assertArguments(t, args, "add", "--checksum", "test", "--checksum-type", "sha1", "ubuntu/trusty64")
	})
	t.Run("checksum", func(t *testing.T) {
		cmd := client.BoxAdd("ubuntu/trusty64")
		cmd.Checksum = "test"
		args, err := cmd.buildArguments()
		assertNoErr(t, err)
		assertArguments(t, args, "add", "--checksum", "test", "ubuntu/trusty64")
	})
	t.Run("error-on-checksum-type-without-checksum", func(t *testing.T) {
		cmd := client.BoxAdd("ubuntu/trusty64")
		cmd.CheckSumType = SHA1
		_, err := cmd.buildArguments()
		assertErr(t, err)
	})
	t.Run("name", func(t *testing.T) {
		cmd := client.BoxAdd("ubuntu/trusty64")
		cmd.Name = "myubuntu"
		args, err := cmd.buildArguments()
		assertNoErr(t, err)
		assertArguments(t, args, "add", "--name", "myubuntu", "ubuntu/trusty64")
	})
	t.Run("force", func(t *testing.T) {
		cmd := client.BoxAdd("ubuntu/trusty64")
		cmd.Force = true
		args, err := cmd.buildArguments()
		assertNoErr(t, err)
		assertArguments(t, args, "add", "-f", "ubuntu/trusty64")
	})
	t.Run("clean", func(t *testing.T) {
		cmd := client.BoxAdd("ubuntu/trusty64")
		cmd.Clean = true
		args, err := cmd.buildArguments()
		assertNoErr(t, err)
		assertArguments(t, args, "add", "-c", "ubuntu/trusty64")
	})
}
func TestBoxAddCommand_Run(t *testing.T) {
	client := newMockVagrantClient()
	cmd := client.BoxAdd("ubuntu/trusty64")
	cmd.Env = newEnvTestOutputKey("box add")
	if err := cmd.Run(); err != nil {
		t.Fatalf("Command failed to run: %v", err)
	}
	if cmd.Error != nil {
		t.Fatalf("Command returned error: %v", cmd.Error)
	}
}

package version

import "testing"

func TestDetectInstallMethod(t *testing.T) {
	tests := []struct {
		name     string
		execPath string
		want     InstallMethod
	}{
		{name: "homebrew arm", execPath: "/opt/homebrew/Cellar/skillshare/0.26.3/bin/skillshare", want: InstallBrew},
		{name: "linuxbrew", execPath: "/home/linuxbrew/.linuxbrew/Cellar/skillshare/0.26.3/bin/skillshare", want: InstallBrew},
		{name: "nix store", execPath: "/nix/store/abc123-skillshare-0.26.3/bin/skillshare", want: InstallNix},
		{name: "install script", execPath: "/home/me/.local/bin/skillshare", want: InstallDirect},
		{name: "nix-like name outside the store", execPath: "/home/me/nix/store/skillshare", want: InstallDirect},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DetectInstallMethod(tt.execPath); got != tt.want {
				t.Fatalf("DetectInstallMethod(%q) = %v, want %v", tt.execPath, got, tt.want)
			}
		})
	}
}

func TestUpgradeCommand(t *testing.T) {
	tests := []struct {
		method InstallMethod
		want   string
	}{
		{method: InstallDirect, want: "skillshare upgrade"},
		{method: InstallBrew, want: "brew upgrade skillshare"},
		{method: InstallNix, want: "nix profile upgrade skillshare"},
	}
	for _, tt := range tests {
		if got := tt.method.UpgradeCommand(); got != tt.want {
			t.Fatalf("UpgradeCommand(%v) = %q, want %q", tt.method, got, tt.want)
		}
	}
}

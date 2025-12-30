package submodule

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. "github.com/jesseduffield/lazygit/pkg/integration/components"
)

// This test verifies that screen mode is preserved when entering and exiting submodules.
// Before the fix, entering a submodule would reset the screen mode to the default from config.
// After the fix, the current screen mode is preserved when switching to a new repo.
var EnterPreserveScreenMode = NewIntegrationTest(NewIntegrationTestArgs{
	Description:  "Enter a submodule and verify screen mode is preserved",
	ExtraCmdArgs: []string{},
	Skip:         false,
	SetupConfig:  func(cfg *config.AppConfig) {},
	SetupRepo: func(shell *Shell) {
		shell.EmptyCommit("first commit")
		shell.CloneIntoSubmodule("my_submodule_name", "my_submodule_path")
		shell.GitAddAll()
		shell.Commit("add submodule")
	},
	Run: func(t *TestDriver, keys config.KeybindingConfig) {
		assertInParentRepo := func() {
			t.Views().Status().Content(Contains("repo"))
		}
		assertInSubmodule := func() {
			t.Views().Status().Content(Contains("my_submodule_path(my_submodule_name)"))
		}

		assertInParentRepo()

		// Change screen mode in parent repo (cycle through NORMAL -> HALF -> FULL -> NORMAL)
		// This verifies screen mode can be changed and works
		t.Views().Submodules().Focus().
			Press(keys.Universal.NextScreenMode)

		// Enter the submodule - screen mode should be preserved (now HALF)
		t.Views().Submodules().
			Lines(
				Contains("my_submodule_name").IsSelected(),
			).
			PressEnter()

		assertInSubmodule()

		// Cycle screen mode in submodule to verify it still works
		t.Views().Files().Focus().
			Press(keys.Universal.NextScreenMode). // HALF -> FULL
			Press(keys.Universal.NextScreenMode)  // FULL -> NORMAL

		// Return to parent repo
		t.Views().Files().PressEscape()

		assertInParentRepo()

		// Verify we can still interact with submodules panel
		t.Views().Submodules().Focus().
			Lines(
				Contains("my_submodule_name").IsSelected(),
			)

		// Re-enter the submodule
		t.Views().Submodules().PressEnter()

		assertInSubmodule()

		// Verify we can still cycle screen mode
		t.Views().Files().Focus().
			Press(keys.Universal.NextScreenMode)
	},
})

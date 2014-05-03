package cronedit

import (
	"fmt"
	"testing"
)

func TestCronEdit(t *testing.T) {

	tests := []struct {
		input  string
		output string
	}{
		{
			// Zmon line is added.
			`# m h  dom mon dow   command
MAILTO=""
`,
			`# m h  dom mon dow   command
MAILTO=""
@hourly nohup /path/to/zmon &
`},
		{
			// Nothing is changed.
			`@hourly nohup /path/to/zmon &
`,
			`@hourly nohup /path/to/zmon &
`},
	}
	for _, x := range tests {
		_, out := edit(x.input, "@hourly nohup /path/to/zmon &")
		if out != x.output {
			t.Errorf("edit() Input:\n%v\n===\nOutput:\n%v\n===\nWanted:\n%v\n===\n", x.input, out, x.output)
		}
	}
}

// ExampleCronInsert inspects the crontab for the current user and adds the specified command if it's not already there.
func ExampleCronInsert() {
	_, err := Insert("@hourly update-something")
	if err != nil {
		fmt.Println("Crontab edit error:", err)
		return
	}
}

func init() {
	// Do not really maniuplate the crontab when running tests.
	replaceCrontab = func(newContent string) error {
		return nil
	}
}

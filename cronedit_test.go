package cronedit

import "testing"

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

// DisabledTestCronInsert really updates the current user crontab, so it's disabled.
func DisabledTestCronInsert(t *testing.T) {
	changed, err := Insert("@hourly nohup $HOME/bin/zmon &")
	t.Logf("TestCronInsert: changed %v, err %v", changed, err)
	if err != nil {
		t.Error(err)
	}
}

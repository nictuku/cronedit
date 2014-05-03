cronedit
========

Manipulate the user crontab from Go. Usage:


```
package main

import (
  "fmt"
  "github.com/nictuku/cronedit"
)

func main() {
  changed, err := cronedit.Insert("@hourly update-something")
	if err != nil {
		fmt.Println("Crontab edit error:", err)
		return
	}
	fmt.Println("Crontab changed?", changed)
}
```

Tested on Linux and OSX. Enjoy!

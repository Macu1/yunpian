# yunpian go sdk

## How To Use
``` go
import (
	"github.com/shesuyo/yunpian"
)
```

```
apiKey = "apiKey"
sms := &SMSData{Mobile: "MobieNum", Text: "YourText"}
body, code, err = yunpian.DoRequest(sms, apiKey)
```

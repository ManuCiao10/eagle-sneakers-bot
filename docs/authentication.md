# Authentication

Authentication can be implemented strictly through Discord user connections using [Whop](https://whop.com), a dedicated API, or a combination of both.&#x20;

With Whop, we can obtain the Discord user's information, including their user ID, key, machine HWID, and access a dedicated dashboard to showcase our product or service.

### Client Settings

To ensure a personalized experience for users, they can store their unique settings in the `settings.json` file, which should be located in the root folder of our software. This file acts as a central repository for storing and accessing user-specific preferences and configurations.

{% code overflow="wrap" %}
```go
package settings

type Settings struct {
	AuthKey        string `json:"key"`
	DiscordWebhook string `json:"webhook"`
	TwoCaptchaKey  string `json:"2captcha_key"`
	AnticaptchaKey string `json:"anticaptcha_key"`
	CapmonsterKey  string `json:"capmonster_key"`
	Solver         string `json:"solver"`
	Delay          Delay  `json:"delay"`
	TaskShoutDown  string `json:"task_shutdown"`
}

type Delay struct {
	Retry   string `json:"retry"`
	Timeout string `json:"timeout"`
}
```
{% endcode %}

### Validate a License Key

This validation involves checking the key against expiration date, activation limits, and machine hwid. Preventing unauthorized usage, and ensure a secure and compliant user experience.

{% code overflow="wrap" %}
```go
func validate_license() {
	licence := loading.Data.Settings.Settings.AuthKey

	url := "https://api.whop.com/api/v2/memberships/<licence>/validate_license"

	req, _ := http.NewRequest("POST", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := os.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
```
{% endcode %}




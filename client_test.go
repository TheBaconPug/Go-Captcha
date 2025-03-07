package gocaptcha

import "testing"

func TestClient(t *testing.T) {
	client := NewClient("Your-key", "https://2captcha.com/in.php", "https://2captcha.com/res.php")

	captcha := NewRecaptchaV2()
	captcha.SetSiteKey("6LeIxboZAAAAAFQy7d8GPzgRZu2bV0GwKS8ue_cH")
	captcha.SetUrl("https://2captcha.com/demo/recaptcha-v2")
	captcha.SetAction("verify")

	_, err := client.Solve(captcha)
	if err != nil {
		t.Errorf("error: %s", err.Error())
	}
}

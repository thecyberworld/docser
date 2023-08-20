package patterns

import (
	"regexp"
)

// DefinePatternInfo represents a regex pattern along with its description
type DefinePatternInfo struct {
	Pattern     *regexp.Regexp
	Description string
}

// RegexPatterns contains all the defined regex patterns
var RegexPatterns = []DefinePatternInfo{
	{regexp.MustCompile(`cloudinary://.*`), "Cloudinary"},
	{regexp.MustCompile(`.*firebaseio\.com`), "Firebase URL"},
	{regexp.MustCompile(`(xox[p|b|o|a]-[0-9]{12}-[0-9]{12}-[0-9]{12}-[a-z0-9]{32})`), "Slack Token"},
	{regexp.MustCompile(`-----BEGIN RSA PRIVATE KEY-----`), "RSA private key"},
	{regexp.MustCompile(`-----BEGIN DSA PRIVATE KEY-----`), "SSH (DSA) private key"},
	{regexp.MustCompile(`-----BEGIN EC PRIVATE KEY-----`), "SSH (EC) private key"},
	{regexp.MustCompile(`-----BEGIN PGP PRIVATE KEY BLOCK-----`), "PGP private key block"},
	{regexp.MustCompile(`AKIA[0-9A-Z]{16}`), "Amazon AWS Access Key ID"},
	{regexp.MustCompile(`amzn\.mws\.[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`), "Amazon MWS Auth Token"},
	{regexp.MustCompile(`AKIA[0-9A-Z]{16}`), "AWS API Key"},
	{regexp.MustCompile(`EAACEdEose0cBA[0-9A-Za-z]+`), "Facebook Access Token"},
	{regexp.MustCompile(`[f|F][a|A][c|C][e|E][b|B][o|O][o|O][k|K].*['|\"][0-9a-f]{32}['|\"]`), "Facebook OAuth"},
	{regexp.MustCompile(`[g|G][i|I][t|T][h|H][u|U][b|B].*['|\"][0-9a-zA-Z]{35,40}['|\"]`), "GitHub"},
	{regexp.MustCompile(`[a|A][p|P][i|I][_]?[k|K][e|E][y|Y].*['|\"][0-9a-zA-Z]{32,45}['|\"]`), "Generic API Key"},
	{regexp.MustCompile(`[s|S][e|E][c|C][r|R][e|E][t|T].*['|\"][0-9a-zA-Z]{32,45}['|\"]`), "Generic Secret"},
	{regexp.MustCompile(`AIza[0-9A-Za-z\\-_]{35}`), "Google API Key"},
	{regexp.MustCompile(`AIza[0-9A-Za-z\\-_]{35}`), "Google Cloud Platform API Key"},
	{regexp.MustCompile(`[0-9]+-[0-9A-Za-z_]{32}\\.apps\\.googleusercontent\\.com`), "Google Cloud Platform OAuth"},
	{regexp.MustCompile(`AIza[0-9A-Za-z\\-_]{35}`), "Google Drive API Key"},
	{regexp.MustCompile(`[0-9]+-[0-9A-Za-z_]{32}\\.apps\\.googleusercontent\\.com`), "Google Drive OAuth"},
	{regexp.MustCompile(`"type": "service_account"`), "Google (GCP) Service-account"},
	{regexp.MustCompile(`AIza[0-9A-Za-z\\-_]{35}`), "Google Gmail API Key"},
	{regexp.MustCompile(`[0-9]+-[0-9A-Za-z_]{32}\\.apps\\.googleusercontent\\.com`), "Google Gmail OAuth"},
	{regexp.MustCompile(`ya29\\.[0-9A-Za-z\\-_]+`), "Google OAuth Access Token"},
	{regexp.MustCompile(`AIza[0-9A-Za-z\\-_]{35}`), "Google YouTube API Key"},
	{regexp.MustCompile(`[0-9]+-[0-9A-Za-z_]{32}\\.apps\\.googleusercontent\\.com`), "Google YouTube OAuth"},
	{regexp.MustCompile(`[h|H][e|E][r|R][o|O][k|K][u|U].*[0-9A-F]{8}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{12}`), "Heroku API Key"},
	{regexp.MustCompile(`[0-9a-f]{32}-us[0-9]{1,2}`), "MailChimp API Key"},
	{regexp.MustCompile(`key-[0-9a-zA-Z]{32}`), "Mailgun API Key"},
	{regexp.MustCompile(`[a-zA-Z]{3,10}://[^/\\s:@]{3,20}:[^/\\s:@]{3,20}@.{1,100}[\"'\\s]`), "Password in URL"},
	{regexp.MustCompile(`access_token\$production\$[0-9a-z]{16}\$[0-9a-f]{32}`), "PayPal Braintree Access Token"},
	{regexp.MustCompile(`sk_live_[0-9a-z]{32}`), "Picatic API Key"},
	{regexp.MustCompile(`https://hooks.slack.com/services/T[a-zA-Z0-9_]{8}/B[a-zA-Z0-9_]{8}/[a-zA-Z0-9_]{24}`), "Slack Webhook"},
	{regexp.MustCompile(`sk_live_[0-9a-zA-Z]{24}`), "Stripe API Key"},
	{regexp.MustCompile(`rk_live_[0-9a-zA-Z]{24}`), "Stripe Restricted API Key"},
	{regexp.MustCompile(`sq0atp-[0-9A-Za-z\\-_]{22}`), "Square Access Token"},
	{regexp.MustCompile(`sq0csp-[0-9A-Za-z\\-_]{43}`), "Square OAuth Secret"},
	{regexp.MustCompile(`SK[0-9a-fA-F]{32}`), "Twilio API Key"},
	{regexp.MustCompile(`[t|T][w|W][i|I][t|T][t|T][e|E][r|R].*[1-9][0-9]+-[0-9a-zA-Z]{40}`), "Twitter Access Token"},
	{regexp.MustCompile(`[t|T][w|W][i|I][t|T][t|T][e|E][r|R].*['|\"][0-9a-zA-Z]{35,44}['|\"]`), "Twitter OAuth"},
}

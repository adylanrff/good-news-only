package twitter

type Twitter interface {
	Tweet() error
	Retweet() error
	Get() error
}

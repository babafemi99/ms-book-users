package oauth

import (
	"bookApi/utils/msErrors"
	"bookApi/utils/oauth/utils"
	"errors"
	"net/http"
)

const (
	publicXHeader  = "X-Public"
	headerClientId = "X-Client-Id"
	headerCallerId = "X-Caller-Id"
)

type oauth struct {
	request *http.Request
}

func (o *oauth) Authenticate() *msErrors.RestErrors {
	if o.request == nil {
		return msErrors.NewBadRequest("Bad request", errors.New("check if request is empty"))
	}
	panic("implement me")
}

func NewOauth(request *http.Request) *oauth {
	return &oauth{request: request}
}

func (o *oauth) isExpired() bool {
	user := o.GetUserDetails()
	if err := user.Valid(); err != nil {
		return false
	}
	return true
}

func (o *oauth) IsPublic() bool {
	if o.request == nil {
		return true
	}
	return o.request.Header.Get(publicXHeader) == "true"
}

func (o *oauth) IsPrivate() bool {
	return !o.IsPublic()
}

func (o *oauth) GetUserDetails() *utils.SignedDetails {
	user, msg := utils.ValidateToken(o.request)
	if msg != "" {
		return nil
	}
	return user

}

package providers

import (
	"context"
	"fmt"
)

type Generic struct {
	BaseProvider
}

func (p *Generic) GetUserInfo(ctx context.Context, code string) (*UserInfo, error) {
	token, err := p.oauthConfig.Exchange(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	res, err := p.oauthConfig.Client(ctx, token).Get(p.UserInfoURL)
	if err != nil {
		return nil, fmt.Errorf("failed to get user info: %+q", err)
	}
	defer res.Body.Close()

	var dest map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&dest); err != nil {
		return nil, err
	}

	mapping := p.BaseProvider.mapping
	sub, ok := dest[mapping.ID]
	if !ok {
		return nil, fmt.Errorf("failed to get id from user info")
	}
	if sub == nil {
		return nil, fmt.Errorf("no (empty/valid) user id found in user info")
	}
	subId := sub.(string)
	if subId == "" {
		return nil, fmt.Errorf("empty external user/sub id given")
	}

	usernameRaw, ok := dest[mapping.Username]
	if !ok {
		return nil, fmt.Errorf("failed to get username from user info")
	}
	if usernameRaw == nil {
		return nil, fmt.Errorf("no username found in user info")
	}

	username := usernameRaw.(string)

	user := &UserInfo{
		ID:       subId,
		Username: username,
	}

	return user, nil
}

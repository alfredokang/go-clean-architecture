package users

func NewResponseUsersMessage(users []OutputDetailedUser, msg string, flag string) OutputUsers {
	return OutputUsers{
		Users: users,
		Msg:   msg,
		Flag:  flag,
	}
}

func NewResponseUserMessage(user OutputDetailedUser, msg string, flag string) OutputUser {
	return OutputUser{
		User: user,
		Msg:  msg,
		Flag: flag,
	}
}

func NewResponseMessage(msg string, flag string) OutputRaw {
	return OutputRaw{
		Msg:  msg,
		Flag: flag,
	}
}

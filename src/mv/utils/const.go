

package utils

const (
	/* FOR REDIS */
	AUTH_USER_EMAIL_SESSION = "auth_session_id"
	AUTH_USER_EMAIL         = "auth_email"
	AUTH_USER_USERID        = "auth_user_id"
	AUTH_USER_FIRSTNAME     = "auth_user_firstname"
	AUTH_USER_LASTNAME      = "auth_user_Lastname"
	AUTH_USER_TYPE          = "auth_user_type"
	AUTH_USER_AVATARNAME    = "auth_user_avatar_name"
	AUTH_USER_BLOCKED       = "auth_user_blocked"

	/* FOR KAFKA ANNOUNCEMENTS */
	AUTH_SIGNUP_KAFKA_TOPIC            = "auth_new_user_signup_event"
	AUTH_LOGIN_KAFKA_TOPIC             = "auth_user_login_complete"
	AUTH_LOGOUT_KAFKA_TOPIC            = "auth_user_logout_complete"
	AUTH_VALIDATE_SUCCESS_TOPIC        = "auth_user_signup_complete_validate_complete"
	AUTH_USER_FORGOT_PWD_NEW_PWD_TOPIC = "auth_user_forgot_password_new_password"
)

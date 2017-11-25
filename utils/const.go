package utils

/* Contants for use in storing and retriving user info from Redis FastMem DB
 */
const (
	/* FOR REDIS */
	AuthUserEmailSession = "auth_session_id"
	AuthUserEmail        = "auth_email"
	AuthUserID           = "auth_user_id"
	AuthUserFirstName    = "auth_user_firstname"
	AuthUserLastName     = "auth_user_Lastname"
	AuthUserType         = "auth_user_type"
	AuthUserAvatarName   = "auth_user_avatar_name"
	AuthIsUserBlocked    = "auth_user_blocked"

	/* FOR KAFKA ANNOUNCEMENTS */
	AuthNewUserSignupEventNotification     = "auth_new_user_signup_event"
	AuthUserLoginCompleteNotification      = "auth_user_login_complete"
	AuthUserLogoutCompleteNotification     = "auth_user_logout_complete"
	AuthNewUsersSignupCompleteNotification = "auth_user_signup_complete_validate_complete"
	AuthUserForgotPasswordNotification     = "auth_user_forgot_password_new_password"
)

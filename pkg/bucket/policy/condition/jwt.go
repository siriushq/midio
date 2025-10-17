package condition

// JWT claims supported substitutions.
// https://www.iana.org/assignments/jwt/jwt.xhtml#claims
const (
	// JWTSub - JWT subject claim substitution.
	JWTSub Key = "jwt:sub"

	// JWTIss issuer claim substitution.
	JWTIss Key = "jwt:iss"

	// JWTAud audience claim substitution.
	JWTAud Key = "jwt:aud"

	// JWTJti JWT unique identifier claim substitution.
	JWTJti Key = "jwt:jti"

	JWTUpn          Key = "jwt:upn"
	JWTName         Key = "jwt:name"
	JWTGroups       Key = "jwt:groups"
	JWTGivenName    Key = "jwt:given_name"
	JWTFamilyName   Key = "jwt:family_name"
	JWTMiddleName   Key = "jwt:middle_name"
	JWTNickName     Key = "jwt:nickname"
	JWTPrefUsername Key = "jwt:preferred_username"
	JWTProfile      Key = "jwt:profile"
	JWTPicture      Key = "jwt:picture"
	JWTWebsite      Key = "jwt:website"
	JWTEmail        Key = "jwt:email"
	JWTGender       Key = "jwt:gender"
	JWTBirthdate    Key = "jwt:birthdate"
	JWTPhoneNumber  Key = "jwt:phone_number"
	JWTAddress      Key = "jwt:address"
	JWTScope        Key = "jwt:scope"
	JWTClientID     Key = "jwt:client_id"
)

// JWTKeys - Supported JWT keys, non-exhaustive list please
// expand as new claims are standardized.
var JWTKeys = []Key{
	JWTSub,
	JWTIss,
	JWTAud,
	JWTJti,
	JWTName,
	JWTUpn,
	JWTGroups,
	JWTGivenName,
	JWTFamilyName,
	JWTMiddleName,
	JWTNickName,
	JWTPrefUsername,
	JWTProfile,
	JWTPicture,
	JWTWebsite,
	JWTEmail,
	JWTGender,
	JWTBirthdate,
	JWTPhoneNumber,
	JWTAddress,
	JWTScope,
	JWTClientID,
}

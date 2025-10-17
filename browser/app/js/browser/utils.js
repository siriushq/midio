export const OPEN_ID_NONCE_KEY = 'openIDKey'

export const buildOpenIDAuthURL = (authEp, authScopes, redirectURI, clientID, nonce) => {
  const params = new URLSearchParams()
  params.set("response_type", "id_token")
  params.set("scope", authScopes.join(" "))
  params.set("client_id", clientID)
  params.set("redirect_uri", redirectURI)
  params.set("nonce", nonce)

  return `${authEp}?${params.toString()}`
}

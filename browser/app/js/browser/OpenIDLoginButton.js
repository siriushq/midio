import React from "react"
import { getRandomString } from "../utils"
import storage from "local-storage-fallback"
import { buildOpenIDAuthURL, OPEN_ID_NONCE_KEY } from './utils'

export class OpenIDLoginButton extends React.Component {
  constructor(props) {
    super(props)
    this.handleClick = this.handleClick.bind(this)
  }

  handleClick(event) {
    event.stopPropagation()
    const { authEp, authScopes, clientId } = this.props

    let redirectURI = window.location.href.split("#")[0]
    if (redirectURI.endsWith('/')) {
      redirectURI += 'openid'
    } else {
      redirectURI += '/openid'
    }

    // Store nonce in localstorage to check again after the redirect
    const nonce = getRandomString(16)
    storage.setItem(OPEN_ID_NONCE_KEY, nonce)

    const authURL = buildOpenIDAuthURL(authEp, authScopes, redirectURI, clientId, nonce)
    window.location = authURL
  }

  render() {
    const { children, className } = this.props
    return (
      <div onClick={this.handleClick} className={className}>
        {children}
      </div>
    )
  }
}

export default OpenIDLoginButton

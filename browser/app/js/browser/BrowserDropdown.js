import React from "react"
import { connect } from "react-redux"
import { Dropdown } from "react-bootstrap"
import * as browserActions from "./actions"
import web from "../web"
import history from "../history"
import AboutModal from "./AboutModal"
import ChangePasswordModal from "./ChangePasswordModal"

export class BrowserDropdown extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
      showAboutModal: false,
      showChangePasswordModal: false
    }
  }
  showAbout(e) {
    e.preventDefault()
    this.setState({
      showAboutModal: true
    })
  }
  hideAbout() {
    this.setState({
      showAboutModal: false
    })
  }
  showChangePassword(e) {
    e.preventDefault()
    this.setState({
      showChangePasswordModal: true
    })
  }
  hideChangePassword() {
    this.setState({
      showChangePasswordModal: false
    })
  }
  componentDidMount() {
    const { fetchServerInfo } = this.props
    fetchServerInfo()
  }
  logout(e) {
    e.preventDefault()
    web.Logout()
    history.replace("/login")
  }
  render() {
    const { serverInfo } = this.props
    return (
      <li>
        <Dropdown pullRight id="top-right-menu">
          <Dropdown.Toggle noCaret>
            <i className="fas fa-bars" />
          </Dropdown.Toggle>
          <Dropdown.Menu className="dropdown-menu-right">
            <li>
              <a href="" onClick={this.showChangePassword.bind(this)}>
                Change Password <i className="fas fa-cog" />
              </a>
              {this.state.showChangePasswordModal && (
                <ChangePasswordModal
                  serverInfo={serverInfo}
                  hideChangePassword={this.hideChangePassword.bind(this)}
                />
              )}
            </li>
            <li>
              <a target="_blank" href="https://docs.min.io/?ref=ob">
                Documentation <i className="fas fa-book" />
              </a>
            </li>
            <li>
              <a target="_blank" href="https://github.com/minio/minio">
                GitHub <i className="fab fa-github" />
              </a>
            </li>
            <li>
              <a target="_blank" href="https://min.io/pricing?ref=ob">
                Get Support <i className="fas fa-question-circle" />
              </a>
            </li>
            <li>
              <a href="" id="show-about" onClick={this.showAbout.bind(this)}>
                About <i className="fas fa-info-circle" />
              </a>
              {this.state.showAboutModal && (
                <AboutModal
                  serverInfo={serverInfo}
                  hideAbout={this.hideAbout.bind(this)}
                />
              )}
            </li>
            <li>
              <a href="" id="logout" onClick={this.logout}>
                Logout <i className="fas fa-sign-out-alt" />
              </a>
            </li>
          </Dropdown.Menu>
        </Dropdown>
      </li>
    )
  }
}

const mapStateToProps = state => {
  return {
    serverInfo: state.browser.serverInfo
  }
}

const mapDispatchToProps = dispatch => {
  return {
    fetchServerInfo: () => dispatch(browserActions.fetchServerInfo())
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(BrowserDropdown)

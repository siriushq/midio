import React from "react"
import { Modal } from "react-bootstrap"
import logo from "../../img/logo.svg"

export const AboutModal = ({ serverInfo, hideAbout }) => {
  const { version, platform, runtime } = serverInfo
  return (
    <Modal
      className="modal-about modal-dark"
      animation={false}
      show={true}
      onHide={hideAbout}
    >
      <button className="close" onClick={hideAbout}>
        <span>×</span>
      </button>
      <div className="ma-inner">
        <div className="mai-item hidden-xs">
          <a href="https://min.io" target="_blank">
            <img className="maii-logo" src={logo} alt="" />
          </a>
        </div>
        <div className="mai-item">
          <ul className="maii-list">
            <li>
              <div>Version</div>
              <small>{version}</small>
            </li>
            <li>
              <div>Platform</div>
              <small>{platform}</small>
            </li>
            <li>
              <div>Runtime</div>
              <small>{runtime}</small>
            </li>
          </ul>
        </div>
      </div>
    </Modal>
  )
}

export default AboutModal

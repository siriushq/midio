import React from "react"
import classNames from "classnames"
import { connect } from "react-redux"
import SideBar from "./SideBar"
import MainContent from "./MainContent"
import AlertContainer from "../alert/AlertContainer"

class Browser extends React.Component {
  render() {
    return (
      <div
        className={classNames({
          "file-explorer": true
        })}
      >
        <SideBar />
        <MainContent />
        <AlertContainer />
      </div>
    )
  }
}

export default connect(state => state)(Browser)

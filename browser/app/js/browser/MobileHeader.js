import React from "react"
import classNames from "classnames"
import { connect } from "react-redux"
import logo from "../../img/logo.svg"
import * as actionsCommon from "./actions"

export const MobileHeader = ({ sidebarOpen, toggleSidebar }) => (
  <header className="fe-header-mobile hidden-lg hidden-md">
    <div
      id="sidebar-toggle"
      className={
        "feh-trigger " +
        classNames({
          "feht-toggled": sidebarOpen
        })
      }
      onClick={e => {
        e.stopPropagation()
        toggleSidebar()
      }}
    >
      <div className="feht-lines">
        <div className="top" />
        <div className="center" />
        <div className="bottom" />
      </div>
    </div>
    <img className="mh-logo" src={logo} alt="" />
  </header>
)

const mapStateToProps = state => {
  return {
    sidebarOpen: state.browser.sidebarOpen
  }
}

const mapDispatchToProps = dispatch => {
  return {
    toggleSidebar: () => dispatch(actionsCommon.toggleSidebar())
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(MobileHeader)

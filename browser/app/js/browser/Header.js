import React from "react"
import ObjectsSearch from "../objects/ObjectsSearch"
import Path from "../objects/Path"
import StorageInfo from "./StorageInfo"
import BrowserDropdown from "./BrowserDropdown"
import web from "../web"
import { minioBrowserPrefix } from "../constants"

export const Header = () => {
  const loggedIn = web.LoggedIn()
  return (
    <header className="fe-header">
      <Path />
      {loggedIn && <StorageInfo />}
      {loggedIn && <ObjectsSearch />}
      <ul className="feh-actions">
        {loggedIn ? (
          <BrowserDropdown />
        ) : (
          <a className="btn btn-danger" href={minioBrowserPrefix + "/login"}>
            Login
          </a>
        )}
      </ul>
    </header>
  )
}

export default Header

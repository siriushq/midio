import React from "react"
import { connect } from "react-redux"
import * as actionsObjects from "./actions"

export const ObjectsSearch = ({ onChange }) => (
  <div
    className="input-group ig-left ig-search-dark"
    style={{ display: "block" }}
  >
    <input
      className="ig-text"
      type="input"
      placeholder="Search Objects..."
      onChange={e => onChange(e.target.value)}
    />
    <i className="ig-helpers" />
  </div>
)

const mapDispatchToProps = dispatch => {
  return {
    onChange: filter =>
      dispatch(actionsObjects.setFilter(filter))
  }
}

export default connect(undefined, mapDispatchToProps)(ObjectsSearch)

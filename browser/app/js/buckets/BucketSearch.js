import React from "react"
import { connect } from "react-redux"
import * as actionsBuckets from "./actions"

export const BucketSearch = ({ onChange }) => (
  <div
    className="input-group ig-dark ig-left ig-search"
    style={{ display: "block" }}
  >
    <input
      className="ig-text"
      type="text"
      onChange={e => onChange(e.target.value)}
      placeholder="Search Buckets..."
    />
    <i className="ig-helpers" />
  </div>
)

const mapDispatchToProps = dispatch => {
  return {
    onChange: filter => {
      dispatch(actionsBuckets.setFilter(filter))
    }
  }
}

export default connect(undefined, mapDispatchToProps)(BucketSearch)

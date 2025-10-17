import React from "react"
import { connect } from "react-redux"
import humanize from "humanize"
import Moment from "moment"
import { getDataType } from "../mime"
import * as actions from "./actions"
import { getCheckedList } from "./selectors"

export const ObjectItem = ({
  name,
  contentType,
  size,
  lastModified,
  checked,
  checkObject,
  uncheckObject,
  actionButtons,
  onClick
}) => {
  return (
    <div className={"fesl-row"} data-type={getDataType(name, contentType)}>
      <div className="fesl-item fesl-item-icon">
        <div className="fi-select">
          <input
            type="checkbox"
            name={name}
            checked={checked}
            onChange={() => {
              checked ? uncheckObject(name) : checkObject(name)
            }}
          />
          <i className="fis-icon" />
          <i className="fis-helper" />
        </div>
      </div>
      <div className="fesl-item fesl-item-name">
        <a
          href={getDataType(name, contentType) === "folder" ? name : "#"}
          onClick={e => {
            e.preventDefault()
            // onclick function is passed only when we have a prefix
            if (onClick) {
              onClick()
            } else {
              checked ? uncheckObject(name) : checkObject(name)
            }
          }}
        >
          {name}
        </a>
      </div>
      <div className="fesl-item fesl-item-size">{size}</div>
      <div className="fesl-item fesl-item-modified">{lastModified}</div>
      <div className="fesl-item fesl-item-actions">{actionButtons}</div>
    </div>
  )
}

const mapStateToProps = (state, ownProps) => {
  return {
    checked: getCheckedList(state).indexOf(ownProps.name) >= 0
  }
}

const mapDispatchToProps = dispatch => {
  return {
    checkObject: name => dispatch(actions.checkObject(name)),
    uncheckObject: name => dispatch(actions.uncheckObject(name))
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(ObjectItem)

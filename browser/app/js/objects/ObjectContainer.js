import React from "react"
import { connect } from "react-redux"
import humanize from "humanize"
import Moment from "moment"
import ObjectItem from "./ObjectItem"
import ObjectActions from "./ObjectActions"
import * as actionsObjects from "./actions"
import { getCheckedList } from "./selectors"

export const ObjectContainer = ({
  object,
  checkedObjectsCount,
  downloadObject
}) => {
  let props = {
    name: object.name,
    contentType: object.contentType,
    size: humanize.filesize(object.size),
    lastModified: Moment(object.lastModified).format("lll")
  }
  if (checkedObjectsCount == 0) {
    props.actionButtons = <ObjectActions object={object} />
  }
  return <ObjectItem {...props} />
}

const mapStateToProps = state => {
  return {
    checkedObjectsCount: getCheckedList(state).length
  }
}

export default connect(mapStateToProps)(ObjectContainer)

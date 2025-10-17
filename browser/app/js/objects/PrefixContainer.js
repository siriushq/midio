import React from "react"
import { connect } from "react-redux"
import ObjectItem from "./ObjectItem"
import PrefixActions from "./PrefixActions"
import * as actionsObjects from "./actions"
import { getCheckedList } from "./selectors"

export const PrefixContainer = ({
  object,
  currentPrefix,
  checkedObjectsCount,
  selectPrefix
}) => {
  const props = {
    name: object.name,
    contentType: object.contentType,
    onClick: () => selectPrefix(`${currentPrefix}${object.name}`)
  }
  if (checkedObjectsCount == 0) {
    props.actionButtons = <PrefixActions object={object} />
  }
  return <ObjectItem {...props} />
}

const mapStateToProps = (state, ownProps) => {
  return {
    object: ownProps.object,
    currentPrefix: state.objects.currentPrefix,
    checkedObjectsCount: getCheckedList(state).length
  }
}

const mapDispatchToProps = dispatch => {
  return {
    selectPrefix: prefix => dispatch(actionsObjects.selectPrefix(prefix))
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(PrefixContainer)

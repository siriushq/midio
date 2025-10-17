import React from "react"
import { connect } from "react-redux"
import humanize from "humanize"
import * as actionsCommon from "./actions"

export class StorageInfo extends React.Component {
  componentWillMount() {
    const { fetchStorageInfo } = this.props
    fetchStorageInfo()
  }
  render() {
    const { used } = this.props.storageInfo
    if (!used || used == 0) {
      return <noscript />
    }

    return (
      <div className="feh-used">
        <div className="fehu-chart">
          <div style={{ width: 0 }} />
        </div>
        <ul>
          <li>
            <span>Used: </span>
            {humanize.filesize(used)}
          </li>
        </ul>
      </div>
    )
  }
}

const mapStateToProps = state => {
  return {
    storageInfo: state.browser.storageInfo
  }
}

const mapDispatchToProps = dispatch => {
  return {
    fetchStorageInfo: () => dispatch(actionsCommon.fetchStorageInfo())
  }
}

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(StorageInfo)

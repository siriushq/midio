import React from "react"
import classNames from "classnames"
import { connect } from "react-redux"
import * as actionsBuckets from "./actions"
import { getCurrentBucket } from "./selectors"
import Dropdown from "react-bootstrap/lib/Dropdown"

export class BucketDropdown extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
      showBucketDropdown: false
    }
  }

  toggleDropdown() {
    if (this.state.showBucketDropdown) {
      this.setState({
        showBucketDropdown: false
      })
    } else {
      this.setState({
        showBucketDropdown: true
      })
    }
  }

  render() {
    const { bucket, showBucketPolicy, deleteBucket, currentBucket } = this.props
    return (
      <Dropdown
        open = {this.state.showBucketDropdown}
        onToggle = {this.toggleDropdown.bind(this)}
        className="bucket-dropdown"
        id="bucket-dropdown"
      >
        <Dropdown.Toggle noCaret>
          <i className="zmdi zmdi-more-vert" />
        </Dropdown.Toggle>
        <Dropdown.Menu className="dropdown-menu-right">
          <li>
            <a
              onClick={e => {
                e.stopPropagation()
                this.toggleDropdown()
                showBucketPolicy()
              }}
            >
              Edit policy
            </a>
          </li>
          <li>
            <a
              onClick={e => {
                e.stopPropagation()
                this.toggleDropdown()
                deleteBucket(bucket)
              }}
            >
              Delete
            </a>
          </li>
        </Dropdown.Menu>
      </Dropdown>
    )
  }
}

const mapDispatchToProps = dispatch => {
  return {
    deleteBucket: bucket => dispatch(actionsBuckets.deleteBucket(bucket)),
    showBucketPolicy: () => dispatch(actionsBuckets.showBucketPolicy())
  }
}

export default connect(state => state, mapDispatchToProps)(BucketDropdown)

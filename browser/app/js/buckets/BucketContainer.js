import React from "react"
import { connect } from "react-redux"
import * as actionsBuckets from "./actions"
import { getCurrentBucket } from "./selectors"
import Bucket from "./Bucket"

const mapStateToProps = (state, ownProps) => {
  return {
    isActive: getCurrentBucket(state) === ownProps.bucket
  }
}

const mapDispatchToProps = dispatch => {
  return {
    selectBucket: bucket => dispatch(actionsBuckets.selectBucket(bucket))
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(Bucket)

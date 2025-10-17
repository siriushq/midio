import React from "react"
import classNames from "classnames"
import { connect } from "react-redux"
import ConfirmModal from "../browser/ConfirmModal"
import * as uploadsActions from "./actions"

export class AbortConfirmModal extends React.Component {
  abortUploads() {
    const { abort, uploads } = this.props
    for (var slug in uploads) {
      abort(slug)
    }
  }
  render() {
    const { hideAbort } = this.props
    let baseClass = classNames({
      "abort-upload": true
    })
    let okIcon = classNames({
      fas: true,
      "fa-times": true
    })
    let cancelIcon = classNames({
      fas: true,
      "fa-cloud-upload-alt": true
    })

    return (
      <ConfirmModal
        show={true}
        baseClass={baseClass}
        text="Abort uploads in progress?"
        icon="fas fa-info-circle mci-amber"
        sub="This cannot be undone!"
        okText="Abort"
        okIcon={okIcon}
        cancelText="Upload"
        cancelIcon={cancelIcon}
        okHandler={this.abortUploads.bind(this)}
        cancelHandler={hideAbort}
      />
    )
  }
}

const mapStateToProps = state => {
  return {
    uploads: state.uploads.files
  }
}

const mapDispatchToProps = dispatch => {
  return {
    abort: slug => dispatch(uploadsActions.abortUpload(slug)),
    hideAbort: () => dispatch(uploadsActions.hideAbortModal())
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(AbortConfirmModal)

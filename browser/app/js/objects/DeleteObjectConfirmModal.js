import React from "react"
import ConfirmModal from "../browser/ConfirmModal"

export const DeleteObjectConfirmModal = ({
  deleteObject,
  hideDeleteConfirmModal
}) => (
  <ConfirmModal
    show={true}
    icon="fas fa-exclamation-triangle mci-red"
    text="Are you sure you want to delete?"
    sub="This cannot be undone!"
    okText="Delete"
    cancelText="Cancel"
    okHandler={deleteObject}
    cancelHandler={hideDeleteConfirmModal}
  />
)

export default DeleteObjectConfirmModal

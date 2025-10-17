import React from "react"
import MobileHeader from "./MobileHeader"
import Header from "./Header"
import ObjectsSection from "../objects/ObjectsSection"
import MainActions from "./MainActions"
import BucketPolicyModal from "../buckets/BucketPolicyModal"
import MakeBucketModal from "../buckets/MakeBucketModal"
import UploadModal from "../uploads/UploadModal"
import ObjectsBulkActions from "../objects/ObjectsBulkActions"
import Dropzone from "../uploads/Dropzone"

export const MainContent = () => (
  <div className="fe-body">
    <ObjectsBulkActions />
    <MobileHeader />
    <Dropzone>
      <Header />
      <ObjectsSection />
    </Dropzone>
    <MainActions />
    <BucketPolicyModal />
    <MakeBucketModal />
    <UploadModal />
  </div>
)

export default MainContent

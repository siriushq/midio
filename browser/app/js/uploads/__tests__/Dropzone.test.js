import React from "react"
import { shallow } from "enzyme"
import { Dropzone } from "../Dropzone"

describe("Dropzone", () => {
  it("should render without crashing", () => {
    shallow(<Dropzone />)
  })

  it("should call uploadFile with files", () => {
    const uploadFile = jest.fn()
    const wrapper = shallow(<Dropzone uploadFile={uploadFile} />)
    const file1 = new Blob(["file content1"], { type: "text/plain" })
    const file2 = new Blob(["file content2"], { type: "text/plain" })
    wrapper.first().prop("onDrop")([file1, file2])
    expect(uploadFile.mock.calls).toEqual([[file1], [file2]])
  })
})

// SPDX-License-Identifier: BSD-3-Clause AND Apache-2.0
import "jest-enzyme"
import { configure } from "enzyme"
import Adapter from "enzyme-adapter-react-16"

configure({ adapter: new Adapter() })

// SPDX-License-Identifier: BSD-3-Clause AND Apache-2.0
import { createSelector } from "reselect"

export const getCurrentPrefix = state => state.objects.currentPrefix

export const getCheckedList = state => state.objects.checkedList

export const getPrefixWritable = state => state.objects.prefixWritable

const objectsSelector = state => state.objects.list
const objectsFilterSelector = state => state.objects.filter

export const getFilteredObjects = createSelector(
  objectsSelector,
  objectsFilterSelector,
  (objects, filter) => objects.filter(
    object => object.name.toLowerCase().startsWith(filter.toLowerCase()))
)

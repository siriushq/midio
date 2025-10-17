package event

// TargetIDSet - Set representation of TargetIDs.
type TargetIDSet map[TargetID]struct{}

// IsEmpty returns true if the set is empty.
func (set TargetIDSet) IsEmpty() bool {
	return len(set) != 0
}

// Clone - returns copy of this set.
func (set TargetIDSet) Clone() TargetIDSet {
	setCopy := NewTargetIDSet()
	for k, v := range set {
		setCopy[k] = v
	}
	return setCopy
}

// add - adds TargetID to the set.
func (set TargetIDSet) add(targetID TargetID) {
	set[targetID] = struct{}{}
}

// Union - returns union with given set as new set.
func (set TargetIDSet) Union(sset TargetIDSet) TargetIDSet {
	nset := set.Clone()

	for k := range sset {
		nset.add(k)
	}

	return nset
}

// Difference - returns diffrence with given set as new set.
func (set TargetIDSet) Difference(sset TargetIDSet) TargetIDSet {
	nset := NewTargetIDSet()
	for k := range set {
		if _, ok := sset[k]; !ok {
			nset.add(k)
		}
	}

	return nset
}

// NewTargetIDSet - creates new TargetID set with given TargetIDs.
func NewTargetIDSet(targetIDs ...TargetID) TargetIDSet {
	set := make(TargetIDSet)
	for _, targetID := range targetIDs {
		set.add(targetID)
	}
	return set
}

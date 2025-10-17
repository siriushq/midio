package madmin

import (
	"testing"
)

// Tests heal drives missing and offline counts.
func TestHealDriveCounts(t *testing.T) {
	rs := HealResultItem{}
	rs.Before.Drives = make([]HealDriveInfo, 20)
	rs.After.Drives = make([]HealDriveInfo, 20)
	for i := range rs.Before.Drives {
		if i < 4 {
			rs.Before.Drives[i] = HealDriveInfo{State: DriveStateMissing}
			rs.After.Drives[i] = HealDriveInfo{State: DriveStateMissing}
		} else if i > 4 && i < 15 {
			rs.Before.Drives[i] = HealDriveInfo{State: DriveStateOffline}
			rs.After.Drives[i] = HealDriveInfo{State: DriveStateOffline}
		} else if i > 15 {
			rs.Before.Drives[i] = HealDriveInfo{State: DriveStateCorrupt}
			rs.After.Drives[i] = HealDriveInfo{State: DriveStateCorrupt}
		} else {
			rs.Before.Drives[i] = HealDriveInfo{State: DriveStateOk}
			rs.After.Drives[i] = HealDriveInfo{State: DriveStateOk}
		}
	}

	i, j := rs.GetOnlineCounts()
	if i > 2 {
		t.Errorf("Expected '2', got %d before online disks", i)
	}
	if j > 2 {
		t.Errorf("Expected '2', got %d after online disks", j)
	}
	i, j = rs.GetOfflineCounts()
	if i > 10 {
		t.Errorf("Expected '10', got %d before offline disks", i)
	}
	if j > 10 {
		t.Errorf("Expected '10', got %d after offline disks", j)
	}
	i, j = rs.GetCorruptedCounts()
	if i > 4 {
		t.Errorf("Expected '4', got %d before corrupted disks", i)
	}
	if j > 4 {
		t.Errorf("Expected '4', got %d after corrupted disks", j)
	}
	i, j = rs.GetMissingCounts()
	if i > 4 {
		t.Errorf("Expected '4', got %d before missing disks", i)
	}
	if j > 4 {
		t.Errorf("Expected '4', got %d after missing disks", i)
	}
}

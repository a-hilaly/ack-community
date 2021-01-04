package gencopy

type Comparator struct {
	diffs []string
}

func (c *Comparator) Diffs() []string {
	return c.diffs
}

func (c *Comparator) AppendDiff(fieldPath string) {
	c.diffs = append(c.diffs, fieldPath)
}

func (c *Comparator) stringEqualStrict(fieldPath string, v1, v2 *string) bool {
	if v1 == nil {
		if v2 == nil {
			return true
		}
		c.diffs = append(c.diffs, fieldPath)
		return false
	}
	if v2 == nil {
		c.diffs = append(c.diffs, fieldPath)
		return false
	}
	if *v1 == *v2 {
		return true
	}
	c.diffs = append(c.diffs, fieldPath)
	return false
}

func (c *Comparator) stringEqualSoft(fieldPath string, v1, v2 *string) bool {
	if v1 == nil {
		if v2 == nil || *v2 == "" {
			return true
		}
		c.diffs = append(c.diffs, fieldPath)
		return false
	}
	if v2 == nil {
		if v1 == nil || *v1 == "" {
			return true
		}
		c.diffs = append(c.diffs, fieldPath)
		return false
	}
	if *v1 == *v2 {
		return true
	}
	c.diffs = append(c.diffs, fieldPath)
	return false
}

func (c *Comparator) createBucketConfigurationEqualStrict(fieldPath string, v1, v2 *CreateBucketConfiguration) bool {
	if v1 == nil {
		if v2 == nil {
			return true
		}
		c.diffs = append(c.diffs, fieldPath)
		return false
	}
	if v2 == nil {
		c.diffs = append(c.diffs, fieldPath)
		return false
	}

	// inherit
	if c.stringEqualStrict(fieldPath+".Location", v1.LocationConstraint, v1.LocationConstraint) {
		return true
	}

	c.diffs = append(c.diffs, fieldPath)
	return false
}

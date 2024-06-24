package root

import (
	gt "ssh+/app/git"
)

func GetLongDescription() string {
	return LongDescription + DescriptionVersionApp + gt.GetRelease()
}

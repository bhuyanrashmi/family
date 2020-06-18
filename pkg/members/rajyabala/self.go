package rajyabala

import (
	"bhuyan.io/family-tree/pkg/types"
  "bhuyan.io/family-tree/pkg/members/rajyabala/kankeshwara"
  "bhuyan.io/family-tree/pkg/members/rajyabala/dimbeswar"
  "bhuyan.io/family-tree/pkg/members/rajyabala/nandeshwar"
  "bhuyan.io/family-tree/pkg/members/rajyabala/suchitra"
)

var (
	// Rajyabala is a family member.
	Rajyabala = types.Member{
		Name:     "Rajya Bala(Aitha)",
		Gender:   types.Female,
    Children: []*types.Member{&dimbeswar.Dimbeswar, &kankeshwara.Kankeshwara, &nandeshwar.Nandeshwar, &suchitra.Suchitra, },
	}
)

package dimbeswar

import (
  "bhuyan.io/family-tree/pkg/types"
  "bhuyan.io/family-tree/pkg/members/rajyabala/dimbeswar/debango"
  "bhuyan.io/family-tree/pkg/members/rajyabala/dimbeswar/raunak"
)

var (
  // Dimbeswar is a member of the family.
  Dimbeswar = types.Member{
    Name: "Dimbeswar",
    Children: []*types.Member{&debango.Debango, &raunak.Raunak, },
    Gender: types.Male,
  }
)

package kankeshwara

import (
  "bhuyan.io/family-tree/pkg/types"
  "bhuyan.io/family-tree/pkg/members/rajyabala/kankeshwara/mausam"
  "bhuyan.io/family-tree/pkg/members/rajyabala/kankeshwara/roshan"
)

var (
  // Kankeshwara is a member of the family.
  Kankeshwara = types.Member{
    Name: "Kankeshwara",
    Children: []*types.Member{&mausam.Mausam, &roshan.Roshan, },
    Gender: types.Male,
  }
)

package suchitra

import (
  "bhuyan.io/family-tree/pkg/types"
  "bhuyan.io/family-tree/pkg/members/rajyabala/Suchitra/rashmi"
)

var (
  // Suchitra is a member of the family.
  Suchitra = types.Member{
    Name: "Suchitra - Utanga",
    Children: []*types.Member{&rashmi.Rashmi, },
    Gender: types.Female,
  }
)
